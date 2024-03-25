package database_test

import (
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NethermindEth/near-sffl/aggregator/database"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/tests"
)

func TestFetchUnknownStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry messages.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	value := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   tests.Keccak256(4),
	}

	err = db.StoreStateRootUpdate(value)
	assert.Nil(t, err)

	var entry messages.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(value.RollupId, value.BlockHeight, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry messages.MessageBlsAggregation

	err = db.FetchStateRootUpdateAggregation(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   tests.Keccak256(4),
	}

	err = db.StoreStateRootUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreStateRootUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry messages.MessageBlsAggregation

	err = db.FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry messages.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(1, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	value := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	err = db.StoreOperatorSetUpdate(value)
	assert.Nil(t, err)

	var entry messages.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(value.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry messages.MessageBlsAggregation

	err = db.FetchOperatorSetUpdateAggregation(1, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []coretypes.RollupOperator{
			{Pubkey: bls.NewG1Point(big.NewInt(3), big.NewInt(4)), Weight: big.NewInt(5)},
		},
	}

	err = db.StoreOperatorSetUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	value := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	err = db.StoreOperatorSetUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry messages.MessageBlsAggregation

	err = db.FetchOperatorSetUpdateAggregation(msg.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg1 := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 1,
		Timestamp:   0,
		StateRoot:   tests.Keccak256(4),
	}

	msgDigest1, err := msg1.Digest()
	assert.Nil(t, err)

	aggregation1 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest1,
	}

	msg2 := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   1,
		StateRoot:   tests.Keccak256(4),
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

	err = db.StoreStateRootUpdate(msg1)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(msg1, aggregation1)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdate(msg2)
	assert.Nil(t, err)

	err = db.StoreStateRootUpdateAggregation(msg2, aggregation2)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdate(msg3)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(msg3, aggregation3)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdate(msg4)
	assert.Nil(t, err)

	err = db.StoreOperatorSetUpdateAggregation(msg4, aggregation4)
	assert.Nil(t, err)

	var result messages.CheckpointMessages

	err = db.FetchCheckpointMessages(0, 3, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg1, msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation1, aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3, aggregation4},
	})

	err = db.FetchCheckpointMessages(1, 3, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3, aggregation4},
	})

	err = db.FetchCheckpointMessages(1, 2, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation2},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg3},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation3},
	})

	err = db.FetchCheckpointMessages(4, 10, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{},
	})
}
