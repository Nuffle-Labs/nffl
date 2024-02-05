package aggregator

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/sha3"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/NethermindEth/near-sffl/aggregator/types"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
)

func TestProcessSignedCheckpointTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var FROM_NEAR_BLOCK = uint64(3)
	var TO_NEAR_BLOCK = uint64(4)

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[bls.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}
	aggregator, _, mockBlsAggServ, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	signedCheckpointTaskResponse, err := createMockSignedCheckpointTaskResponse(MockTask{
		TaskNum:       TASK_INDEX,
		BlockNumber:   BLOCK_NUMBER,
		FromNearBlock: FROM_NEAR_BLOCK,
		ToNearBlock:   TO_NEAR_BLOCK,
	}, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedCheckpointTaskResponseDigest, err := core.GetCheckpointTaskResponseDigest(&signedCheckpointTaskResponse.TaskResponse)
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedCheckpointTaskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId)
	err = aggregator.ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse, nil)
	assert.Nil(t, err)
}

func keccak256(num uint64) [32]byte {
	var hash [32]byte
	hasher := sha3.NewLegacyKeccak256()
	binary.Write(hasher, binary.LittleEndian, num)
	copy(hash[:], hasher.Sum(nil)[:32])

	return hash
}

// mocks an operator signing on a task response
func createMockSignedCheckpointTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*SignedCheckpointTaskResponse, error) {
	taskResponse := &taskmanager.CheckpointTaskResponse{
		ReferenceTaskIndex:     mockTask.TaskNum,
		StateRootUpdatesRoot:   keccak256(mockTask.FromNearBlock),
		OperatorSetUpdatesRoot: keccak256(mockTask.ToNearBlock),
	}
	taskResponseHash, err := core.GetCheckpointTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedCheckpointTaskResponse := &SignedCheckpointTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedCheckpointTaskResponse, nil
}
