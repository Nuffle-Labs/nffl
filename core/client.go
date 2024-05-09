package core

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BLOCK_CHUNK_SIZE = 1000
	BLOCK_MAX_RANGE  = 10000
	RESUB_INTERVAL   = 5 * time.Minute
	REINIT_INTERVAL  = 1 * time.Minute
)

type SafeEthClient struct {
	eth.Client

	clientLock            sync.RWMutex
	reinitSubscribersLock sync.Mutex
	wg                    sync.WaitGroup
	logger                logging.Logger
	rpcUrl                string
	reinitInterval        time.Duration
	reinitSubscribers     []chan bool
	reinitC               chan struct{}
	closeC                chan struct{}
	collector             *rpccalls.Collector
}

func NewSafeEthClient(rpcUrl string, logger logging.Logger, opts ...SafeEthClientOption) (*SafeEthClient, error) {
	safeClient := &SafeEthClient{
		logger:         logger,
		rpcUrl:         rpcUrl,
		reinitInterval: REINIT_INTERVAL,
		reinitC:        make(chan struct{}),
		closeC:         make(chan struct{}),
	}

	for _, opt := range opts {
		opt(safeClient)
	}

	client, err := safeClient.createClient()
	if err != nil {
		return nil, err
	}
	safeClient.Client = client

	safeClient.wg.Add(1)
	go safeClient.handleReinit()

	return safeClient, nil
}

type SafeEthClientOption func(*SafeEthClient)

func WithReinitInterval(interval time.Duration) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.reinitInterval = interval
	}
}

func WithCollector(collector *rpccalls.Collector) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.collector = collector
	}
}

func (c *SafeEthClient) createClient() (eth.Client, error) {
	if c.collector == nil {
		client, err := eth.NewClient(c.rpcUrl)
		if err != nil {
			return nil, err
		}
		return client, nil
	} else {
		client, err := eth.NewInstrumentedClient(c.rpcUrl, c.collector)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
}

func (c *SafeEthClient) reinit() bool {
	c.clientLock.Lock()
	defer c.clientLock.Unlock()

	client, err := c.createClient()
	if err != nil {
		c.logger.Errorf("Failed to reinitialize client: %v", err)

		return false
	}

	c.Client = client
	return true
}

func (c *SafeEthClient) notifySubscribers(success bool) {
	c.reinitSubscribersLock.Lock()
	defer c.reinitSubscribersLock.Unlock()

	for _, ch := range c.reinitSubscribers {
		ch <- success
		close(ch)
	}
	c.reinitSubscribers = nil
}

func (c *SafeEthClient) handleReinit() {
	defer c.wg.Done()

	reinitTicker := time.NewTicker(c.reinitInterval)
	defer reinitTicker.Stop()

	isReinitializing := false
	handleEvent := func() {
		isReinitializing = true

		go func() {
			success := c.reinit()
			if success {
				isReinitializing = false
				reinitTicker.Stop()
			} else {
				reinitTicker.Reset(c.reinitInterval)
			}

			c.notifySubscribers(success)
		}()
	}

	for {
		select {
		case <-c.closeC:
			return
		case <-c.reinitC:
			if isReinitializing {
				continue
			}

			handleEvent()
		case <-reinitTicker.C:
			handleEvent()
		}
	}
}

func (c *SafeEthClient) WatchReinit() <-chan bool {
	ch := make(chan bool, 1)
	c.reinitSubscribersLock.Lock()
	c.reinitSubscribers = append(c.reinitSubscribers, ch)
	c.reinitSubscribersLock.Unlock()
	return ch
}

type SafeSubscription struct {
	sub          ethereum.Subscription
	lock         sync.Mutex
	errC         chan error
	unsubscribed bool
}

func NewSafeSubscription(sub ethereum.Subscription) *SafeSubscription {
	return &SafeSubscription{
		sub:  sub,
		errC: make(chan error, 1),
	}
}

func (s *SafeSubscription) Err() <-chan error {
	return s.errC
}

func (s *SafeSubscription) Unsubscribe() {
	s.lock.Lock()

	if s.unsubscribed {
		s.lock.Unlock()
		return
	}

	s.sub.Unsubscribe()
	s.unsubscribed = true
	s.lock.Unlock()

	<-s.errC
}

func (s *SafeSubscription) SetUnderlyingSub(sub ethereum.Subscription) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.sub.Unsubscribe()
	s.sub = sub
}

func (c *SafeEthClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	currentBlock, err := c.Client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	ch2 := make(chan types.Log)

	sub, err := c.Client.SubscribeFilterLogs(ctx, q, ch2)
	if err != nil {
		c.logger.Errorf("Failed to resubscribe: %v", err)
		return nil, err
	}

	safeSub := NewSafeSubscription(sub)
	lastBlock := currentBlock

	resubFilterLogs := func() error {
		currentBlock, err := c.Client.BlockNumber(ctx)
		if err != nil {
			c.logger.Errorf("Failed to get current block number: %v", err)
			return err
		}

		fromBlock := max(lastBlock+1, currentBlock-BLOCK_MAX_RANGE)

		for fromBlock < currentBlock {
			if fromBlock > currentBlock {
				break
			}

			toBlock := fromBlock + BLOCK_CHUNK_SIZE

			targetBlock := big.NewInt(int64(toBlock))
			if fromBlock+BLOCK_CHUNK_SIZE > currentBlock {
				targetBlock = nil
			}

			logs, err := c.Client.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   targetBlock,
				Addresses: q.Addresses,
				Topics:    q.Topics,
			})
			if err != nil {
				c.logger.Errorf("Failed to get missed logs: %v", err)
				return err
			} else {
				for _, log := range logs {
					ch2 <- log
				}
			}

			fromBlock += BLOCK_CHUNK_SIZE
		}

		return nil
	}

	resub := func() error {
		c.clientLock.RLock()
		defer c.clientLock.RUnlock()

		err := resubFilterLogs()
		if err != nil {
			return err
		}

		sub, err = c.Client.SubscribeFilterLogs(ctx, q, ch2)
		if err != nil {
			c.logger.Errorf("Failed to resubscribe: %v", err)
			return err
		}

		safeSub.SetUnderlyingSub(sub)

		return nil
	}

	lastBlock = currentBlock

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		ticker := time.NewTicker(RESUB_INTERVAL)
		defer ticker.Stop()

		reinitC := c.WatchReinit()

		for {
			select {
			case <-safeSub.Err():
				return
			case success := <-reinitC:
				reinitC = c.WatchReinit()
				if success {
					err := resub()
					if err != nil {
						c.handleClientError(err)
					}
				}
			case log := <-ch2:
				lastBlock = max(lastBlock, log.BlockNumber)
				ch <- log
			case <-ticker.C:
				err := resub()
				if err != nil {
					c.handleClientError(err)
				}
			case <-sub.Err():
				err := resub()
				if err != nil {
					c.handleClientError(err)
				}
			case <-ctx.Done():
				safeSub.Unsubscribe()
				return
			}
		}
	}()

	return safeSub, nil
}

func (c *SafeEthClient) Close() {
	c.clientLock.Lock()
	defer c.clientLock.Unlock()

	close(c.closeC)
	c.wg.Wait()
}

func (c *SafeEthClient) isConnectionError(err error) bool {
	if err == nil {
		return false
	}

	isConnectionReset := strings.Contains(err.Error(), "connection reset")
	isConnectionRefused := strings.Contains(err.Error(), "connection refused")

	return isConnectionReset || isConnectionRefused
}

func (c *SafeEthClient) handleClientError(err error) {
	if c.isConnectionError(err) {
		c.triggerReinit()
	}
}

func (c *SafeEthClient) triggerReinit() {
	c.reinitC <- struct{}{}
}

func (c *SafeEthClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	sub, err := c.Client.SubscribeNewHead(ctx, ch)
	if err != nil {
		return nil, err
	}

	safeSub := NewSafeSubscription(sub)

	resub := func() error {
		c.clientLock.RLock()
		defer c.clientLock.RUnlock()

		sub, err = c.Client.SubscribeNewHead(ctx, ch)
		if err != nil {
			return err
		}

		safeSub.SetUnderlyingSub(sub)

		return nil
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		reinitC := c.WatchReinit()

		for {
			select {
			case <-safeSub.Err():
				return
			case success := <-reinitC:
				reinitC = c.WatchReinit()
				if success {
					err := resub()
					if err != nil {
						c.handleClientError(err)
					}
				}
			case <-sub.Err():
				err := resub()
				if err != nil {
					c.handleClientError(err)
				}
			case <-ctx.Done():
				safeSub.Unsubscribe()
				return
			}
		}
	}()

	return safeSub, nil
}

func (c *SafeEthClient) GetClientAndVersion() (string, error) {
	client, err := ethclient.Dial(c.rpcUrl)
	if err != nil {
		return "", err
	}

	var clientVersion string
	err = client.Client().Call(&clientVersion, "web3_clientVersion")
	if err != nil {
		return "unavailable", nil
	}
	return clientVersion, nil
}
