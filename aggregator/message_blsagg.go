package aggregator

// BLS aggregator service for SFFL messages, based on eigensdk-go's blsagg

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	MessageAlreadyInitializedErrorFn = func(messageDigest aggtypes.MessageDigest) error {
		return fmt.Errorf("message %x already initialized", messageDigest)
	}
	MessageExpiredError    = fmt.Errorf("message expired")
	MessageNotFoundErrorFn = func(messageDigest aggtypes.MessageDigest) error {
		return fmt.Errorf("message %x not initialized or already completed", messageDigest)
	}
	OperatorNotPartOfMessageQuorumErrorFn = func(operatorId types.OperatorId, messageDigest aggtypes.MessageDigest) error {
		return fmt.Errorf("operator %x not part of message %x's quorum", operatorId, messageDigest)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

type MessageBlsAggregationServiceResponse struct {
	Err                          error
	EthBlockNumber               uint64
	MessageDigest                aggtypes.MessageDigest
	NonSignersPubkeysG1          []*bls.G1Point
	QuorumApksG1                 []*bls.G1Point
	SignersApkG2                 *bls.G2Point
	SignersAggSigG1              *bls.Signature
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

type AggregatedOperators struct {
	signersApkG2               *bls.G2Point
	signersAggSigG1            *bls.Signature
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
	signersOperatorIdsSet      map[types.OperatorId]bool
}

type SignedMessageDigest struct {
	MessageDigest               aggtypes.MessageDigest
	BlsSignature                *bls.Signature
	OperatorId                  bls.OperatorId
	SignatureVerificationErrorC chan error
}

type MessageBlsAggregationService interface {
	InitializeNewMessage(
		messageDigest aggtypes.MessageDigest,
		quorumNumbers []types.QuorumNum,
		quorumThresholdPercentages []types.QuorumThresholdPercentage,
		timeToExpiry time.Duration,
	) error

	ProcessNewSignature(
		ctx context.Context,
		messageDigest aggtypes.MessageDigest,
		blsSignature *bls.Signature,
		operatorId bls.OperatorId,
	) error

	GetResponseChannel() <-chan MessageBlsAggregationServiceResponse
}

type MessageBlsAggregatorService struct {
	aggregatedResponsesC   chan MessageBlsAggregationServiceResponse
	signedMessageDigestsCs map[aggtypes.MessageDigest]chan SignedMessageDigest
	messageChansMutex      sync.RWMutex
	avsRegistryService     avsregistry.AvsRegistryService
	ethClient              eth.EthClient
	logger                 logging.Logger
}

var _ MessageBlsAggregationService = (*MessageBlsAggregatorService)(nil)

func NewMessageBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, ethClient eth.EthClient, logger logging.Logger) *MessageBlsAggregatorService {
	return &MessageBlsAggregatorService{
		aggregatedResponsesC:   make(chan MessageBlsAggregationServiceResponse),
		signedMessageDigestsCs: make(map[aggtypes.MessageDigest]chan SignedMessageDigest),
		messageChansMutex:      sync.RWMutex{},
		avsRegistryService:     avsRegistryService,
		ethClient:              ethClient,
		logger:                 logger,
	}
}

func (a *MessageBlsAggregatorService) GetResponseChannel() <-chan MessageBlsAggregationServiceResponse {
	return a.aggregatedResponsesC
}

func (a *MessageBlsAggregatorService) InitializeNewMessage(
	messageDigest aggtypes.MessageDigest,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
) error {
	if _, taskExists := a.signedMessageDigestsCs[messageDigest]; taskExists {
		return MessageAlreadyInitializedErrorFn(messageDigest)
	}
	signedMessageDigestsC := make(chan SignedMessageDigest)
	a.messageChansMutex.Lock()
	a.signedMessageDigestsCs[messageDigest] = signedMessageDigestsC
	a.messageChansMutex.Unlock()
	go a.singleMessageAggregatorGoroutineFunc(messageDigest, quorumNumbers, quorumThresholdPercentages, timeToExpiry, signedMessageDigestsC)
	return nil
}

func (a *MessageBlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	messageDigest aggtypes.MessageDigest,
	blsSignature *bls.Signature,
	operatorId bls.OperatorId,
) error {
	a.messageChansMutex.Lock()
	messageC, taskInitialized := a.signedMessageDigestsCs[messageDigest]
	a.messageChansMutex.Unlock()
	if !taskInitialized {
		return MessageNotFoundErrorFn(messageDigest)
	}
	signatureVerificationErrorC := make(chan error)

	select {
	case messageC <- SignedMessageDigest{
		MessageDigest:               messageDigest,
		BlsSignature:                blsSignature,
		OperatorId:                  operatorId,
		SignatureVerificationErrorC: signatureVerificationErrorC,
	}:
		return <-signatureVerificationErrorC
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *MessageBlsAggregatorService) singleMessageAggregatorGoroutineFunc(
	messageDigest aggtypes.MessageDigest,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	signedMessageDigestsC <-chan SignedMessageDigest,
) {
	defer a.closeMessageGoroutine(messageDigest)

	curBlockNum, err := a.ethClient.BlockNumber(context.Background())

	if err != nil {
		a.logger.Fatal("Aggregator failed to get current block number", "err", err)
	}

	operatorsAvsStateDict, err := a.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(curBlockNum))
	if err != nil {
		// TODO: how should we handle such an error?
		a.logger.Fatal("Aggregator failed to get operators state from avs registry", "err", err)
	}

	quorumsAvsStakeDict, err := a.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(curBlockNum))
	if err != nil {
		a.logger.Fatal("Aggregator failed to get quorums state from avs registry", "err", err)
	}

	totalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}

	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	messageExpiredTimer := time.NewTimer(timeToExpiry)

	quorumThresholdPercentagesMap := make(map[types.QuorumNum]types.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}

	aggregatedOperatorsDict := map[types.TaskResponseDigest]AggregatedOperators{}
	for {
		select {
		case signedMessageDigest := <-signedMessageDigestsC:
			a.logger.Debug("Message goroutine received new signed message digest", "messageDigest", messageDigest)

			signedMessageDigest.SignatureVerificationErrorC <- a.verifySignature(messageDigest, signedMessageDigest, operatorsAvsStateDict)
			digestAggregatedOperators, ok := aggregatedOperatorsDict[signedMessageDigest.MessageDigest]
			if !ok {
				digestAggregatedOperators = AggregatedOperators{
					// we've already verified that the operator is part of the task's quorum, so we don't need checks here
					signersApkG2:               bls.NewZeroG2Point().Add(operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey),
					signersAggSigG1:            signedMessageDigest.BlsSignature,
					signersOperatorIdsSet:      map[types.OperatorId]bool{signedMessageDigest.OperatorId: true},
					signersTotalStakePerQuorum: operatorsAvsStateDict[signedMessageDigest.OperatorId].StakePerQuorum,
				}
			} else {
				digestAggregatedOperators.signersAggSigG1.Add(signedMessageDigest.BlsSignature)
				digestAggregatedOperators.signersApkG2.Add(operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey)
				digestAggregatedOperators.signersOperatorIdsSet[signedMessageDigest.OperatorId] = true
				for quorumNum, stake := range operatorsAvsStateDict[signedMessageDigest.OperatorId].StakePerQuorum {
					if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
						digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
					}
					digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
				}
			}
			// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
			// because of https://github.com/golang/go/issues/3117
			aggregatedOperatorsDict[signedMessageDigest.MessageDigest] = digestAggregatedOperators

			if checkIfStakeThresholdsMet(digestAggregatedOperators.signersTotalStakePerQuorum, totalStakePerQuorum, quorumThresholdPercentagesMap) {
				nonSignersOperatorIds := []types.OperatorId{}
				for operatorId := range operatorsAvsStateDict {
					if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
						nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
					}
				}

				indices, err := a.avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, uint32(curBlockNum), quorumNumbers, nonSignersOperatorIds)
				if err != nil {
					a.logger.Error("Failed to get check signatures indices", "err", err)
					a.aggregatedResponsesC <- MessageBlsAggregationServiceResponse{
						Err: err,
					}
					return
				}

				messageBlsAggregationServiceResponse := MessageBlsAggregationServiceResponse{
					Err:                          nil,
					EthBlockNumber:               curBlockNum,
					MessageDigest:                messageDigest,
					NonSignersPubkeysG1:          getG1PubkeysOfNonSigners(digestAggregatedOperators.signersOperatorIdsSet, operatorsAvsStateDict),
					QuorumApksG1:                 quorumApksG1,
					SignersApkG2:                 digestAggregatedOperators.signersApkG2,
					SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
					NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
					QuorumApkIndices:             indices.QuorumApkIndices,
					TotalStakeIndices:            indices.TotalStakeIndices,
					NonSignerStakeIndices:        indices.NonSignerStakeIndices,
				}

				a.aggregatedResponsesC <- messageBlsAggregationServiceResponse
				return
			}
		case <-messageExpiredTimer.C:
			a.aggregatedResponsesC <- MessageBlsAggregationServiceResponse{
				Err: MessageExpiredError,
			}
			return
		}
	}

}

func (a *MessageBlsAggregatorService) closeMessageGoroutine(messageDigest aggtypes.MessageDigest) {
	a.messageChansMutex.Lock()
	delete(a.signedMessageDigestsCs, messageDigest)
	a.messageChansMutex.Unlock()
}

func (a *MessageBlsAggregatorService) verifySignature(
	messageDigest aggtypes.MessageDigest,
	signedMessageDigest SignedMessageDigest,
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedMessageDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedMessageDigest.OperatorId)
		return OperatorNotPartOfMessageQuorumErrorFn(signedMessageDigest.OperatorId, messageDigest)
	}

	// 0. verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		a.logger.Fatal("Operator G2 pubkey not found")
	}
	a.logger.Debug("Verifying signed message digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"messageDigest", signedMessageDigest.MessageDigest,
		"blsSignature", signedMessageDigest.BlsSignature,
	)
	signatureVerified, err := signedMessageDigest.BlsSignature.Verify(operatorG2Pubkey, signedMessageDigest.MessageDigest)
	if err != nil {
		a.logger.Error(SignatureVerificationError(err).Error())
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		a.logger.Error(IncorrectSignatureError.Error())
		return IncorrectSignatureError
	}
	return nil
}

func checkIfStakeThresholdsMet(
	signedStakePerQuorum map[types.QuorumNum]*big.Int,
	totalStakePerQuorum map[types.QuorumNum]*big.Int,
	quorumThresholdPercentagesMap map[types.QuorumNum]types.QuorumThresholdPercentage,
) bool {
	for quorumNum, quorumThresholdPercentage := range quorumThresholdPercentagesMap {
		signedStake := big.NewInt(0).Mul(signedStakePerQuorum[quorumNum], big.NewInt(100))
		thresholdStake := big.NewInt(0).Mul(totalStakePerQuorum[quorumNum], big.NewInt(int64(quorumThresholdPercentage)))
		if signedStake.Cmp(thresholdStake) < 0 {
			return false
		}
	}
	return true
}

func getG1PubkeysOfNonSigners(signersOperatorIdsSet map[types.OperatorId]bool, operatorAvsStateDict map[[32]byte]types.OperatorAvsState) []*bls.G1Point {
	nonSignersG1Pubkeys := []*bls.G1Point{}
	for operatorId, operator := range operatorAvsStateDict {
		if _, operatorSigned := signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.Pubkeys.G1Pubkey)
		}
	}
	return nonSignersG1Pubkeys
}
