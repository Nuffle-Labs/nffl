package attestor

import (
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/NethermindEth/near-sffl/operator/consumer"
)

const NUM_OF_BLOCKS = 1000

func generateBlockData() consumer.BlockData {
	rand.Seed(time.Now().UnixNano())

	randomRollupId := uint32(rand.Intn(NUM_OF_BLOCKS / 100))
	randomBlockNumber := big.NewInt(int64(rand.Intn(100000)))
	header := types.Header{
		Number: randomBlockNumber,
	}

	return consumer.BlockData{
		RollupId: randomRollupId,
		Block:    *types.NewBlockWithHeader(&header),
	}
}

func subscribe(notifier *Notifier, blocks []consumer.BlockData, subscribedWg *sync.WaitGroup, unsubscribedWg *sync.WaitGroup) {
	for i := 0; i < len(blocks); i++ {
		subscribedWg.Add(1)
		unsubscribedWg.Add(1)

		go func(block consumer.BlockData, notifier *Notifier) {
			predicate := func(mqBlock consumer.BlockData) bool {
				if block.Block.Header().Number.Cmp(mqBlock.Block.Header().Number) != 0 {
					return false
				}

				return true
			}

			blocksC, id := notifier.Subscribe(block.RollupId, predicate)
			subscribedWg.Done()

			defer func() {
				notifier.Unsubscribe(block.RollupId, id)
				unsubscribedWg.Done()
			}()

			for {
				mqBlock := <-blocksC
				if block.Block.Header().Number.Cmp(mqBlock.Block.Header().Number) != 0 {
					continue
				}

				return
			}
		}(blocks[i], notifier)
	}
}

func notify(t *testing.T, notifier *Notifier, blocks []consumer.BlockData) {
	for i := 0; i < len(blocks); i++ {
		err := notifier.Notify(blocks[i].RollupId, blocks[i])
		assert.Nil(t, err)

		// Timeout necessary to let runtime switch context
		time.Sleep(time.Microsecond * 50)
	}
}

func TestNotifier(t *testing.T) {
	notifier := NewNotifier()
	blocks := make([]consumer.BlockData, NUM_OF_BLOCKS)
	for i := 0; i < NUM_OF_BLOCKS; i++ {
		blocks[i] = generateBlockData()
	}

	var subscribedWg sync.WaitGroup
	var unsubscribedWg sync.WaitGroup
	subscribe(&notifier, blocks, &subscribedWg, &unsubscribedWg)
	subscribedWg.Wait()

	notify(t, &notifier, blocks)
	unsubscribedWg.Wait()

	for _, val := range notifier.rollupIdsToSubscribers {
		assert.Equal(t, 0, val.Len())
	}
}

func TestNotifierNotifyUnknownRollup(t *testing.T) {
	block := generateBlockData()
	notifier := NewNotifier()

	err := notifier.Notify(block.RollupId, block)
	assert.Error(t, err, unknownRollupIdError)
}

func TestNotifierSubscribeAndUnsubscribe(t *testing.T) {
	block := generateBlockData()
	notifier := NewNotifier()

	predicate := func(mqBlock consumer.BlockData) bool {
		if block.Block.Header().Number.Cmp(mqBlock.Block.Header().Number) != 0 {
			return false
		}

		return true
	}
	_, id := notifier.Subscribe(block.RollupId, predicate)
	assert.Equal(t, notifier.rollupIdsToSubscribers[block.RollupId].Len(), 1)

	notifier.Unsubscribe(block.RollupId, id)
	assert.Equal(t, notifier.rollupIdsToSubscribers[block.RollupId].Len(), 0)
}
