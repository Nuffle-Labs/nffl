package aggregator

import (
	"encoding/json"
	"fmt"
	"testing"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
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

	value := servicemanager.StateRootUpdateMessage{}
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

	value := servicemanager.StateRootUpdateMessage{}

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

	msg := servicemanager.StateRootUpdateMessage{}
	value := types.MessageBlsAggregationServiceResponse{}
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

	msg := servicemanager.StateRootUpdateMessage{}
	value := types.MessageBlsAggregationServiceResponse{}

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

	value := registryrollup.OperatorSetUpdateMessage{}
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

	value := registryrollup.OperatorSetUpdateMessage{}

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

	msg := registryrollup.OperatorSetUpdateMessage{}
	value := types.MessageBlsAggregationServiceResponse{}
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

	msg := registryrollup.OperatorSetUpdateMessage{}
	value := types.MessageBlsAggregationServiceResponse{}

	err = db.StoreOperatorSetUpdateAggregation(msg, value)
	assert.Nil(t, err)

	var entry types.MessageBlsAggregationServiceResponse

	err = db.FetchOperatorSetUpdateAggregation(msg.Id, &entry)
	assert.Nil(t, err)

	assert.Equal(t, entry, entry)
}
