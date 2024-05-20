package safeclient_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/core/safeclient"
)

type MockNetwork struct {
	blockTicker          *time.Ticker
	blockNum             uint64
	blockSubscribers     []chan<- uint64
	blockSubscribersLock sync.Mutex
}

func NewMockNetwork(ctx context.Context, mockCtrl *gomock.Controller) *MockNetwork {
	mockNetwork := &MockNetwork{}

	go func() {
		mockNetwork.blockTicker = time.NewTicker(1 * time.Second)
		defer mockNetwork.blockTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-mockNetwork.blockTicker.C:
				mockNetwork.blockNum++

				for _, ch := range mockNetwork.blockSubscribers {
					ch <- mockNetwork.blockNum
				}

				mockNetwork.blockSubscribersLock.Lock()
				mockNetwork.blockSubscribers = nil
				mockNetwork.blockSubscribersLock.Unlock()
			}
		}
	}()

	return mockNetwork
}

func (m *MockNetwork) PauseBlockProduction() {
	m.blockTicker.Stop()
}

func (m *MockNetwork) ResumeBlockProduction() {
	m.blockTicker.Reset(1 * time.Second)
}

func (m *MockNetwork) BlockNumber() uint64 {
	return m.blockNum
}

type MockEthClient struct {
	*mocks.MockEthClient
	isClosed             bool
	isPaused             bool
	closeSubscribers     []chan<- bool
	closeSubscribersLock sync.Mutex
}

func (m *MockEthClient) CloseConnection() {
	m.isClosed = true

	fmt.Println("closing")

	m.closeSubscribersLock.Lock()
	defer m.closeSubscribersLock.Unlock()

	fmt.Println("closing subscribers")

	for _, ch := range m.closeSubscribers {
		ch <- true
	}

	fmt.Println("closing done")

	m.closeSubscribers = nil
}

func (m *MockEthClient) PauseHeaderSubscriptions() {
	m.isPaused = true
}

func (m *MockEthClient) ResumeHeaderSubscriptions() {
	m.isPaused = false
}

func (m *MockNetwork) subscribeToBlocks() <-chan uint64 {
	m.blockSubscribersLock.Lock()
	defer m.blockSubscribersLock.Unlock()

	ch := make(chan uint64, 1)
	m.blockSubscribers = append(m.blockSubscribers, ch)

	return ch
}

func (m *MockEthClient) subscribeToClose() <-chan bool {
	m.closeSubscribersLock.Lock()
	defer m.closeSubscribersLock.Unlock()

	ch := make(chan bool, 1)
	m.closeSubscribers = append(m.closeSubscribers, ch)

	return ch
}

func NewMockEthClient(ctx context.Context, mockCtrl *gomock.Controller, mockNetwork *MockNetwork) *MockEthClient {
	fmt.Println("creating mock client")

	mockClient := &MockEthClient{
		MockEthClient: mocks.NewMockEthClient(mockCtrl),
	}

	mockClient.EXPECT().SubscribeNewHead(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
			if mockClient.isClosed {
				return nil, errors.New("connection refused")
			}

			sub := mocks.NewMockSubscription(mockCtrl)

			blockCh := mockNetwork.subscribeToBlocks()
			closeCh := mockClient.subscribeToClose()

			subErrCh := make(chan error)
			stopCh := make(chan struct{})

			sub.EXPECT().Err().Return(subErrCh).AnyTimes()
			sub.EXPECT().Unsubscribe().Do(func() {
				close(stopCh)
			}).AnyTimes()

			go func() {
				for {
					select {
					case <-stopCh:
						return
					case <-ctx.Done():
						return
					case closed := <-closeCh:
						fmt.Println("closed", closed)

						closeCh = mockClient.subscribeToClose()
						if closed {
							subErrCh <- errors.New("connection refused")
						}
					case blockNum := <-blockCh:
						fmt.Println("header block", blockNum, "closed", mockClient.isClosed, "paused", mockClient.isPaused)

						blockCh = mockNetwork.subscribeToBlocks()

						if mockClient.isClosed {
							continue
						}

						if !mockClient.isPaused {
							ch <- &types.Header{Number: big.NewInt(int64(blockNum))}
						}
					}
				}
			}()
			return sub, nil
		},
	).AnyTimes()

	mockClient.EXPECT().SubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
			if mockClient.isClosed {
				return nil, errors.New("connection refused")
			}

			sub := mocks.NewMockSubscription(mockCtrl)

			blockCh := mockNetwork.subscribeToBlocks()
			closeCh := mockClient.subscribeToClose()

			subErrCh := make(chan error)
			stopCh := make(chan struct{})

			sub.EXPECT().Err().Return(subErrCh).AnyTimes()
			sub.EXPECT().Unsubscribe().Do(func() {
				close(stopCh)
			}).AnyTimes()

			go func() {
				for {
					select {
					case <-stopCh:
						fmt.Println("subscription done")
						return
					case <-ctx.Done():
						fmt.Println("subscription done")
						return
					case closed := <-closeCh:
						fmt.Println("closed", closed)

						closeCh = mockClient.subscribeToClose()

						if closed {
							subErrCh <- errors.New("connection refused")
						}
					case blockNum := <-blockCh:
						fmt.Println("log block", blockNum)

						blockCh = mockNetwork.subscribeToBlocks()

						if mockClient.isClosed {
							continue
						}

						ch <- types.Log{BlockNumber: blockNum}
					}
				}
			}()
			return sub, nil
		},
	).AnyTimes()

	mockClient.EXPECT().BlockNumber(gomock.Any()).DoAndReturn(
		func(ctx context.Context) (uint64, error) {
			return mockNetwork.blockNum, nil
		},
	).AnyTimes()

	return mockClient
}

func NewMockSafeClient(ctx context.Context, mockCtrl *gomock.Controller, logger logging.Logger, mockNetwork *MockNetwork) (*safeclient.SafeEthClient, error) {
	client, err := safeclient.NewSafeEthClient("", logger, safeclient.WithCustomCreateClient(func(rpcUrl string, logger logging.Logger) (eth.Client, error) {
		return NewMockEthClient(ctx, mockCtrl, mockNetwork), nil
	}))

	return client, err
}

func TestSubscribeNewHead(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logging.NewZapLogger("development")
	assert.NoError(t, err)

	mockNetwork := NewMockNetwork(ctx, mockCtrl)

	client, err := NewMockSafeClient(ctx, mockCtrl, logger, mockNetwork)
	assert.NoError(t, err)

	headCh := make(chan *types.Header)
	_, err = client.SubscribeNewHead(ctx, headCh)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head)
	}

	mockClient := client.Client.(*MockEthClient)
	mockClient.PauseHeaderSubscriptions()
	select {
	case <-headCh:
		t.Fatal("unexpected head")
	case <-time.After(2 * time.Second):
	}
	mockClient.ResumeHeaderSubscriptions()

	mockNetwork.PauseBlockProduction()
	block := mockNetwork.BlockNumber()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head)
	}

	mockClient.CloseConnection()

	mockNetwork.PauseBlockProduction()
	block = mockNetwork.BlockNumber()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head)
	}

	mockClient = client.Client.(*MockEthClient)
	mockClient.CloseConnection()

	mockNetwork.PauseBlockProduction()
	block = mockNetwork.BlockNumber()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head)
	}
}

func TestSubscribeFilterLogs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logging.NewZapLogger("development")
	assert.NoError(t, err)

	mockNetwork := NewMockNetwork(ctx, mockCtrl)

	client, err := NewMockSafeClient(ctx, mockCtrl, logger, mockNetwork)
	assert.NoError(t, err)

	logCh := make(chan types.Log, 10)
	_, err = client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, logCh)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log)
	}

	mockClient := client.Client.(*MockEthClient)
	mockClient.CloseConnection()

	for i := 4; i <= 6; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log)
	}

	mockClient = client.Client.(*MockEthClient)
	mockClient.CloseConnection()

	for i := 7; i <= 9; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log)
	}
}
