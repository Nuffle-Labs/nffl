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
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/metrics"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"github.com/NethermindEth/near-sffl/operator/mocks"
)

const MOCK_OPERATOR_BLS_PRIVATE_KEY = "69"

// hash of bls_public_key (hardcoded for sk=69)
var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}

func TestOperator(t *testing.T) {
	operator, avsManager, mockConsumer, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

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

		X, ok := big.NewInt(0).SetString("16027015062938738578882736302236067956295942129658001187467262823130911146848", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("17647202624711407226560166949876419852295410380838239126346172357046468756471", 10)
		assert.True(t, ok)
		taskResponseSignature := bls.Signature{G1Point: bls.NewG1Point(X, Y)}

		stateRootUpdatesRoot, err := hex.DecodeString("c3566ef4aad0610b0d273388480d8d21f7d07151bd62c428ec3c74f0ffbebf3c")
		assert.Nil(t, err)

		operatorSetUpdatesRoot, err := hex.DecodeString("5ae69791d810e0ec17aa2ec2f67e443f3f7d380079a7e51ff70009f0533aa61e")
		assert.Nil(t, err)

		signedTaskResponse := &messages.SignedCheckpointTaskResponse{
			TaskResponse: messages.CheckpointTaskResponse{
				ReferenceTaskIndex:     taskIndex,
				StateRootUpdatesRoot:   [32]byte(stateRootUpdatesRoot),
				OperatorSetUpdatesRoot: [32]byte(operatorSetUpdatesRoot),
			},
			BlsSignature: taskResponseSignature,
			OperatorId:   operator.operatorId,
		}

		stateRoot, err := hex.DecodeString("04d855ea9fbfefca9069335296aaa5108fa16d36ecd200bf133a1f5b5a7f5fe2")
		assert.Nil(t, err)

		X, ok = big.NewInt(0).SetString("21116328994885950554238991365696504336803036592749146303103711143274485566828", 10)
		assert.True(t, ok)
		Y, ok = big.NewInt(0).SetString("20497381518544159074554997034942551621756618995010334610971970630114874587485", 10)
		assert.True(t, ok)
		stateRootUpdateMessageSignature := bls.Signature{G1Point: bls.NewG1Point(X, Y)}

		block := types.NewBlockWithHeader(&types.Header{
			Number: big.NewInt(0).SetInt64(2),
			Time:   3,
			Root:   common.Hash(stateRoot),
		})
		signedStateRootUpdateMessage := &messages.SignedStateRootUpdateMessage{
			Message: messages.StateRootUpdateMessage{
				RollupId:    1,
				BlockHeight: block.NumberU64(),
				Timestamp:   block.Header().Time,
				// TODO: update this once commitment data fetching is done
				NearDaTransactionId: [32]byte{1},
				NearDaCommitment:    [32]byte{2},
				StateRoot:           [32]byte(stateRoot),
			},
			BlsSignature: stateRootUpdateMessageSignature,
			OperatorId:   operator.operatorId,
		}

		operatorSetUpdate := &opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock{
			Id:        1,
			Timestamp: block.Header().Time,
			Raw:       types.Log{},
		}
		signedOperatorSetUpdateMessage, err := SignOperatorSetUpdate(messages.OperatorSetUpdateMessage{
			Id:        operatorSetUpdate.Id,
			Timestamp: operatorSetUpdate.Timestamp,
			Operators: make([]coretypes.RollupOperator, 0),
		}, operator.blsKeypair, operator.operatorId)
		assert.Nil(t, err)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := mocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedCheckpointTaskResponseToAggregator(signedTaskResponse)
		mockAggregatorRpcClient.EXPECT().SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage)
		mockAggregatorRpcClient.EXPECT().SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage)
		mockAggregatorRpcClient.EXPECT().GetAggregatedCheckpointMessages(newTaskCreatedEvent.Task.FromTimestamp, newTaskCreatedEvent.Task.ToTimestamp, gomock.Any()).SetArg(2, messages.CheckpointMessages{
			StateRootUpdateMessages:   []messages.StateRootUpdateMessage{signedStateRootUpdateMessage.Message},
			OperatorSetUpdateMessages: []messages.OperatorSetUpdateMessage{signedOperatorSetUpdateMessage.Message},
		})

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
		logger:                       logger,
		checkpointTaskCreatedChan:    make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:        make(chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock),
		operatorSetUpdateMessageChan: make(chan messages.OperatorSetUpdateMessage),
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
