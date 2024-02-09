package aggregator

import (
	"encoding/json"
	"fmt"
	"sync"

	badger "github.com/dgraph-io/badger/v4"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
)

type MessageDatabase struct {
	db     *badger.DB
	dbPath string
	lock   sync.RWMutex
}

func NewMessageDatabase(dbPath string) (*MessageDatabase, error) {
	opt := badger.DefaultOptions(dbPath)

	if dbPath == "" {
		opt = opt.WithInMemory(true)
	}

	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return &MessageDatabase{
		db:     db,
		dbPath: dbPath,
	}, nil
}

func (md *MessageDatabase) Close() error {
	return md.db.Close()
}

func (md *MessageDatabase) Store(prefix string, key string, value any) error {
	md.lock.Lock()
	defer md.lock.Unlock()

	fullKey := prefix + key

	err := md.db.Update(func(txn *badger.Txn) error {
		value, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(fullKey), value)
	})

	return err
}

func (md *MessageDatabase) Fetch(prefix string, key string, value any) error {
	fullKey := prefix + key

	err := md.db.View(func(txn *badger.Txn) error {
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

func (md *MessageDatabase) StoreStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error {
	return md.Store("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), stateRootUpdateMessage)
}

func (md *MessageDatabase) FetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error {
	return md.Fetch("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), stateRootUpdateMessage)
}

func (md *MessageDatabase) StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	return md.Store("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), aggregation)
}

func (md *MessageDatabase) FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *types.MessageBlsAggregationServiceResponse) error {
	return md.Fetch("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), aggregation)
}

func (md *MessageDatabase) StoreOperatorSetUpdate(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage) error {
	return md.Store("operatorSetUpdates", fmt.Sprintf("%d", operatorSetUpdateMessage.Id), operatorSetUpdateMessage)
}

func (md *MessageDatabase) FetchOperatorSetUpdate(id uint32, operatorSetUpdateMessage *registryrollup.OperatorSetUpdateMessage) error {
	return md.Fetch("operatorSetUpdates", fmt.Sprintf("%d", id), operatorSetUpdateMessage)
}

func (md *MessageDatabase) StoreOperatorSetUpdateAggregation(operatorSetUpdateMessage registryrollup.OperatorSetUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	return md.Store("operatorSetUpdates", fmt.Sprintf("%d", operatorSetUpdateMessage.Id), aggregation)
}

func (md *MessageDatabase) FetchOperatorSetUpdateAggregation(id uint32, aggregation *types.MessageBlsAggregationServiceResponse) error {
	return md.Fetch("operatorSetUpdates", fmt.Sprintf("%d", id), aggregation)
}
