package aggregator

import (
	"context"
	"encoding/binary"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/sha3"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	"github.com/NethermindEth/near-sffl/core"
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

	aggregator, _, _, mockBlsAggServ, _, _, mockOperatorRegistrationsServ, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
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
	ctx := context.Background()
	mockBlsAggServ.EXPECT().ProcessNewSignature(ctx, TASK_INDEX, signedCheckpointTaskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId)
	mockOperatorRegistrationsServ.EXPECT().GetOperatorInfoById(ctx, signedCheckpointTaskResponse.OperatorId).Return(eigentypes.OperatorInfo{Pubkeys: MOCK_OPERATOR_PUBKEYS}, true)

	err = aggregator.ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse)
	assert.Nil(t, err)
}

func TestProcessSignedStateRootUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, mockMessageBlsAggServ, _, mockOperatorRegistrationsServ, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	aggregator.clock = core.Clock{Now: func() time.Time { return time.Unix(10_000, 0) }}
	message := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           9_995,
		NearDaCommitment:    keccak256(4),
		NearDaTransactionId: keccak256(5),
		StateRoot:           keccak256(6),
	}

	signedMessage, err := createMockSignedStateRootUpdateMessage(message, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	messageDigest, err := signedMessage.Message.Digest()
	assert.Nil(t, err)

	mockMessageBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), messageDigest, &signedMessage.BlsSignature, signedMessage.OperatorId)
	mockMessageBlsAggServ.EXPECT().InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []eigentypes.QuorumThresholdPercentage{types.MESSAGE_AGGREGATION_QUORUM_THRESHOLD}, types.MESSAGE_TTL, types.MESSAGE_BLS_AGGREGATION_TIMEOUT, uint64(0))
	mockOperatorRegistrationsServ.EXPECT().GetOperatorInfoById(context.Background(), signedMessage.OperatorId).Return(eigentypes.OperatorInfo{Pubkeys: MOCK_OPERATOR_PUBKEYS}, true)

	err = aggregator.ProcessSignedStateRootUpdateMessage(signedMessage)
	assert.Nil(t, err)
}

func TestProcessInvalidSignedStateRootUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockOperatorRegistrationsServ, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	aggregator.clock = core.Clock{Now: func() time.Time { return time.Unix(10_000, 0) }}
	message := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           9_995,
		NearDaCommitment:    keccak256(4),
		NearDaTransactionId: keccak256(5),
		StateRoot:           keccak256(6),
	}

	signedMessage, err := createMockSignedStateRootUpdateMessage(message, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedMessage.BlsSignature = *newInvalidSignature()

	mockOperatorRegistrationsServ.EXPECT().GetOperatorInfoById(context.Background(), signedMessage.OperatorId).Return(eigentypes.OperatorInfo{Pubkeys: MOCK_OPERATOR_PUBKEYS}, true)
	err = aggregator.ProcessSignedStateRootUpdateMessage(signedMessage)
	assert.Equal(t, err.Error(), "Invalid signature")
}

func TestProcessOperatorSetUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, mockAvsReader, _, _, _, mockMessageBlsAggServ, mockOperatorRegistrationsServ, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	aggregator.clock = core.Clock{Now: func() time.Time { return time.Unix(10_000, 0) }}
	message := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 9_995,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	signedMessage, err := createMockSignedOperatorSetUpdateMessage(message, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	messageDigest, err := signedMessage.Message.Digest()
	assert.Nil(t, err)

	ctx := context.Background()
	mockAvsReader.EXPECT().GetOperatorSetUpdateBlock(ctx, uint64(1)).Return(uint32(10), nil)

	mockMessageBlsAggServ.EXPECT().ProcessNewSignature(ctx, messageDigest,
		&signedMessage.BlsSignature, signedMessage.OperatorId)
	mockMessageBlsAggServ.EXPECT().InitializeMessageIfNotExists(messageDigest, coretypes.QUORUM_NUMBERS, []eigentypes.QuorumThresholdPercentage{types.MESSAGE_AGGREGATION_QUORUM_THRESHOLD}, types.MESSAGE_TTL, types.MESSAGE_BLS_AGGREGATION_TIMEOUT, uint64(9))
	mockOperatorRegistrationsServ.EXPECT().GetOperatorInfoById(context.Background(), signedMessage.OperatorId).Return(eigentypes.OperatorInfo{Pubkeys: MOCK_OPERATOR_PUBKEYS}, true)

	err = aggregator.ProcessSignedOperatorSetUpdateMessage(signedMessage)
	assert.Nil(t, err)
}

func TestGetAggregatedCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	var checkpointMessages messages.CheckpointMessages
	mockDb.EXPECT().FetchCheckpointMessages(uint64(1), uint64(2)).Return(&checkpointMessages, nil)
	_, err = aggregator.GetAggregatedCheckpointMessages(uint64(1), uint64(2))
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

func newInvalidSignature() *bls.Signature {
	return bls.NewZeroSignature()
}
