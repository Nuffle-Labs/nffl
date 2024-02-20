package attestor

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/NethermindEth/near-sffl/operator/consumer"
)

const NUM_OF_BLOCKS = 1000

func generateBlockData() consumer.BlockData {
	rand.Seed(time.Now().UnixNano())

	randomRollupId := uint32(rand.Intn(NUM_OF_BLOCKS / 100))
	randomBlockNumber := big.NewInt(int64(rand.Intn(100000))) // Random Block Number between 0 and 99999
	header := types.Header{
		Number: randomBlockNumber,
	}

	return consumer.BlockData{
		RollupId: randomRollupId,
		Block:    *types.NewBlockWithHeader(&header),
	}
}

func subscribe(notifier *Notifier, blocks []consumer.BlockData) {
	for i := 0; i < len(blocks); i++ {
		go func(block consumer.BlockData, notifier *Notifier) {
			blocksC, id := notifier.Subscribe(block.RollupId)
			defer notifier.Unsubscribe(block.RollupId, id)

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

		// Timeout necessary to let runtime switch cotext
		time.Sleep(time.Microsecond * 50)
	}
}

func TestNotifier(t *testing.T) {
	notifier := NewNotifier()
	blocks := make([]consumer.BlockData, NUM_OF_BLOCKS)
	for i := 0; i < NUM_OF_BLOCKS; i++ {
		blocks[i] = generateBlockData()
	}

	subscribe(&notifier, blocks)
	time.Sleep(time.Second)

	notify(t, &notifier, blocks)
	time.Sleep(time.Second)

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

	_, id := notifier.Subscribe(block.RollupId)
	assert.Equal(t, notifier.rollupIdsToSubscribers[block.RollupId].Len(), 1)

	notifier.Unsubscribe(block.RollupId, id)
	assert.Equal(t, notifier.rollupIdsToSubscribers[block.RollupId].Len(), 0)
}
