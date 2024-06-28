// File directly copied from https://github.com/Layr-Labs/eigensdk-go/blob/cc3b8d2dd3390ce407aadc76130c4b3fa583f9b5/services/bls_aggregation/blsagg.go
// with slight modifications to naming

package aggregator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

var (
	TaskInitializationErrorFn = func(err error, taskIndex eigentypes.TaskIndex) error {
		return fmt.Errorf("Failed to initialize task %d: %w", taskIndex, err)
	}
	TaskAlreadyInitializedErrorFn = func(taskIndex eigentypes.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredErrorFn = func(taskIndex eigentypes.TaskIndex) error {
		return fmt.Errorf("task %d expired", taskIndex)
	}
	TaskNotFoundErrorFn = func(taskIndex eigentypes.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorFn = func(operatorId eigentypes.OperatorId, taskIndex eigentypes.TaskIndex) error {
		return fmt.Errorf("operator %x not part of task %d's quorum", operatorId, taskIndex)
	}
)

type aggregatedOperators struct {
	signersApkG2               *bls.G2Point
	signersAggSigG1            *bls.Signature
	signersTotalStakePerQuorum map[eigentypes.QuorumNum]*big.Int
	signersOperatorIdsSet      map[eigentypes.OperatorId]bool
}

type TaskBlsAggregationService interface {
	InitializeNewTask(
		taskIndex eigentypes.TaskIndex,
		taskCreatedBlock uint32,
		quorumNumbers eigentypes.QuorumNums,
		quorumThresholdPercentages eigentypes.QuorumThresholdPercentages,
		timeToExpiry time.Duration,
	) error

	ProcessNewSignature(
		ctx context.Context,
		taskIndex eigentypes.TaskIndex,
		taskResponseDigest eigentypes.TaskResponseDigest,
		blsSignature *bls.Signature,
		operatorId eigentypes.OperatorId,
	) error

	GetResponseChannel() <-chan types.TaskBlsAggregationServiceResponse
}

type TaskBlsAggregatorService struct {
	aggregatedResponsesC chan types.TaskBlsAggregationServiceResponse
	signedTaskRespsCs    map[eigentypes.TaskIndex]chan eigentypes.SignedTaskResponseDigest
	taskChansMutex       sync.RWMutex
	avsRegistryService   avsregistry.AvsRegistryService
	logger               logging.Logger
}

var _ TaskBlsAggregationService = (*TaskBlsAggregatorService)(nil)

func NewTaskBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, logger logging.Logger) *TaskBlsAggregatorService {
	return &TaskBlsAggregatorService{
		aggregatedResponsesC: make(chan types.TaskBlsAggregationServiceResponse),
		signedTaskRespsCs:    make(map[eigentypes.TaskIndex]chan eigentypes.SignedTaskResponseDigest),
		taskChansMutex:       sync.RWMutex{},
		avsRegistryService:   avsRegistryService,
		logger:               logger,
	}
}

func (a *TaskBlsAggregatorService) GetResponseChannel() <-chan types.TaskBlsAggregationServiceResponse {
	return a.aggregatedResponsesC
}

func (a *TaskBlsAggregatorService) InitializeNewTask(
	taskIndex eigentypes.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers eigentypes.QuorumNums,
	quorumThresholdPercentages eigentypes.QuorumThresholdPercentages,
	timeToExpiry time.Duration,
) error {
	a.logger.Debug("AggregatorService initializing new task", "taskIndex", taskIndex, "taskCreatedBlock", taskCreatedBlock, "quorumNumbers", quorumNumbers, "quorumThresholdPercentages", quorumThresholdPercentages, "timeToExpiry", timeToExpiry)
	if _, taskExists := a.signedTaskRespsCs[taskIndex]; taskExists {
		return TaskAlreadyInitializedErrorFn(taskIndex)
	}
	signedTaskRespsC := make(chan eigentypes.SignedTaskResponseDigest)
	a.taskChansMutex.Lock()
	a.signedTaskRespsCs[taskIndex] = signedTaskRespsC
	a.taskChansMutex.Unlock()
	go a.singleTaskAggregatorGoroutineFunc(taskIndex, taskCreatedBlock, quorumNumbers, quorumThresholdPercentages, timeToExpiry, signedTaskRespsC)
	return nil
}

func (a *TaskBlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	taskIndex eigentypes.TaskIndex,
	taskResponseDigest eigentypes.TaskResponseDigest,
	blsSignature *bls.Signature,
	operatorId eigentypes.OperatorId,
) error {
	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.signedTaskRespsCs[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return TaskNotFoundErrorFn(taskIndex)
	}
	signatureVerificationErrorC := make(chan error)
	select {
	case taskC <- eigentypes.SignedTaskResponseDigest{
		TaskResponseDigest:          taskResponseDigest,
		BlsSignature:                blsSignature,
		OperatorId:                  operatorId,
		SignatureVerificationErrorC: signatureVerificationErrorC,
	}:
		return <-signatureVerificationErrorC
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *TaskBlsAggregatorService) singleTaskAggregatorGoroutineFunc(
	taskIndex eigentypes.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers eigentypes.QuorumNums,
	quorumThresholdPercentages []eigentypes.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	signedTaskRespsC <-chan eigentypes.SignedTaskResponseDigest,
) {
	defer a.closeTaskGoroutine(taskIndex)

	quorumThresholdPercentagesMap := make(map[eigentypes.QuorumNum]eigentypes.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}
	operatorsAvsStateDict, err := a.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		a.aggregatedResponsesC <- types.TaskBlsAggregationServiceResponse{
			Err: TaskInitializationErrorFn(fmt.Errorf("AggregatorService failed to get operators state from avs registry at blockNum %d: %w", taskCreatedBlock, err), taskIndex),
		}
		return
	}
	quorumsAvsStakeDict, err := a.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		a.aggregatedResponsesC <- types.TaskBlsAggregationServiceResponse{
			Err: TaskInitializationErrorFn(fmt.Errorf("Aggregator failed to get quorums state from avs registry: %w", err), taskIndex),
		}
		return
	}
	totalStakePerQuorum := make(map[eigentypes.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}
	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	taskExpiredTimer := time.NewTimer(timeToExpiry)

	aggregatedOperatorsDict := map[eigentypes.TaskResponseDigest]aggregatedOperators{}
	for {
		select {
		case signedTaskResponseDigest := <-signedTaskRespsC:
			a.logger.Debug("Task goroutine received new signed task response digest", "taskIndex", taskIndex, "signedTaskResponseDigest", signedTaskResponseDigest)
			err := a.verifySignature(taskIndex, signedTaskResponseDigest, operatorsAvsStateDict)
			signedTaskResponseDigest.SignatureVerificationErrorC <- err
			if err != nil {
				continue
			}
			digestAggregatedOperators, ok := aggregatedOperatorsDict[signedTaskResponseDigest.TaskResponseDigest]
			if !ok {
				digestAggregatedOperators = aggregatedOperators{
					signersApkG2:               bls.NewZeroG2Point().Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey),
					signersAggSigG1:            signedTaskResponseDigest.BlsSignature,
					signersOperatorIdsSet:      map[eigentypes.OperatorId]bool{signedTaskResponseDigest.OperatorId: true},
					signersTotalStakePerQuorum: operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum,
				}
			} else {
				digestAggregatedOperators.signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
				digestAggregatedOperators.signersApkG2.Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey)
				digestAggregatedOperators.signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] = true
				for quorumNum, stake := range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
					if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
						digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
					}
					digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
				}
			}
			aggregatedOperatorsDict[signedTaskResponseDigest.TaskResponseDigest] = digestAggregatedOperators

			if checkIfStakeThresholdsMet(digestAggregatedOperators.signersTotalStakePerQuorum, totalStakePerQuorum, quorumThresholdPercentagesMap) {
				nonSignersOperatorIds := []eigentypes.OperatorId{}
				for operatorId := range operatorsAvsStateDict {
					if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
						nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
					}
				}

				sort.SliceStable(nonSignersOperatorIds, func(i, j int) bool {
					iOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[i][:])
					jOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[j][:])
					return iOprInt.Cmp(jOprInt) == -1
				})

				nonSignersG1Pubkeys := []*bls.G1Point{}
				for _, operatorId := range nonSignersOperatorIds {
					operator := operatorsAvsStateDict[operatorId]
					nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.OperatorInfo.Pubkeys.G1Pubkey)
				}

				indices, err := a.avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, taskCreatedBlock, quorumNumbers, nonSignersOperatorIds)
				if err != nil {
					a.aggregatedResponsesC <- types.TaskBlsAggregationServiceResponse{
						Err: utils.WrapError(errors.New("Failed to get check signatures indices"), err),
					}
					return
				}
				taskBlsAggregationServiceResponse := types.TaskBlsAggregationServiceResponse{
					Err:                nil,
					TaskIndex:          taskIndex,
					TaskResponseDigest: signedTaskResponseDigest.TaskResponseDigest,
					TaskBlsAggregation: messages.TaskBlsAggregation{
						NonSignersPubkeysG1:          nonSignersG1Pubkeys,
						QuorumApksG1:                 quorumApksG1,
						SignersApkG2:                 digestAggregatedOperators.signersApkG2,
						SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
						NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
						QuorumApkIndices:             indices.QuorumApkIndices,
						TotalStakeIndices:            indices.TotalStakeIndices,
						NonSignerStakeIndices:        indices.NonSignerStakeIndices,
					},
				}
				a.aggregatedResponsesC <- taskBlsAggregationServiceResponse
				taskExpiredTimer.Stop()
				return
			}
		case <-taskExpiredTimer.C:
			a.aggregatedResponsesC <- types.TaskBlsAggregationServiceResponse{
				Err: TaskExpiredErrorFn(taskIndex),
			}
			return
		}
	}

}

func (a *TaskBlsAggregatorService) closeTaskGoroutine(taskIndex eigentypes.TaskIndex) {
	a.taskChansMutex.Lock()
	delete(a.signedTaskRespsCs, taskIndex)
	a.taskChansMutex.Unlock()
}

func (a *TaskBlsAggregatorService) verifySignature(
	taskIndex eigentypes.TaskIndex,
	signedTaskResponseDigest eigentypes.SignedTaskResponseDigest,
	operatorsAvsStateDict map[eigentypes.OperatorId]eigentypes.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedTaskResponseDigest.OperatorId)
		return OperatorNotPartOfTaskQuorumErrorFn(signedTaskResponseDigest.OperatorId, taskIndex)
	}

	operatorG2Pubkey := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		a.logger.Error("Operator G2 pubkey not found", "operatorId", signedTaskResponseDigest.OperatorId, "taskId", taskIndex)
		return fmt.Errorf("taskId %d: Operator G2 pubkey not found (operatorId: %x)", taskIndex, signedTaskResponseDigest.OperatorId)
	}
	a.logger.Debug("Verifying signed task response digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"taskResponseDigest", signedTaskResponseDigest.TaskResponseDigest,
		"blsSignature", signedTaskResponseDigest.BlsSignature,
	)
	signatureVerified, err := signedTaskResponseDigest.BlsSignature.Verify(operatorG2Pubkey, signedTaskResponseDigest.TaskResponseDigest)
	if err != nil {
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		return IncorrectSignatureError
	}
	return nil
}
