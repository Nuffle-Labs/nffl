package blsagg

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
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

var (
	MessageAlreadyInitializedErrorFn = func(messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("message 0x%x already initialized", messageDigest)
	}
	MessageExpiredError    = fmt.Errorf("message expired")
	MessageNotFoundErrorFn = func(messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("message 0x%x not initialized or already completed", messageDigest)
	}
	OperatorNotPartOfMessageQuorumErrorFn = func(operatorId eigentypes.OperatorId, messageDigest coretypes.MessageDigest) error {
		return fmt.Errorf("operator 0x%x not part of message 0x%x's quorum", operatorId, messageDigest)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError    = errors.New("Signature verification failed. Incorrect Signature.")
	MessageDigestNotFoundError = errors.New("Message digest not found")
)

type MessageBlsAggregationStatus int32

const (
	MessageBlsAggregationStatusNone MessageBlsAggregationStatus = iota
	MessageBlsAggregationStatusFullStakeThresholdMet
	MessageBlsAggregationStatusThresholdNotReached
	MessageBlsAggregationStatusThresholdReached
)

type MessageBlsAggregationServiceResponse struct {
	messages.MessageBlsAggregation

	MessageKey coretypes.MessageKey
	Message    interface{}

	Status   MessageBlsAggregationStatus
	Finished bool
	Err      error
}

type AggregatedOperators struct {
	signersApkG2               *bls.G2Point
	signersAggSigG1            *bls.Signature
	signersTotalStakePerQuorum map[eigentypes.QuorumNum]*big.Int
	signersOperatorIdsSet      map[eigentypes.OperatorId]bool
}

type SignedMessage struct {
	Message                     MessageBlsAggregationServiceMessage
	MessageDigest               coretypes.MessageDigest
	BlsSignature                *bls.Signature
	OperatorId                  eigentypes.OperatorId
	SignatureVerificationErrorC chan error
}

type signedMessageDigestValidationInfo struct {
	operatorsAvsStateDict         map[eigentypes.OperatorId]eigentypes.OperatorAvsState
	quorumsAvsStakeDict           map[eigentypes.QuorumNum]eigentypes.QuorumAvsState
	totalStakePerQuorum           map[eigentypes.QuorumNum]*big.Int
	quorumApksG1                  []*bls.G1Point
	aggregatedOperatorsDict       map[eigentypes.TaskResponseDigest]AggregatedOperators
	quorumThresholdPercentagesMap map[eigentypes.QuorumNum]eigentypes.QuorumThresholdPercentage
	quorumNumbers                 []eigentypes.QuorumNum
	ethBlockNumber                uint64
}

type MessageBlsAggregationServiceMessage interface {
	Digest() (coretypes.MessageDigest, error)
	Key() coretypes.MessageKey
}

type MessageBlsAggregationService interface {
	InitializeMessageIfNotExists(
		messageKey coretypes.MessageKey,
		quorumNumbers []eigentypes.QuorumNum,
		quorumThresholdPercentages []eigentypes.QuorumThresholdPercentage,
		timeToExpiry time.Duration,
		aggregationTimeout time.Duration,
		ethBlockNumber uint64,
	) error

	ProcessNewSignature(
		ctx context.Context,
		message MessageBlsAggregationServiceMessage,
		blsSignature *bls.Signature,
		operatorId eigentypes.OperatorId,
	) error

	GetResponseChannel() <-chan MessageBlsAggregationServiceResponse
}

type MessageBlsAggregatorService struct {
	aggregatedResponsesC chan MessageBlsAggregationServiceResponse
	signedMessageCs      map[coretypes.MessageKey]chan SignedMessage
	messageChansLock     sync.RWMutex
	avsRegistryService   avsregistry.AvsRegistryService
	ethClient            eth.Client
	logger               logging.Logger
}

var _ MessageBlsAggregationService = (*MessageBlsAggregatorService)(nil)

func NewMessageBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, ethClient eth.Client, logger logging.Logger) *MessageBlsAggregatorService {
	return &MessageBlsAggregatorService{
		aggregatedResponsesC: make(chan MessageBlsAggregationServiceResponse),
		signedMessageCs:      make(map[coretypes.MessageKey]chan SignedMessage),
		messageChansLock:     sync.RWMutex{},
		avsRegistryService:   avsRegistryService,
		ethClient:            ethClient,
		logger:               logger,
	}
}

func (mbas *MessageBlsAggregatorService) GetResponseChannel() <-chan MessageBlsAggregationServiceResponse {
	return mbas.aggregatedResponsesC
}

func (mbas *MessageBlsAggregatorService) initializeMessageChan(messageKey coretypes.MessageKey) chan SignedMessage {
	mbas.messageChansLock.Lock()
	defer mbas.messageChansLock.Unlock()

	if _, taskExists := mbas.signedMessageCs[messageKey]; taskExists {
		return nil
	}

	signedMessageC := make(chan SignedMessage)
	mbas.signedMessageCs[messageKey] = signedMessageC

	return signedMessageC
}

func (mbas *MessageBlsAggregatorService) InitializeMessageIfNotExists(
	messageKey coretypes.MessageKey,
	quorumNumbers []eigentypes.QuorumNum,
	quorumThresholdPercentages []eigentypes.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	aggregationTimeout time.Duration,
	ethBlockNumber uint64,
) error {
	signedMessageC := mbas.initializeMessageChan(messageKey)
	if signedMessageC == nil {
		return nil
	}

	validationInfo, err := mbas.fetchValidationInfo(quorumNumbers, quorumThresholdPercentages, ethBlockNumber)
	if err != nil {
		return err
	}

	go mbas.singleMessageAggregatorGoroutineFunc(
		messageKey,
		validationInfo,
		timeToExpiry,
		aggregationTimeout,
		signedMessageC,
	)

	return nil
}

func (mbas *MessageBlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	message MessageBlsAggregationServiceMessage,
	blsSignature *bls.Signature,
	operatorId eigentypes.OperatorId,
) error {
	mbas.messageChansLock.RLock()
	messageC, taskInitialized := mbas.signedMessageCs[message.Key()]
	mbas.messageChansLock.RUnlock()

	if !taskInitialized {
		return MessageNotFoundErrorFn(message.Key())
	}
	signatureVerificationErrorC := make(chan error)

	messageDigest, err := message.Digest()
	if err != nil {
		return MessageDigestNotFoundError
	}

	select {
	case messageC <- SignedMessage{
		Message:                     message,
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
	messageKey coretypes.MessageKey,
	validationInfo *signedMessageDigestValidationInfo,
	timeToExpiry time.Duration,
	aggregationTimeout time.Duration,
	signedMessageC <-chan SignedMessage,
) {
	defer mbas.closeMessageGoroutine(messageKey)

	message, shouldWaitForFullStake := mbas.handleSignedMessagePreThreshold(messageKey, validationInfo, timeToExpiry, signedMessageC)
	if shouldWaitForFullStake {
		mbas.handleSignedMessageThresholdReached(message, validationInfo, signedMessageC, aggregationTimeout)
	}
}

func (mbas *MessageBlsAggregatorService) handleSignedMessagePreThreshold(
	messageKey coretypes.MessageKey,
	validationInfo *signedMessageDigestValidationInfo,
	timeToExpiry time.Duration,
	signedMessageC <-chan SignedMessage,
) (MessageBlsAggregationServiceMessage, bool) {
	messageExpiredTimer := time.NewTimer(timeToExpiry)
	defer messageExpiredTimer.Stop()

	for {
		select {
		case signedMessage := <-signedMessageC:
			mbas.logger.Debug("Message goroutine received new signed message", "key", messageKey)

			err := mbas.handleSignedMessageDigest(signedMessage, validationInfo)
			signedMessage.SignatureVerificationErrorC <- err
			if err != nil {
				continue
			}

			aggregation := mbas.getMessageBlsAggregationResponse(signedMessage.Message, signedMessage.MessageDigest, validationInfo, false)

			if aggregation.Status == MessageBlsAggregationStatusThresholdNotReached {
				continue
			}

			mbas.aggregatedResponsesC <- aggregation

			if aggregation.Status == MessageBlsAggregationStatusThresholdReached {
				return signedMessage.Message, true
			} else if aggregation.Status == MessageBlsAggregationStatusFullStakeThresholdMet {
				return signedMessage.Message, false
			}
		case <-messageExpiredTimer.C:
			mbas.logger.Debug("Message expired", "key", messageKey)
			mbas.aggregatedResponsesC <- MessageBlsAggregationServiceResponse{
				MessageBlsAggregation: messages.MessageBlsAggregation{
					EthBlockNumber: validationInfo.ethBlockNumber,
				},
				MessageKey: messageKey,
				Message:    nil,
				Finished:   true,
				Err:        MessageExpiredError,
			}

			return nil, true
		}
	}
}

func (mbas *MessageBlsAggregatorService) handleSignedMessageThresholdReached(
	message MessageBlsAggregationServiceMessage,
	validationInfo *signedMessageDigestValidationInfo,
	signedMessageC <-chan SignedMessage,
	aggregationTimeout time.Duration,
) {
	thresholdReachedTimer := time.NewTimer(aggregationTimeout)
	defer thresholdReachedTimer.Stop()

	messageKey := message.Key()
	messageDigest, err := message.Digest()
	if err != nil {
		mbas.logger.Fatal("Failed to get message digest, should be unreachable", "err", err)
		return
	}

	for {
		select {
		case signedMessage := <-signedMessageC:
			mbas.logger.Debug("Message goroutine received new signed message", "key", messageKey)

			err := mbas.handleSignedMessageDigest(signedMessage, validationInfo)
			signedMessage.SignatureVerificationErrorC <- err
			if err != nil {
				continue
			}

			aggregation := mbas.getMessageBlsAggregationResponse(signedMessage.Message, signedMessage.MessageDigest, validationInfo, false)
			mbas.aggregatedResponsesC <- aggregation

			if aggregation.Status == MessageBlsAggregationStatusFullStakeThresholdMet {
				return
			}
		case <-thresholdReachedTimer.C:
			mbas.logger.Debug("Message expired", "key", messageKey)
			mbas.aggregatedResponsesC <- mbas.getMessageBlsAggregationResponse(message, messageDigest, validationInfo, true)
			return
		}
	}
}

func (mbas *MessageBlsAggregatorService) fetchValidationInfo(quorumNumbers []eigentypes.QuorumNum, quorumThresholdPercentages []eigentypes.QuorumThresholdPercentage, ethBlockNumber uint64) (*signedMessageDigestValidationInfo, error) {
	if ethBlockNumber == 0 {
		curEthBlockNumber, err := mbas.ethClient.BlockNumber(context.Background())
		if err != nil {
			mbas.logger.Fatal("Aggregator failed to get current block number", "err", err)
		}

		ethBlockNumber = curEthBlockNumber
	}

	operatorsAvsStateDict, err := mbas.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(ethBlockNumber))
	if err != nil {
		mbas.logger.Error("Aggregator failed to get operators state from avs registry", "err", err)
		return nil, err
	}

	quorumsAvsStakeDict, err := mbas.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, uint32(ethBlockNumber))
	if err != nil {
		mbas.logger.Error("Aggregator failed to get quorums state from avs registry", "err", err)
		return nil, err
	}

	totalStakePerQuorum := make(map[eigentypes.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}

	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	quorumThresholdPercentagesMap := make(map[eigentypes.QuorumNum]eigentypes.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}

	return &signedMessageDigestValidationInfo{
		operatorsAvsStateDict:         operatorsAvsStateDict,
		quorumsAvsStakeDict:           quorumsAvsStakeDict,
		totalStakePerQuorum:           totalStakePerQuorum,
		quorumApksG1:                  quorumApksG1,
		aggregatedOperatorsDict:       make(map[eigentypes.TaskResponseDigest]AggregatedOperators),
		quorumThresholdPercentagesMap: quorumThresholdPercentagesMap,
		quorumNumbers:                 quorumNumbers,
		ethBlockNumber:                ethBlockNumber,
	}, nil
}

type SignedMessageHandlingResult int32

const (
	SignedMessageHandlingResultThresholdReached SignedMessageHandlingResult = iota
	SignedMessageHandlingResultFullStakeThresholdMet
	SignedMessageHandlingResultSuccess
	SignedMessageHandlingResultError
)

func (mbas *MessageBlsAggregatorService) handleSignedMessageDigest(signedMessage SignedMessage, validationInfo *signedMessageDigestValidationInfo) error {
	err := mbas.verifySignature(signedMessage, validationInfo.operatorsAvsStateDict)
	if err != nil {
		return err
	}

	digest, err := signedMessage.Message.Digest()
	if err != nil {
		return err
	}

	digestAggregatedOperators, ok := validationInfo.aggregatedOperatorsDict[digest]
	if !ok {
		digestAggregatedOperators = AggregatedOperators{
			// we've already verified that the operator is part of the task's quorum, so we don't need checks here
			signersApkG2:               bls.NewZeroG2Point().Add(validationInfo.operatorsAvsStateDict[signedMessage.OperatorId].OperatorInfo.Pubkeys.G2Pubkey),
			signersAggSigG1:            signedMessage.BlsSignature,
			signersOperatorIdsSet:      map[eigentypes.OperatorId]bool{signedMessage.OperatorId: true},
			signersTotalStakePerQuorum: validationInfo.operatorsAvsStateDict[signedMessage.OperatorId].StakePerQuorum,
		}
	} else {
		digestAggregatedOperators.signersAggSigG1.Add(signedMessage.BlsSignature)
		digestAggregatedOperators.signersApkG2.Add(validationInfo.operatorsAvsStateDict[signedMessage.OperatorId].OperatorInfo.Pubkeys.G2Pubkey)
		digestAggregatedOperators.signersOperatorIdsSet[signedMessage.OperatorId] = true
		for quorumNum, stake := range validationInfo.operatorsAvsStateDict[signedMessage.OperatorId].StakePerQuorum {
			if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
				digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
			}
			digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
		}
	}
	// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
	// because of https://github.com/golang/go/issues/3117
	validationInfo.aggregatedOperatorsDict[digest] = digestAggregatedOperators

	return nil
}

func (mbas *MessageBlsAggregatorService) getMessageBlsAggregationStatus(messageDigest coretypes.MessageDigest, validationInfo *signedMessageDigestValidationInfo) (MessageBlsAggregationStatus, error) {
	digestAggregatedOperators, ok := validationInfo.aggregatedOperatorsDict[messageDigest]
	if !ok {
		return MessageBlsAggregationStatusNone, MessageNotFoundErrorFn(messageDigest)
	}

	fullStakeThresholdMet := checkIfFullStakeThresholdMet(digestAggregatedOperators.signersTotalStakePerQuorum, validationInfo.totalStakePerQuorum)
	if fullStakeThresholdMet {
		return MessageBlsAggregationStatusFullStakeThresholdMet, nil
	}

	if checkIfStakeThresholdsMet(digestAggregatedOperators.signersTotalStakePerQuorum, validationInfo.totalStakePerQuorum, validationInfo.quorumThresholdPercentagesMap) {
		return MessageBlsAggregationStatusThresholdReached, nil
	}

	return MessageBlsAggregationStatusThresholdNotReached, nil
}

func (mbas *MessageBlsAggregatorService) getMessageBlsAggregationResponse(message MessageBlsAggregationServiceMessage, messageDigest coretypes.MessageDigest, validationInfo *signedMessageDigestValidationInfo, forceFinished bool) MessageBlsAggregationServiceResponse {
	defaultAggregation := messages.MessageBlsAggregation{
		MessageDigest:  messageDigest,
		EthBlockNumber: validationInfo.ethBlockNumber,
	}

	digestAggregatedOperators, ok := validationInfo.aggregatedOperatorsDict[messageDigest]
	if !ok {
		return MessageBlsAggregationServiceResponse{
			MessageBlsAggregation: defaultAggregation,
			Message:               message,
			Err:                   MessageNotFoundErrorFn(messageDigest),
		}
	}

	status, err := mbas.getMessageBlsAggregationStatus(messageDigest, validationInfo)
	if err != nil {
		return MessageBlsAggregationServiceResponse{
			MessageBlsAggregation: defaultAggregation,
			Message:               message,
			Err:                   err,
		}
	}

	nonSignersOperatorIds := []eigentypes.OperatorId{}
	for operatorId := range validationInfo.operatorsAvsStateDict {
		if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
		}
	}

	indices, err := mbas.avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, uint32(validationInfo.ethBlockNumber), validationInfo.quorumNumbers, nonSignersOperatorIds)
	if err != nil {
		mbas.logger.Error("Failed to get check signatures indices", "err", err)
		return MessageBlsAggregationServiceResponse{
			MessageBlsAggregation: defaultAggregation,
			Message:               message,
			Err:                   err,
		}
	}

	aggregation, err := messages.StandardizeMessageBlsAggregation(
		messages.MessageBlsAggregation{
			EthBlockNumber:               uint64(validationInfo.ethBlockNumber),
			MessageDigest:                messageDigest,
			NonSignersPubkeysG1:          getG1PubkeysOfNonSigners(digestAggregatedOperators.signersOperatorIdsSet, validationInfo.operatorsAvsStateDict),
			QuorumApksG1:                 validationInfo.quorumApksG1,
			SignersApkG2:                 digestAggregatedOperators.signersApkG2,
			SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
			NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
			QuorumApkIndices:             indices.QuorumApkIndices,
			TotalStakeIndices:            indices.TotalStakeIndices,
			NonSignerStakeIndices:        indices.NonSignerStakeIndices,
		},
	)
	if err != nil {
		mbas.logger.Error("Failed to format aggregation", "err", err)
		return MessageBlsAggregationServiceResponse{
			MessageBlsAggregation: defaultAggregation,
			Message:               message,
			Err:                   err,
		}
	}

	fullStakeThresholdMet := status == MessageBlsAggregationStatusFullStakeThresholdMet

	return MessageBlsAggregationServiceResponse{
		Err:                   nil,
		Status:                status,
		Finished:              fullStakeThresholdMet || forceFinished,
		MessageBlsAggregation: aggregation,
		Message:               message,
	}
}

func (mbas *MessageBlsAggregatorService) closeMessageGoroutine(messageDigest coretypes.MessageDigest) {
	mbas.messageChansLock.Lock()
	delete(mbas.signedMessageCs, messageDigest)
	mbas.messageChansLock.Unlock()
}

func (mbas *MessageBlsAggregatorService) verifySignature(
	signedMessage SignedMessage,
	operatorsAvsStateDict map[eigentypes.OperatorId]eigentypes.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedMessage.OperatorId]
	if !ok {
		mbas.logger.Warn("Operator not found. Skipping message", "operator", signedMessage.OperatorId)
		return OperatorNotPartOfMessageQuorumErrorFn(signedMessage.OperatorId, signedMessage.MessageDigest)
	}

	// 0. verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedMessage.OperatorId].OperatorInfo.Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		mbas.logger.Fatal("Operator G2 pubkey not found")
	}
	mbas.logger.Debug("Verifying signed message digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"messageDigest", signedMessage.MessageDigest,
		"blsSignature", signedMessage.BlsSignature,
	)
	signatureVerified, err := signedMessage.BlsSignature.Verify(operatorG2Pubkey, signedMessage.MessageDigest)
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
	signedStakePerQuorum map[eigentypes.QuorumNum]*big.Int,
	totalStakePerQuorum map[eigentypes.QuorumNum]*big.Int,
	quorumThresholdPercentagesMap map[eigentypes.QuorumNum]eigentypes.QuorumThresholdPercentage,
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

func checkIfFullStakeThresholdMet(
	signedStakePerQuorum map[eigentypes.QuorumNum]*big.Int,
	totalStakePerQuorum map[eigentypes.QuorumNum]*big.Int,
) bool {
	for quorumNum, signedStake := range signedStakePerQuorum {
		if signedStake.Cmp(totalStakePerQuorum[quorumNum]) != 0 {
			return false
		}
	}
	return true
}

func getG1PubkeysOfNonSigners(signersOperatorIdsSet map[eigentypes.OperatorId]bool, operatorAvsStateDict map[eigentypes.OperatorId]eigentypes.OperatorAvsState) []*bls.G1Point {
	nonSignersG1Pubkeys := []*bls.G1Point{}
	for operatorId, operator := range operatorAvsStateDict {
		if _, operatorSigned := signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.OperatorInfo.Pubkeys.G1Pubkey)
		}
	}
	return nonSignersG1Pubkeys
}
