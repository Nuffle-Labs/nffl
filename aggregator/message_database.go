package aggregator

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"

	badger "github.com/dgraph-io/badger/v4"
)

type MessageDatabase struct {
	db     *badger.DB
	dbPath string
	mu     sync.RWMutex
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

func (messageDb *MessageDatabase) Close() error {
	return messageDb.db.Close()
}

func (messageDb *MessageDatabase) Store(prefix string, key string, value any) error {
	messageDb.mu.Lock()
	defer messageDb.mu.Unlock()

	fullKey := prefix + key

	err := messageDb.db.Update(func(txn *badger.Txn) error {
		value, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return txn.Set([]byte(fullKey), value)
	})

	return err
}

func (messageDb *MessageDatabase) Fetch(prefix string, key string, value any) error {
	fullKey := prefix + key

	err := messageDb.db.View(func(txn *badger.Txn) error {
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

func (messageDb *MessageDatabase) StoreStateRootUpdate(stateRootUpdateMessage servicemanager.StateRootUpdateMessage) error {
	return messageDb.Store("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), stateRootUpdateMessage)
}

func (messageDb *MessageDatabase) FetchStateRootUpdate(rollupId uint32, blockHeight uint64, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) error {
	return messageDb.Fetch("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), stateRootUpdateMessage)
}

func (messageDb *MessageDatabase) StoreStateRootUpdateAggregation(stateRootUpdateMessage servicemanager.StateRootUpdateMessage, aggregation types.MessageBlsAggregationServiceResponse) error {
	return messageDb.Store("stateRootUpdates", fmt.Sprintf("%d_%d", stateRootUpdateMessage.RollupId, stateRootUpdateMessage.BlockHeight), aggregation)
}

func (messageDb *MessageDatabase) FetchStateRootUpdateAggregation(rollupId uint32, blockHeight uint64, aggregation *types.MessageBlsAggregationServiceResponse) error {
	return messageDb.Fetch("stateRootUpdates", fmt.Sprintf("%d_%d", rollupId, blockHeight), aggregation)
}
