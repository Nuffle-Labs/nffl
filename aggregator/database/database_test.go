package database_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NethermindEth/near-sffl/aggregator/database"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/tests"
)

func TestFetchUnknownStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry servicemanager.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	value := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   tests.Keccak256(4),
	}

	err = db.StoreStateRootUpdate(value)
	assert.Nil(t, err)

	var entry servicemanager.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(value.RollupId, value.BlockHeight, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry aggtypes.MessageBlsAggregationServiceResponse

	err = db.FetchStateRootUpdateAggregation(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   tests.Keccak256(4),
	}

	err = db.StoreStateRootUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := core.GetStateRootUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	err = db.StoreStateRootUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry aggtypes.MessageBlsAggregationServiceResponse

	err = db.FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry registryrollup.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(1, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	value := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}

	err = db.StoreOperatorSetUpdate(value)
	assert.Nil(t, err)

	var entry registryrollup.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(value.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	var entry aggtypes.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(1, &entry)
	assert.NotNil(t, err)
}

func TestStoreAndFetchOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}

	err = db.StoreOperatorSetUpdate(msg)
	assert.Nil(t, err)

	msgDigest, err := core.GetOperatorSetUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	err = db.StoreOperatorSetUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry aggtypes.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(msg.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := database.NewDatabase("")
	assert.Nil(t, err)

	msg1 := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 1,
		Timestamp:   0,
		StateRoot:   tests.Keccak256(4),
	}

	msgDigest1, err := core.GetStateRootUpdateMessageDigest(&msg1)
	assert.Nil(t, err)

	aggregation1 := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest1,
	}

	msg2 := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   1,
		StateRoot:   tests.Keccak256(4),
	}

	msgDigest2, err := core.GetStateRootUpdateMessageDigest(&msg2)
	assert.Nil(t, err)

	aggregation2 := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest2,
	}

	msg3 := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}

	msgDigest3, err := core.GetOperatorSetUpdateMessageDigest(&msg3)
	assert.Nil(t, err)

	aggregation3 := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest3,
	}

	msg4 := registryrollup.OperatorSetUpdateMessage{
		Id:        2,
		Timestamp: 3,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}

	msgDigest4, err := core.GetOperatorSetUpdateMessageDigest(&msg4)
	assert.Nil(t, err)

	aggregation4 := aggtypes.MessageBlsAggregationServiceResponse{
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

	var result coretypes.CheckpointMessages

	err = db.FetchCheckpointMessages(0, 3, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, coretypes.CheckpointMessages{
		StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{msg1, msg2},
		StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{aggregation1, aggregation2},
		OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{aggregation3, aggregation4},
	})

	err = db.FetchCheckpointMessages(1, 3, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, coretypes.CheckpointMessages{
		StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{aggregation2},
		OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{msg3, msg4},
		OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{aggregation3, aggregation4},
	})

	err = db.FetchCheckpointMessages(1, 2, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, coretypes.CheckpointMessages{
		StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{msg2},
		StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{aggregation2},
		OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{msg3},
		OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{aggregation3},
	})

	err = db.FetchCheckpointMessages(4, 10, &result)
	assert.Nil(t, err)
	assert.Equal(t, result, coretypes.CheckpointMessages{
		StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{},
		StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{},
		OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{},
		OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{},
	})
}
