package database_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NethermindEth/near-sffl/aggregator/database"
	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/tests"
)

func TestFetchUnknownStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	entry, err := db.FetchStateRootUpdate(1, 2)
	assert.NotNil(t, err)
	assert.Nil(t, entry)
}

func TestStoreAndFetchStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	value := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaTransactionId: tests.Keccak256(4),
		NearDaCommitment:    tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}

	_, err = db.StoreStateRootUpdate(value)
	assert.Nil(t, err)

	entry, err := db.FetchStateRootUpdate(value.RollupId, value.BlockHeight)
	assert.Nil(t, err)
	assert.NotNil(t, entry)

	assert.Equal(t, *entry, value)
}

func TestFetchUnknownStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	entry, err := db.FetchStateRootUpdateAggregation(1, 2)
	assert.NotNil(t, err)
	assert.Nil(t, entry)
}

func TestStoreAndFetchStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaTransactionId: tests.Keccak256(4),
		NearDaCommitment:    tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}

	msgModel, err := db.StoreStateRootUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreStateRootUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	entry, err := db.FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight)
	assert.Nil(t, err)
	assert.NotNil(t, entry)

	assert.Equal(t, *entry, value)
}

func TestStateRootUpdateAggregationReplace(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaTransactionId: tests.Keccak256(4),
		NearDaCommitment:    tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}

	msgModel, err := db.StoreStateRootUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreStateRootUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	var count int64
	db.DB().Model(&models.MessageBlsAggregation{}).Count(&count)
	assert.Equal(t, count, int64(1))
}

func TestFetchUnknownOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	entry, err := db.FetchOperatorSetUpdate(1)
	assert.NotNil(t, err)
	assert.Nil(t, entry)
}

func TestStoreAndFetchOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	value := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	_, err = db.StoreOperatorSetUpdate(value)
	assert.Nil(t, err)

	entry, err := db.FetchOperatorSetUpdate(value.Id)
	assert.Nil(t, err)
	assert.NotNil(t, entry)

	assert.Equal(t, *entry, value)
}

func TestFetchUnknownOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	entry, err := db.FetchOperatorSetUpdateAggregation(1)
	assert.NotNil(t, err)
	assert.Nil(t, entry)
}

func TestStoreAndFetchOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	msgModel, err := db.StoreOperatorSetUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreOperatorSetUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	entry, err := db.FetchOperatorSetUpdateAggregation(msg.Id)
	assert.Nil(t, err)
	assert.NotNil(t, entry)

	assert.Equal(t, *entry, value)
}

func TestOperatorSetUpdateAggregationReplace(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	msgModel, err := db.StoreOperatorSetUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreOperatorSetUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(msgModel, value)
	assert.Nil(t, err)

	var count int64
	db.DB().Model(&models.MessageBlsAggregation{}).Count(&count)
	assert.Equal(t, count, int64(1))
}

func TestFetchCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg1 := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         1,
		Timestamp:           0,
		NearDaTransactionId: tests.Keccak256(4),
		NearDaCommitment:    tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}

	msgDigest1, err := msg1.Digest()
	assert.Nil(t, err)

	aggregation1 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest1,
	}

	msg2 := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           1,
		NearDaTransactionId: tests.Keccak256(4),
		NearDaCommitment:    tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}

	msgDigest2, err := msg2.Digest()
	assert.Nil(t, err)

	aggregation2 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest2,
	}

	msg3 := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	msgDigest3, err := msg3.Digest()
	assert.Nil(t, err)

	aggregation3 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest3,
	}

	msg4 := messages.OperatorSetUpdateMessage{
		Id:        2,
		Timestamp: 3,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	msgDigest4, err := msg4.Digest()
	assert.Nil(t, err)

	aggregation4 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest4,
	}

	var stateRootUpdateMsgModel *models.StateRootUpdateMessage
	var operatorSetUpdateMsgModel *models.OperatorSetUpdateMessage

	stateRootUpdateMsgModel, err = db.StoreStateRootUpdate(msg1)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(stateRootUpdateMsgModel, aggregation1)
	assert.Nil(t, err)

	stateRootUpdateMsgModel, err = db.StoreStateRootUpdate(msg2)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(stateRootUpdateMsgModel, aggregation2)
	assert.Nil(t, err)

	operatorSetUpdateMsgModel, err = db.StoreOperatorSetUpdate(msg3)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(operatorSetUpdateMsgModel, aggregation3)
	assert.Nil(t, err)

	operatorSetUpdateMsgModel, err = db.StoreOperatorSetUpdate(msg4)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(operatorSetUpdateMsgModel, aggregation4)
	assert.Nil(t, err)

	result, err := db.FetchCheckpointMessages(0, 3)
	assert.Nil(t, err)
	assert.Equal(t, *result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg1, msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation1, aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3, aggregation4},
	})

	result, err = db.FetchCheckpointMessages(1, 3)
	assert.Nil(t, err)
	assert.Equal(t, *result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3, aggregation4},
	})

	result, err = db.FetchCheckpointMessages(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, *result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3},
	})

	result, err = db.FetchCheckpointMessages(4, 10)
	assert.Nil(t, err)
	assert.Equal(t, *result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{},
	})
}

func TestFetchCheckpointMessages_TimestampTooLarge(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	t.Run("fromTimestamp too large", func(t *testing.T) {
		_, err := db.FetchCheckpointMessages(uint64(0x8000000000000000), 0)
		assert.NotNil(t, err)
	})

	t.Run("toTimestamp too large", func(t *testing.T) {
		_, err := db.FetchCheckpointMessages(0, uint64(0x8000000000000000))
		assert.NotNil(t, err)
	})
}

func TestFetchCheckpointMessages_InvalidRange(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	_, err = db.FetchCheckpointMessages(101, 100)
	assert.NotNil(t, err)
}

func TestStoreStateRootUpdate_LargeMsgValues(t *testing.T) {
	t.Skip("Currently impossible to store all uint64 values in the DB")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase(":memory:")
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{
		RollupId:            math.MaxUint32,
		BlockHeight:         math.MaxUint64, // TODO: Cannot be stored, maximum possible value is `math.MaxInt64`
		Timestamp:           math.MaxUint64, // TODO: Cannot be stored, maximum possible value is `math.MaxInt64`
		NearDaTransactionId: [32]byte{0xFF},
		NearDaCommitment:    [32]byte{0xFF},
		StateRoot:           [32]byte{0xFF},
	}
	_, err = db.StoreStateRootUpdate(msg)
	assert.Nil(t, err)

	stored, err := db.FetchStateRootUpdate(math.MaxUint32, math.MaxUint64)
	assert.NotNil(t, stored)
	assert.Nil(t, err)

	assert.Equal(t, &msg, stored)
}
