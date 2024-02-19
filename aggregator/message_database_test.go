package aggregator

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
)

func TestFetchUnknown(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry string

	err = db.Fetch("prefix", "key", &entry)
	assert.NotNil(t, err)
}

func TestFetchKnown(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	db.db.Update(func(txn *badger.Txn) error {
		entry, err := json.Marshal("value")
		if err != nil {
			return err
		}

		return txn.Set([]byte("prefixkey"), entry)
	})

	var entry string

	err = db.Fetch("prefix", "key", &entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, "value")
}

func TestStoreAndFetch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry = "value"

	err = db.Store("prefix", "key", entry)
	assert.Nil(t, err)

	var fetchedEntry string

	err = db.Fetch("prefix", "key", &fetchedEntry)
	assert.Nil(t, err)

	assert.Equal(t, fetchedEntry, entry)
}

func TestFetchUnknownStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry servicemanager.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestFetchKnownStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	value := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
	}
	prefix := "stateRootUpdates"
	key := fmt.Sprintf("%d_%d", value.RollupId, value.BlockHeight)

	db.db.Update(func(txn *badger.Txn) error {
		entry, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(prefix+key), entry)
	})

	var entry servicemanager.StateRootUpdateMessage

	err = db.FetchStateRootUpdate(value.RollupId, value.BlockHeight, &entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, value)
}

func TestStoreAndFetchStateRootUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	value := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
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

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchStateRootUpdateAggregation(1, 2, &entry)
	assert.NotNil(t, err)
}

func TestFetchKnownStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	msg := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
	}
	msgDigest, err := core.GetStateRootUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := types.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}
	prefix := "stateRootUpdateAggregations"
	key := fmt.Sprintf("%d_%d", msg.RollupId, msg.BlockHeight)

	db.db.Update(func(txn *badger.Txn) error {
		entry, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(prefix+key), entry)
	})

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, &entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, value)
}

func TestStoreAndFetchStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	msg := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
	}
	msgDigest, err := core.GetStateRootUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := types.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	err = db.StoreStateRootUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}

func TestFetchUnknownOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry registryrollup.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(1, &entry)
	assert.NotNil(t, err)
}

func TestFetchKnownOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	value := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}
	prefix := "operatorSetUpdates"
	key := fmt.Sprintf("%d", value.Id)

	db.db.Update(func(txn *badger.Txn) error {
		entry, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(prefix+key), entry)
	})

	var entry registryrollup.OperatorSetUpdateMessage

	err = db.FetchOperatorSetUpdate(value.Id, &entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, value)
}

func TestStoreAndFetchOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
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

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(1, &entry)
	assert.NotNil(t, err)
}

func TestFetchKnownOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	msg := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}
	msgDigest, err := core.GetOperatorSetUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := types.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}
	prefix := "operatorSetUpdateAggregations"
	key := fmt.Sprintf("%d", msg.Id)

	db.db.Update(func(txn *badger.Txn) error {
		entry, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(prefix+key), entry)
	})

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(msg.Id, &entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, value)
}

func TestStoreAndFetchOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, err := NewMessageDatabase("")
	assert.Nil(t, err)

	msg := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
		Operators: []registryrollup.OperatorsOperator{
			{Pubkey: registryrollup.BN254G1Point{X: big.NewInt(3), Y: big.NewInt(4)}, Weight: big.NewInt(5)},
		},
	}
	msgDigest, err := core.GetOperatorSetUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	value := types.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	err = db.StoreOperatorSetUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(msg.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}
