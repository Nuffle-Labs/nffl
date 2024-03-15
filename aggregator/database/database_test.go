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
