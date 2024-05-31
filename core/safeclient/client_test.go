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
	"github.com/ethereum/go-ethereum/common"
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
	mockNetwork := &MockNetwork{
		blockTicker: time.NewTicker(1 * time.Second),
	}

	mockNetwork.blockTicker.Stop()

	go func() {
		mockNetwork.blockTicker.Reset(1 * time.Second)
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
	sub := m.SubscribeToBlocks()
	block := <-sub
	fmt.Println("paused at block", block)

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

func (m *MockNetwork) SubscribeToBlocks() <-chan uint64 {
	m.blockSubscribersLock.Lock()
	defer m.blockSubscribersLock.Unlock()

	ch := make(chan uint64, 1)
	m.blockSubscribers = append(m.blockSubscribers, ch)

	return ch
}

func (m *MockEthClient) ReopenConnection() {
	m.isClosed = false
}

func (m *MockEthClient) PauseHeaderSubscriptions() {
	m.isPaused = true
}

func (m *MockEthClient) ResumeHeaderSubscriptions() {
	m.isPaused = false
}

func (m *MockEthClient) subscribeToClose() <-chan bool {
	m.closeSubscribersLock.Lock()
	defer m.closeSubscribersLock.Unlock()

	ch := make(chan bool, 1)
	m.closeSubscribers = append(m.closeSubscribers, ch)

	return ch
}

func NewMockEthClientFromNetwork(ctx context.Context, mockCtrl *gomock.Controller, mockNetwork *MockNetwork) *MockEthClient {
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

			blockCh := mockNetwork.SubscribeToBlocks()
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

						blockCh = mockNetwork.SubscribeToBlocks()

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

			blockCh := mockNetwork.SubscribeToBlocks()
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

						blockCh = mockNetwork.SubscribeToBlocks()

						if mockClient.isClosed {
							continue
						}

						log := types.Log{BlockNumber: blockNum, Index: uint(blockNum)}

						ch <- log
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

func NewMockSafeClientFromNetwork(ctx context.Context, mockCtrl *gomock.Controller, logger logging.Logger, mockNetwork *MockNetwork) (*safeclient.SafeEthClient, error) {
	client, err := safeclient.NewSafeEthClient("", logger, safeclient.WithCustomCreateClient(func(rpcUrl string, logger logging.Logger) (eth.Client, error) {
		return NewMockEthClientFromNetwork(ctx, mockCtrl, mockNetwork), nil
	}))

	return client, err
}

func NewMockClientControllable(ctx context.Context, mockCtrl *gomock.Controller, headerProxyC <-chan *types.Header, logProxyC <-chan types.Log, blockNum *uint64) (mockClient *MockEthClient) {
	fmt.Println("creating mock client")

	mockClient = &MockEthClient{
		MockEthClient: mocks.NewMockEthClient(mockCtrl),
	}

	mockClient.EXPECT().SubscribeNewHead(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
			if mockClient.isClosed {
				return nil, errors.New("connection refused")
			}

			sub := mocks.NewMockSubscription(mockCtrl)

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
					case header := <-headerProxyC:
						if mockClient.isClosed {
							continue
						}

						if !mockClient.isPaused {
							ch <- header
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
					case log := <-logProxyC:
						if mockClient.isClosed {
							continue
						}

						if !mockClient.isPaused {
							ch <- log
						}
					}
				}
			}()

			return sub, nil
		},
	).AnyTimes()

	mockClient.EXPECT().BlockNumber(gomock.Any()).DoAndReturn(
		func(ctx context.Context) (uint64, error) {
			return *blockNum, nil
		},
	).AnyTimes()

	return mockClient
}

func NewMockSafeClientControllable(ctx context.Context, mockCtrl *gomock.Controller, logger logging.Logger, headerProxyC <-chan *types.Header, logProxyC <-chan types.Log, blockNum *uint64) (*safeclient.SafeEthClient, error) {
	client, err := safeclient.NewSafeEthClient("", logger, safeclient.WithCustomCreateClient(func(rpcUrl string, logger logging.Logger) (eth.Client, error) {
		mockClient := NewMockClientControllable(ctx, mockCtrl, headerProxyC, logProxyC, blockNum)
		return mockClient, nil
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

	client, err := NewMockSafeClientFromNetwork(ctx, mockCtrl, logger, mockNetwork)
	assert.NoError(t, err)

	defer client.Close()

	mockClient := client.Client.(*MockEthClient)

	headCh := make(chan *types.Header)
	flushHeadCh := func() {
		select {
		case <-headCh:
		case <-time.After(100 * time.Millisecond):
		}
	}

	_, err = client.SubscribeNewHead(ctx, headCh)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head.Number.Uint64())
	}

	mockClient.PauseHeaderSubscriptions()
	select {
	case <-headCh:
		t.Fatal("unexpected head")
	case <-time.After(2 * time.Second):
	}
	mockClient.ResumeHeaderSubscriptions()

	mockNetwork.PauseBlockProduction()
	block := mockNetwork.BlockNumber()
	flushHeadCh()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head.Number.Uint64())
	}

	mockClient.CloseConnection()
	time.Sleep(2 * time.Second)
	mockClient.ReopenConnection()

	mockNetwork.PauseBlockProduction()
	block = mockNetwork.BlockNumber()
	flushHeadCh()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head.Number.Uint64())
	}

	mockClient.CloseConnection()
	time.Sleep(2 * time.Second)
	mockClient.ReopenConnection()

	mockNetwork.PauseBlockProduction()
	block = mockNetwork.BlockNumber()
	flushHeadCh()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		head := <-headCh
		assert.Equal(t, uint64(i), head.Number.Uint64())

		fmt.Println("head", head.Number.Uint64())
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

	client, err := NewMockSafeClientFromNetwork(ctx, mockCtrl, logger, mockNetwork)
	assert.NoError(t, err)

	defer client.Close()

	mockClient := client.Client.(*MockEthClient)

	logCh := make(chan types.Log)
	flushLogCh := func() {
		select {
		case <-logCh:
		case <-time.After(100 * time.Millisecond):
		}
	}

	_, err = client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, logCh)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log.BlockNumber)
	}

	mockClient.CloseConnection()
	time.Sleep(2 * time.Second)
	mockClient.ReopenConnection()

	mockNetwork.PauseBlockProduction()
	block := mockNetwork.BlockNumber()
	flushLogCh()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log.BlockNumber)
	}

	mockClient.CloseConnection()
	time.Sleep(2 * time.Second)
	mockClient.ReopenConnection()

	mockNetwork.PauseBlockProduction()
	block = mockNetwork.BlockNumber()
	flushLogCh()
	mockNetwork.ResumeBlockProduction()

	for i := block + 1; i <= block+3; i++ {
		log := <-logCh
		assert.Equal(t, uint64(i), log.BlockNumber)

		fmt.Println("log", log.BlockNumber)
	}
}

func TestLogCache(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logging.NewZapLogger("development")
	assert.NoError(t, err)

	blockNum := uint64(0)
	headerProxyC := make(chan *types.Header)
	logProxyC := make(chan types.Log)

	client, err := NewMockSafeClientControllable(ctx, mockCtrl, logger, headerProxyC, logProxyC, &blockNum)
	assert.NoError(t, err)

	defer client.Close()

	logCh := make(chan types.Log, 10)
	_, err = client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, logCh)
	assert.NoError(t, err)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
				fmt.Println("sending log")
				logProxyC <- types.Log{BlockNumber: 1, BlockHash: common.Hash{1}}
			}
		}
	}()

	time.Sleep(2 * time.Second)

	assert.Equal(t, 1, len(logCh))

	logProxyC <- types.Log{BlockNumber: 2, BlockHash: common.Hash{2}}

	time.Sleep(2 * time.Second)

	assert.Equal(t, 2, len(logCh))

	log := <-logCh
	assert.Equal(t, uint64(1), log.BlockNumber)
	log = <-logCh
	assert.Equal(t, uint64(2), log.BlockNumber)
}

func TestSubscribeFilterLogs_Unsubscribe(t *testing.T) {
	logger, err := logging.NewZapLogger("development")
	assert.NoError(t, err)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockEthClient(mockCtrl)
	mockClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(1_000), nil)
	mockClient.EXPECT().SubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
			errChann := make(chan error)

			sub := mocks.NewMockSubscription(mockCtrl)
			sub.EXPECT().Unsubscribe().Do(func() { close(errChann) })
			sub.EXPECT().Err().Return(errChann)

			return sub, nil
		},
	)

	client, err := safeclient.NewSafeEthClient("", logger, safeclient.WithCustomCreateClient(func(string, logging.Logger) (eth.Client, error) { return mockClient, nil }))
	assert.NoError(t, err)
	assert.NotNil(t, client)
	defer client.Close()

	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(900),
		ToBlock:   big.NewInt(1_100),
	}
	logCh := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), filterQuery, logCh)
	assert.NoError(t, err)
	assert.NotNil(t, sub)
}
