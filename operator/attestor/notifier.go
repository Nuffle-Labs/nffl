package attestor

import (
	"container/list"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"sync"
)

// Notifier Broadcasts block from some rollup
// to subscribers
type Notifier struct {
	rollupIdsToSubscribers map[uint32]*list.List
	notifierLock           sync.Mutex
}

func NewNotifier() Notifier {
	return Notifier{
		rollupIdsToSubscribers: make(map[uint32]*list.List),
	}
}

func (notifier *Notifier) Subscribe(rollupId uint32) (chan consumer.BlockData, *list.Element) {
	notifier.notifierLock.Lock()
	defer notifier.notifierLock.Lock()

	if _, exists := notifier.rollupIdsToSubscribers[rollupId]; !exists {
		notifier.rollupIdsToSubscribers[rollupId] = list.New()
	}

	notifierC := make(chan consumer.BlockData)
	id := notifier.rollupIdsToSubscribers[rollupId].PushBack(notifierC)

	return notifierC, id
}

func (notifier *Notifier) Notify(rollupId uint32, block consumer.BlockData) error {
	notifier.notifierLock.Lock()
	defer notifier.notifierLock.Lock()

	subscribers, exists := notifier.rollupIdsToSubscribers[rollupId]
	if !exists {
		return unknownRollupIdError
	}

	for el := subscribers.Front(); el != nil; el = el.Next() {
		subscriber, ok := el.Value.(chan consumer.BlockData)
		if !ok {
			panic("Notifier: unreachable")
		}

		subscriber <- block
	}

	return nil
}

func (notifier *Notifier) Unsubscribe(rollupId uint32, el *list.Element) {
	notifier.notifierLock.Lock()
	defer notifier.notifierLock.Lock()

	notifier.rollupIdsToSubscribers[rollupId].Remove(el)
}
