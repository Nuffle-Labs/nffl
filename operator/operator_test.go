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
	"github.com/Layr-Labs/eigensdk-go/metrics"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	safeclientmocks "github.com/NethermindEth/near-sffl/core/safeclient/mocks"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"github.com/NethermindEth/near-sffl/operator/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

const MOCK_OPERATOR_BLS_PRIVATE_KEY = "69"

// hash of bls_public_key (hardcoded for sk=69)
var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}

func TestOperator(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	operator, avsManager, mockConsumer, mockClient, err := createMockOperator(mockCtrl)
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
				QuorumNumbers:    coretypes.QUORUM_NUMBERS_BYTES,
				QuorumThreshold:  uint32(aggtypes.TASK_QUORUM_THRESHOLD),
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)

		X, ok := big.NewInt(0).SetString("17661335844121276474564666950358303590922297397044732665699090882239279390831", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("12547054754368699128461691059748124385981650477514768718744243985037962958213", 10)
		assert.True(t, ok)
		taskResponseSignature := bls.Signature{G1Point: bls.NewG1Point(X, Y)}

		stateRootUpdatesRoot, err := hex.DecodeString("f5d0fae03562975bbec0032ca44d563d0a44d75b6e84b63f4b3c0cbc6e3aeb2c")
		assert.Nil(t, err)

		operatorSetUpdatesRoot, err := hex.DecodeString("47c2698d40a371cd35f175f5a7696d74cf4b2161cb2d6139671a627062fb596d")
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

		X, ok = big.NewInt(0).SetString("13800588136723745139086930992795855876038340441781760930411710515590881085702", 10)
		assert.True(t, ok)
		Y, ok = big.NewInt(0).SetString("13177653237352420178350570044696862987004831431618036354360834407331725483722", 10)
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
		signedOperatorSetUpdateMessage, err := operator.SignOperatorSetUpdate(&messages.OperatorSetUpdateMessage{
			Id:        operatorSetUpdate.Id,
			Timestamp: operatorSetUpdate.Timestamp,
			Operators: make([]coretypes.RollupOperator, 0),
		})
		assert.Nil(t, err)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := mocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedCheckpointTaskResponseToAggregator(signedTaskResponse)
		mockAggregatorRpcClient.EXPECT().SendSignedStateRootUpdateToAggregator(signedStateRootUpdateMessage)
		mockAggregatorRpcClient.EXPECT().SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdateMessage)
		mockAggregatorRpcClient.EXPECT().GetAggregatedCheckpointMessages(newTaskCreatedEvent.Task.FromTimestamp, newTaskCreatedEvent.Task.ToTimestamp).Return(&messages.CheckpointMessages{
			StateRootUpdateMessages:   []messages.StateRootUpdateMessage{signedStateRootUpdateMessage.Message},
			OperatorSetUpdateMessages: []messages.OperatorSetUpdateMessage{signedOperatorSetUpdateMessage.Message},
		}, nil)

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
		mockReader.EXPECT().GetOperatorSetUpdateDelta(gomock.Any(), operatorSetUpdate.Id).Return(make([]opsetupdatereg.RollupOperatorsOperator, 0), nil)

		avsManager.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())

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

		mockClient.EXPECT().Close()

		cancel()

		time.Sleep(1 * time.Second)
	})
}

func createMockOperator(mockCtrl *gomock.Controller) (*Operator, *AvsManager, *mocks.MockConsumer, *safeclientmocks.MockSafeClient, error) {
	logger := sdklogging.NewNoopLogger()
	reg := prometheus.NewRegistry()
	noopMetrics := metrics.NewNoopMetrics()

	blsPrivateKey, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	operatorKeypair := bls.NewKeyPair(blsPrivateKey)

	hasher := messages.NewHasher([32]byte{})
	mockAttestor := mocks.NewMockAttestor(hasher, operatorKeypair, MOCK_OPERATOR_ID)
	avsManager := &AvsManager{
		logger:                       logger,
		checkpointTaskCreatedChan:    make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:        make(chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock),
		operatorSetUpdateMessageChan: make(chan messages.OperatorSetUpdateMessage),
	}
	mockClient := safeclientmocks.NewMockSafeClient(mockCtrl)

	operator := &Operator{
		logger:        logger,
		blsKeypair:    operatorKeypair,
		metricsReg:    reg,
		metrics:       noopMetrics,
		operatorId:    MOCK_OPERATOR_ID,
		attestor:      mockAttestor,
		avsManager:    avsManager,
		listener:      &SelectiveOperatorListener{},
		ethClient:     mockClient,
		messageHasher: hasher,
	}

	return operator, avsManager, mockAttestor.MockGetConsumer(), mockClient, nil
}
