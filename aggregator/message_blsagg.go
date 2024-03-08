package aggregator

// BLS aggregator service for SFFL messages, based on eigensdk-go's blsagg

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

var (
	MessageAlreadyInitializedErrorFn = func(messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("message %x already initialized", messageDigest)
	}
	MessageExpiredError    = fmt.Errorf("message expired")
	MessageNotFoundErrorFn = func(messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("message %x not initialized or already completed", messageDigest)
	}
	OperatorNotPartOfMessageQuorumErrorFn = func(operatorId types.OperatorId, messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("operator %x not part of message %x's quorum", operatorId, messageDigest)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

type AggregatedOperators struct {
	signersApkG2               *bls.G2Point
	signersAggSigG1            *bls.Signature
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
	signersOperatorIdsSet      map[types.OperatorId]bool
}

type SignedMessageDigest struct {
	MessageDigest               coretypes.MessageDigest
	BlsSignature                *bls.Signature
	OperatorId                  bls.OperatorId
	SignatureVerificationErrorC chan error
}

type signedMessageDigestValidationInfo struct {
	operatorsAvsStateDict         map[[32]byte]types.OperatorAvsState
	quorumsAvsStakeDict           map[uint8]types.QuorumAvsState
	totalStakePerQuorum           map[uint8]*big.Int
	quorumApksG1                  []*bls.G1Point
	aggregatedOperatorsDict       map[types.TaskResponseDigest]AggregatedOperators
	quorumThresholdPercentagesMap map[uint8]uint32
	quorumNumbers                 []types.QuorumNum
	ethBlockNumber                uint64
}

type MessageBlsAggregationService interface {
	InitializeMessageIfNotExists(
		messageDigest coretypes.MessageDigest,
		quorumNumbers []types.QuorumNum,
		quorumThresholdPercentages []types.QuorumThresholdPercentage,
		timeToExpiry time.Duration,
		ethBlockNumber uint64,
	) error

	ProcessNewSignature(
		ctx context.Context,
		messageDigest coretypes.MessageDigest,
		blsSignature *bls.Signature,
		operatorId bls.OperatorId,
	) error

	GetResponseChannel() <-chan aggtypes.MessageBlsAggregationServiceResponse
}

type MessageBlsAggregatorService struct {
	aggregatedResponsesC   chan aggtypes.MessageBlsAggregationServiceResponse
	signedMessageDigestsCs map[coretypes.MessageDigest]chan SignedMessageDigest
	messageChansLock       sync.RWMutex
	avsRegistryService     avsregistry.AvsRegistryService
	ethClient              eth.EthClient
	logger                 logging.Logger
}

var _ MessageBlsAggregationService = (*MessageBlsAggregatorService)(nil)

func NewMessageBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, ethClient eth.EthClient, logger logging.Logger) *MessageBlsAggregatorService {
	return &MessageBlsAggregatorService{
		aggregatedResponsesC:   make(chan aggtypes.MessageBlsAggregationServiceResponse),
		signedMessageDigestsCs: make(map[coretypes.MessageDigest]chan SignedMessageDigest),
		messageChansLock:       sync.RWMutex{},
		avsRegistryService:     avsRegistryService,
		ethClient:              ethClient,
		logger:                 logger,
	}
}

func (mbas *MessageBlsAggregatorService) GetResponseChannel() <-chan aggtypes.MessageBlsAggregationServiceResponse {
	return mbas.aggregatedResponsesC
}

func (mbas *MessageBlsAggregatorService) InitializeMessageIfNotExists(
	messageDigest coretypes.MessageDigest,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	ethBlockNumber uint64,
) error {
	mbas.messageChansLock.Lock()
	defer mbas.messageChansLock.Unlock()

	if _, taskExists := mbas.signedMessageDigestsCs[messageDigest]; taskExists {
		return nil
	}

	signedMessageDigestsC := make(chan SignedMessageDigest)
	mbas.signedMessageDigestsCs[messageDigest] = signedMessageDigestsC
	go mbas.singleMessageAggregatorGoroutineFunc(messageDigest, quorumNumbers, quorumThresholdPercentages, timeToExpiry, signedMessageDigestsC, ethBlockNumber)

	return nil
}

func (mbas *MessageBlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	messageDigest coretypes.MessageDigest,
	blsSignature *bls.Signature,
	operatorId bls.OperatorId,
) error {
	mbas.messageChansLock.RLock()
	defer mbas.messageChansLock.RUnlock()

	messageC, taskInitialized := mbas.signedMessageDigestsCs[messageDigest]

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

func (mbas *MessageBlsAggregatorService) singleMessageAggregatorGoroutineFunc(
	messageDigest coretypes.MessageDigest,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	signedMessageDigestsC <-chan SignedMessageDigest,
	ethBlockNumber uint64,
) {
	defer mbas.closeMessageGoroutine(messageDigest)

	validationInfo := mbas.fetchValidationInfo(quorumNumbers, quorumThresholdPercentages, ethBlockNumber)
	messageExpiredTimer := time.NewTimer(timeToExpiry)

	for {
		select {
		case signedMessageDigest := <-signedMessageDigestsC:
			mbas.logger.Debug("Message goroutine received new signed message digest", "messageDigest", messageDigest)
			mbas.handleSignedMessageDigest(signedMessageDigest, validationInfo)
			return
		case <-messageExpiredTimer.C:
			mbas.aggregatedResponsesC <- aggtypes.MessageBlsAggregationServiceResponse{
				Err: MessageExpiredError,
			}
			return
		}
	}
}

func (mbas *MessageBlsAggregatorService) fetchValidationInfo(quorumNumbers []types.QuorumNum, quorumThresholdPercentages []types.QuorumThresholdPercentage, ethBlockNumber uint64) signedMessageDigestValidationInfo {
	if ethBlockNumber == 0 {
		curEthBlockNumber, err := mbas.ethClient.BlockNumber(context.Background())
		if err != nil {
			mbas.logger.Fatal("Aggregator failed to get current block number", "err", err)
		}

		ethBlockNumber = curEthBlockNumber
	}

	operatorsAvsStateDict, err := mbas.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(ethBlockNumber))
	if err != nil {
		// TODO: how should we handle such an error?
		mbas.logger.Fatal("Aggregator failed to get operators state from avs registry", "err", err)
	}

	quorumsAvsStakeDict, err := mbas.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(ethBlockNumber))
	if err != nil {
		mbas.logger.Fatal("Aggregator failed to get quorums state from avs registry", "err", err)
	}

	totalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}

	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	quorumThresholdPercentagesMap := make(map[types.QuorumNum]types.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}

	return signedMessageDigestValidationInfo{
		operatorsAvsStateDict:         operatorsAvsStateDict,
		quorumsAvsStakeDict:           quorumsAvsStakeDict,
		totalStakePerQuorum:           totalStakePerQuorum,
		quorumApksG1:                  quorumApksG1,
		aggregatedOperatorsDict:       make(map[types.TaskResponseDigest]AggregatedOperators),
		quorumThresholdPercentagesMap: quorumThresholdPercentagesMap,
		quorumNumbers:                 quorumNumbers,
		ethBlockNumber:                ethBlockNumber,
	}
}

func (mbas *MessageBlsAggregatorService) handleSignedMessageDigest(signedMessageDigest SignedMessageDigest, validationInfo signedMessageDigestValidationInfo) {
	err := mbas.verifySignature(signedMessageDigest, validationInfo.operatorsAvsStateDict)
	signedMessageDigest.SignatureVerificationErrorC <- err

	if err != nil {
		return
	}

	digestAggregatedOperators, ok := validationInfo.aggregatedOperatorsDict[signedMessageDigest.MessageDigest]
	if !ok {
		digestAggregatedOperators = AggregatedOperators{
			// we've already verified that the operator is part of the task's quorum, so we don't need checks here
			signersApkG2:               bls.NewZeroG2Point().Add(validationInfo.operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey),
			signersAggSigG1:            signedMessageDigest.BlsSignature,
			signersOperatorIdsSet:      map[types.OperatorId]bool{signedMessageDigest.OperatorId: true},
			signersTotalStakePerQuorum: validationInfo.operatorsAvsStateDict[signedMessageDigest.OperatorId].StakePerQuorum,
		}
	} else {
		digestAggregatedOperators.signersAggSigG1.Add(signedMessageDigest.BlsSignature)
		digestAggregatedOperators.signersApkG2.Add(validationInfo.operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey)
		digestAggregatedOperators.signersOperatorIdsSet[signedMessageDigest.OperatorId] = true
		for quorumNum, stake := range validationInfo.operatorsAvsStateDict[signedMessageDigest.OperatorId].StakePerQuorum {
			if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
				digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
			}
			digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
		}
	}
	// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
	// because of https://github.com/golang/go/issues/3117
	validationInfo.aggregatedOperatorsDict[signedMessageDigest.MessageDigest] = digestAggregatedOperators

	if !checkIfStakeThresholdsMet(digestAggregatedOperators.signersTotalStakePerQuorum, validationInfo.totalStakePerQuorum, validationInfo.quorumThresholdPercentagesMap) {
		return
	}

	nonSignersOperatorIds := []types.OperatorId{}
	for operatorId := range validationInfo.operatorsAvsStateDict {
		if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
		}
	}

	indices, err := mbas.avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, uint32(validationInfo.ethBlockNumber), validationInfo.quorumNumbers, nonSignersOperatorIds)
	if err != nil {
		mbas.logger.Error("Failed to get check signatures indices", "err", err)
		mbas.aggregatedResponsesC <- aggtypes.MessageBlsAggregationServiceResponse{
			Err: err,
		}
		return
	}

	messageBlsAggregationServiceResponse := aggtypes.MessageBlsAggregationServiceResponse{
		Err:                          nil,
		EthBlockNumber:               validationInfo.ethBlockNumber,
		MessageDigest:                signedMessageDigest.MessageDigest,
		NonSignersPubkeysG1:          getG1PubkeysOfNonSigners(digestAggregatedOperators.signersOperatorIdsSet, validationInfo.operatorsAvsStateDict),
		QuorumApksG1:                 validationInfo.quorumApksG1,
		SignersApkG2:                 digestAggregatedOperators.signersApkG2,
		SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
		NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             indices.QuorumApkIndices,
		TotalStakeIndices:            indices.TotalStakeIndices,
		NonSignerStakeIndices:        indices.NonSignerStakeIndices,
	}

	mbas.aggregatedResponsesC <- messageBlsAggregationServiceResponse
}

func (mbas *MessageBlsAggregatorService) closeMessageGoroutine(messageDigest coretypes.MessageDigest) {
	mbas.messageChansLock.Lock()
	delete(mbas.signedMessageDigestsCs, messageDigest)
	mbas.messageChansLock.Unlock()
}

func (mbas *MessageBlsAggregatorService) verifySignature(
	signedMessageDigest SignedMessageDigest,
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedMessageDigest.OperatorId]
	if !ok {
		mbas.logger.Warnf("Operator %#v not found. Skipping message.", signedMessageDigest.OperatorId)
		return OperatorNotPartOfMessageQuorumErrorFn(signedMessageDigest.OperatorId, signedMessageDigest.MessageDigest)
	}

	// 0. verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedMessageDigest.OperatorId].Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		mbas.logger.Fatal("Operator G2 pubkey not found")
	}
	mbas.logger.Debug("Verifying signed message digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"messageDigest", signedMessageDigest.MessageDigest,
		"blsSignature", signedMessageDigest.BlsSignature,
	)
	signatureVerified, err := signedMessageDigest.BlsSignature.Verify(operatorG2Pubkey, signedMessageDigest.MessageDigest)
	if err != nil {
		mbas.logger.Error(SignatureVerificationError(err).Error())
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		mbas.logger.Error(IncorrectSignatureError.Error())
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
