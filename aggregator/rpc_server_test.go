package aggregator

import (
	"context"
	"encoding/binary"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/sha3"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

func TestProcessSignedCheckpointTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var FROM_NEAR_BLOCK = uint64(3)
	var TO_NEAR_BLOCK = uint64(4)

	aggregator, _, _, mockBlsAggServ, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	signedCheckpointTaskResponse, err := createMockSignedCheckpointTaskResponse(MockTask{
		TaskNum:       TASK_INDEX,
		BlockNumber:   BLOCK_NUMBER,
		FromTimestamp: FROM_NEAR_BLOCK,
		ToTimestamp:   TO_NEAR_BLOCK,
	}, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedCheckpointTaskResponseDigest, err := signedCheckpointTaskResponse.TaskResponse.Digest()
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedCheckpointTaskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId)
	err = aggregator.ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse, nil)
	assert.Nil(t, err)
}

func TestProcessSignedStateRootUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, mockMessageBlsAggServ, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	message := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaCommitment:    keccak256(4),
		NearDaTransactionId: keccak256(5),
		StateRoot:           keccak256(6),
	}

	signedMessage, err := createMockSignedStateRootUpdateMessage(message, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	messageDigest, err := signedMessage.Message.Digest()
	assert.Nil(t, err)

	mockMessageBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), messageDigest,
		&signedMessage.BlsSignature, signedMessage.OperatorId)
	mockMessageBlsAggServ.EXPECT().InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, types.MESSAGE_TTL, uint64(0))
	err = aggregator.ProcessSignedStateRootUpdateMessage(signedMessage, nil)
	assert.Nil(t, err)
}

func TestProcessOperatorSetUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, mockAvsReader, _, _, _, mockMessageBlsAggServ, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	message := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	signedMessage, err := createMockSignedOperatorSetUpdateMessage(message, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	messageDigest, err := signedMessage.Message.Digest()
	assert.Nil(t, err)

	mockAvsReader.EXPECT().GetOperatorSetUpdateBlock(context.Background(), uint64(1)).Return(uint32(10), nil)

	mockMessageBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), messageDigest,
		&signedMessage.BlsSignature, signedMessage.OperatorId)
	mockMessageBlsAggServ.EXPECT().InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, types.MESSAGE_TTL, uint64(9))
	err = aggregator.ProcessSignedOperatorSetUpdateMessage(signedMessage, nil)
	assert.Nil(t, err)
}

func TestGetAggregatedCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	var checkpointMessages messages.CheckpointMessages

	mockDb.EXPECT().FetchCheckpointMessages(uint64(1), uint64(2), &checkpointMessages)
	err = aggregator.GetAggregatedCheckpointMessages(&GetAggregatedCheckpointMessagesArgs{uint64(1), uint64(2)}, &checkpointMessages)
	assert.Nil(t, err)
}

func keccak256(num uint64) [32]byte {
	var hash [32]byte
	hasher := sha3.NewLegacyKeccak256()
	binary.Write(hasher, binary.LittleEndian, num)
	copy(hash[:], hasher.Sum(nil)[:32])

	return hash
}

func createMockSignedCheckpointTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*messages.SignedCheckpointTaskResponse, error) {
	taskResponse := &messages.CheckpointTaskResponse{
		ReferenceTaskIndex:     mockTask.TaskNum,
		StateRootUpdatesRoot:   keccak256(mockTask.FromTimestamp),
		OperatorSetUpdatesRoot: keccak256(mockTask.ToTimestamp),
	}
	taskResponseHash, err := taskResponse.Digest()
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedCheckpointTaskResponse := &messages.SignedCheckpointTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedCheckpointTaskResponse, nil
}

func createMockSignedStateRootUpdateMessage(mockMessage messages.StateRootUpdateMessage, keypair bls.KeyPair) (*messages.SignedStateRootUpdateMessage, error) {
	messageDigest, err := mockMessage.Digest()
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(messageDigest)
	signedStateRootUpdateMessage := &messages.SignedStateRootUpdateMessage{
		Message:      mockMessage,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedStateRootUpdateMessage, nil
}

func createMockSignedOperatorSetUpdateMessage(mockMessage messages.OperatorSetUpdateMessage, keypair bls.KeyPair) (*messages.SignedOperatorSetUpdateMessage, error) {
	messageDigest, err := mockMessage.Digest()
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(messageDigest)
	signedOperatorSetUpdateMessage := &messages.SignedOperatorSetUpdateMessage{
		Message:      mockMessage,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedOperatorSetUpdateMessage, nil
}
