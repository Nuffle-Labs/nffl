package operator

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/NethermindEth/near-sffl/aggregator"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	cstaskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	operatormocks "github.com/NethermindEth/near-sffl/operator/mocks"
)

func TestOperator(t *testing.T) {
	operator, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		var fromNearBlock = uint64(3)
		var toNearBlock = uint64(4)

		newTaskCreatedLog := &cstaskmanager.ContractSFFLTaskManagerCheckpointTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.CheckpointTask{
				FromNearBlock:    fromNearBlock,
				ToNearBlock:      toNearBlock,
				TaskCreatedBlock: 1000,
				QuorumNumbers:    aggtypes.QUORUM_NUMBERS,
				QuorumThreshold:  aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		got := operator.ProcessCheckpointTaskCreatedLog(newTaskCreatedLog)
		want := &cstaskmanager.CheckpointTaskResponse{
			ReferenceTaskIndex:     taskIndex,
			StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var fromNearBlock = uint64(3)
		var toNearBlock = uint64(4)

		// new task event
		newTaskCreatedEvent := &cstaskmanager.ContractSFFLTaskManagerCheckpointTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.CheckpointTask{
				FromNearBlock:    fromNearBlock,
				ToNearBlock:      toNearBlock,
				TaskCreatedBlock: 1000,
				QuorumNumbers:    aggtypes.QUORUM_NUMBERS,
				QuorumThreshold:  aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)
		X, ok := big.NewInt(0).SetString("9996820285347616229516447695531482442433381089408864937966952807215923228881", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("10403462274336311613113322623477208113332192454020049193133394900674966403334", 10)
		assert.True(t, ok)

		signedTaskResponse := &aggregator.SignedCheckpointTaskResponse{
			TaskResponse: cstaskmanager.CheckpointTaskResponse{
				ReferenceTaskIndex:     taskIndex,
				StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			BlsSignature: bls.Signature{
				G1Point: bls.NewG1Point(X, Y),
			},
			OperatorId: operator.operatorId,
		}
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedCheckpointTaskResponseToAggregator(signedTaskResponse)
		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubscribeToNewTasks(operator.checkpointTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}))
		operator.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		operator.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()
		operator.checkpointTaskCreatedChan <- newTaskCreatedEvent
		time.Sleep(1 * time.Second)

		cancel()
	})

}
