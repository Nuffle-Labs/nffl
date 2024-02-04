package aggregator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"

	badger "github.com/dgraph-io/badger/v4"
)

func (agg *Aggregator) startDatabase(ctx context.Context) error {
	db, err := badger.Open(badger.DefaultOptions(agg.databasePath))
	if err != nil {
		return err
	}
	defer db.Close()

	agg.database = db

	return nil
}

func (agg *Aggregator) databaseStore(prefix string, key string, value any) error {
	agg.databaseMu.Lock()
	defer agg.databaseMu.Unlock()

	fullKey := prefix + key

	err := agg.database.Update(func(txn *badger.Txn) error {
		value, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(fullKey), value)
	})

	return err
}

func (agg *Aggregator) databaseRead(prefix string, key string, value any) error {
	fullKey := prefix + key

	err := agg.database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fullKey))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &value)
		})
	})

	return err
}

func (agg *Aggregator) storeStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error {
	return agg.databaseStore("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), stateRootUpdateMessage)
}

func (agg *Aggregator) fetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error {
	return agg.databaseRead("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), stateRootUpdateMessage)
}

func (agg *Aggregator) storeStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	return agg.databaseStore("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), aggregation)
}

func (agg *Aggregator) fetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *types.MessageBlsAggregationServiceResponse) error {
	return agg.databaseRead("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), aggregation)
}
