package operator

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/metrics"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"github.com/NethermindEth/near-sffl/operator/mocks"
	operatormocks "github.com/NethermindEth/near-sffl/operator/mocks"
)

const MOCK_OPERATOR_BLS_PRIVATE_KEY = "69"

// hash of bls_public_key (hardcoded for sk=69)
var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}

func TestOperator(t *testing.T) {
	operator, avsManager, mockConsumer, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		var fromTimestamp = uint64(3)
		var toTimestamp = uint64(4)

		newTaskCreatedLog := &taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated{
			TaskIndex: taskIndex,
			Task: taskmanager.CheckpointTask{
				FromTimestamp:    fromTimestamp,
				ToTimestamp:      toTimestamp,
				TaskCreatedBlock: 1000,
				QuorumNumbers:    coretypes.QUORUM_NUMBERS,
				QuorumThreshold:  aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		got := avsManager.ProcessCheckpointTaskCreatedLog(newTaskCreatedLog)
		want := taskmanager.CheckpointTaskResponse{
			ReferenceTaskIndex:     taskIndex,
			StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var fromTimestamp = uint64(3)
		var toTimestamp = uint64(4)

		// new task event
		newTaskCreatedEvent := &taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated{
			TaskIndex: taskIndex,
			Task: taskmanager.CheckpointTask{
				FromTimestamp:    fromTimestamp,
				ToTimestamp:      toTimestamp,
				TaskCreatedBlock: 1000,
				QuorumNumbers:    coretypes.QUORUM_NUMBERS,
				QuorumThreshold:  aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)

		X, ok := big.NewInt(0).SetString("9996820285347616229516447695531482442433381089408864937966952807215923228881", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("10403462274336311613113322623477208113332192454020049193133394900674966403334", 10)
		assert.True(t, ok)
		taskResponseSignature := bls.Signature{G1Point: bls.NewG1Point(X, Y)}

		signedTaskResponse := &coretypes.SignedCheckpointTaskResponse{
			TaskResponse: taskmanager.CheckpointTaskResponse{
				ReferenceTaskIndex:     taskIndex,
				StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			BlsSignature: taskResponseSignature,
			OperatorId:   operator.operatorId,
		}

		stateRoot, err := hex.DecodeString("04d855ea9fbfefca9069335296aaa5108fa16d36ecd200bf133a1f5b5a7f5fe2")
		assert.Nil(t, err)

		X, ok = big.NewInt(0).SetString("11290223320119541059506650081001398560835662629535951352639311217634399300639", 10)
		assert.True(t, ok)
		Y, ok = big.NewInt(0).SetString("13733759189688049392641040177512184151496910984279779370238986310773365898082", 10)
		assert.True(t, ok)
		stateRootUpdateMessageSignature := bls.Signature{G1Point: bls.NewG1Point(X, Y)}

		block := types.NewBlockWithHeader(&types.Header{
			Number: big.NewInt(0).SetInt64(2),
			Time:   3,
			Root:   common.Hash(stateRoot),
		})
		signedStateRootUpdateMessage := &coretypes.SignedStateRootUpdateMessage{
			Message: servicemanager.StateRootUpdateMessage{
				RollupId:    1,
				BlockHeight: block.NumberU64(),
				Timestamp:   block.Header().Time,
				StateRoot:   [32]byte(stateRoot),
			},
			BlsSignature: stateRootUpdateMessageSignature,
			OperatorId:   operator.operatorId,
		}

		operatorSetUpdate := &opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock{
			Id:        1,
			Timestamp: block.Header().Time,
			Raw:       types.Log{},
		}
		signedOperatorSetUpdateMessage, err := SignOperatorSetUpdate(registryrollup.OperatorSetUpdateMessage{
			Id:        operatorSetUpdate.Id,
			Timestamp: operatorSetUpdate.Timestamp,
			Operators: make([]registryrollup.OperatorsOperator, 0),
		}, operator.blsKeypair, operator.operatorId)
		assert.Nil(t, err)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedCheckpointTaskResponseToAggregator(signedTaskResponse)
		mockAggregatorRpcClient.EXPECT().SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage)
		mockAggregatorRpcClient.EXPECT().SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage)

		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubscribeToNewTasks(avsManager.checkpointTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}), nil)
		mockSubscriber.EXPECT().SubscribeToOperatorSetUpdates(avsManager.operatorSetUpdateChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}), nil)
		avsManager.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		mockReader.EXPECT().GetOperatorSetUpdateDelta(gomock.Any(), operatorSetUpdate.Id).Return(make([]opsetupdatereg.OperatorsOperator, 0), nil)

		avsManager.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()

		avsManager.checkpointTaskCreatedChan <- newTaskCreatedEvent
		avsManager.operatorSetUpdateChan <- operatorSetUpdate
		mockConsumer.MockReceiveBlockData(consumer.BlockData{
			RollupId: signedStateRootUpdateMessage.Message.RollupId,
			Block:    *block,
		})

		time.Sleep(1 * time.Second)
	})
}

func createMockOperator() (*Operator, *AvsManager, *mocks.MockConsumer, error) {
	logger := sdklogging.NewNoopLogger()
	reg := prometheus.NewRegistry()
	noopMetrics := metrics.NewNoopMetrics()

	blsPrivateKey, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	if err != nil {
		return nil, nil, nil, err
	}
	operatorKeypair := bls.NewKeyPair(blsPrivateKey)

	mockAttestor := mocks.NewMockAttestor(operatorKeypair, MOCK_OPERATOR_ID)
	avsManager := &AvsManager{
		logger:                            logger,
		checkpointTaskCreatedChan:         make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:             make(chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock),
		checkpointTaskResponseCreatedChan: make(chan taskmanager.CheckpointTaskResponse),
		operatorSetUpdateMessageChan:      make(chan registryrollup.OperatorSetUpdateMessage),
	}

	operator := &Operator{
		logger:     logger,
		blsKeypair: operatorKeypair,
		metricsReg: reg,
		metrics:    noopMetrics,
		operatorId: MOCK_OPERATOR_ID,
		attestor:   mockAttestor,
		avsManager: avsManager,
	}

	return operator, avsManager, mockAttestor.MockGetConsumer(), nil
}
