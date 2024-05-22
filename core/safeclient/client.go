package safeclient

import (
	"context"
	"math/big"
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
	BLOCK_CHUNK_SIZE = 2000
	BLOCK_MAX_RANGE  = 10000
	RESUB_INTERVAL   = 5 * time.Minute
	HEADER_TIMEOUT   = 1 * time.Minute
)

type SafeClient interface {
	eth.Client

	Close()
}

type SafeEthClient struct {
	eth.Client

	wg             sync.WaitGroup
	logger         logging.Logger
	rpcUrl         string
	closeC         chan struct{}
	closed         bool
	headerTimeout  time.Duration
	resubInterval  time.Duration
	blockChunkSize uint64
	blockMaxRange  uint64

	createClient func(string, logging.Logger) (eth.Client, error)
}

func NewSafeEthClient(rpcUrl string, logger logging.Logger, opts ...SafeEthClientOption) (*SafeEthClient, error) {
	safeClient := &SafeEthClient{
		logger:         logger,
		rpcUrl:         rpcUrl,
		resubInterval:  RESUB_INTERVAL,
		headerTimeout:  HEADER_TIMEOUT,
		blockChunkSize: BLOCK_CHUNK_SIZE,
		blockMaxRange:  BLOCK_MAX_RANGE,
		closeC:         make(chan struct{}),
		createClient:   createDefaultClient,
	}

	for _, opt := range opts {
		opt(safeClient)
	}

	client, err := safeClient.createClient(rpcUrl, logger)
	if err != nil {
		logger.Error("Failed to create client", "err", err)
		return nil, err
	}
	safeClient.Client = client

	logger.Info("Created new SafeEthClient", "rpcUrl", rpcUrl)
	return safeClient, nil
}

type SafeEthClientOption func(*SafeEthClient)

func WithResubInterval(interval time.Duration) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.resubInterval = interval
	}
}

func WithHeaderTimeout(timeout time.Duration) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.headerTimeout = timeout
	}
}

func WithLogFilteringParams(chunkSize, maxRange uint64) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.blockChunkSize = chunkSize
		c.blockMaxRange = maxRange
	}
}

func WithInstrumentedCreateClient(collector *rpccalls.Collector) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.createClient = func(rpcUrl string, logger logging.Logger) (eth.Client, error) {
			return createInstrumentedClient(rpcUrl, collector, logger)
		}
	}
}

func WithCustomCreateClient(createClient func(string, logging.Logger) (eth.Client, error)) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.createClient = createClient
	}
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
	currentBlock, err := c.Client.BlockNumber(ctx)
	if err != nil {
		c.logger.Error("Failed to get current block number", "err", err)
		return nil, err
	}
	c.logger.Debug("Got current block number", "block", currentBlock)

	proxyC := make(chan types.Log, 100)

	sub, err := c.Client.SubscribeFilterLogs(ctx, q, proxyC)
	if err != nil {
		c.logger.Error("Failed to subscribe to logs", "err", err)
		return nil, err
	}
	c.logger.Info("Subscribed to logs")

	safeSub := NewSafeSubscription(sub)
	lastBlock := currentBlock

	resubFilterLogs := func() ([]types.Log, error) {
		currentBlock, err := c.Client.BlockNumber(ctx)
		if err != nil {
			c.logger.Error("Failed to get current block number", "err", err)
			return nil, err
		}
		c.logger.Debug("Got current block number for resub", "block", currentBlock)

		if lastBlock >= currentBlock {
			return nil, nil
		}

		missedLogs := make([]types.Log, 0)
		fromBlock := max(lastBlock, currentBlock-c.blockMaxRange) + 1

		for ; fromBlock < currentBlock; fromBlock += (c.blockChunkSize + 1) {
			toBlock := min(fromBlock+c.blockChunkSize, currentBlock)

			logs, err := c.Client.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(toBlock)),
				Addresses: q.Addresses,
				Topics:    q.Topics,
			})
			if err != nil {
				c.logger.Error("Failed to get missed logs", "err", err)
				return nil, err
			} else {
				c.logger.Info("Got missed logs on resubscribe", "count", len(logs))
				missedLogs = append(missedLogs, logs...)
			}
		}

		return missedLogs, nil
	}

	resub := func() error {
		newSub, err := c.Client.SubscribeFilterLogs(ctx, q, proxyC)
		if err != nil {
			c.logger.Error("Failed to resubscribe to logs", "err", err)
			return err
		}
		c.logger.Info("Resubscribed to logs")

		missedLogs, err := resubFilterLogs()
		if err != nil {
			c.logger.Error("Failed to get missed logs", "err", err)
			newSub.Unsubscribe()
			return err
		}

		sub = newSub
		safeSub.SetUnderlyingSub(sub)

		for _, log := range missedLogs {
			lastBlock = max(lastBlock, log.BlockNumber)
			ch <- log
		}

		return nil
	}

	lastBlock = currentBlock

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		ticker := time.NewTicker(c.resubInterval)
		defer ticker.Stop()

		handleResub := func() {
			err := resub()
			if err != nil {
				c.logger.Error("Failed to resubscribe to logs", "err", err)
				ticker.Reset(c.resubInterval)
			} else {
				ticker.Stop()
			}
		}

		for {
			select {
			case <-safeSub.Err():
				c.logger.Debug("Safe subscription ended")
				return
			case log := <-proxyC:
				// if that's the case, then most likely we got an event on filterLog and are getting the same one in the sub
				if lastBlock > log.BlockNumber {
					continue
				}

				// since resub pushes the missed blocks directly to the channel and updates lastBlock, this is ordered
				lastBlock = log.BlockNumber
				ch <- log
			case <-ticker.C:
				c.logger.Debug("Resub ticker fired")
				handleResub()
			case <-sub.Err():
				c.logger.Info("Underlying subscription ended, resubscribing")
				handleResub()
			case <-c.closeC:
				c.logger.Info("Received close signal, ending subscription")
				safeSub.Unsubscribe()
				return
			case <-ctx.Done():
				c.logger.Info("Context done, ending subscription")
				safeSub.Unsubscribe()
				return
			}
		}
	}()

	return safeSub, nil
}

func (c *SafeEthClient) Close() {
	if c.closed {
		return
	}

	close(c.closeC)
	c.wg.Wait()
	c.logger.Info("SafeEthClient closed")

	c.closed = true
}

func (c *SafeEthClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	sub, err := c.Client.SubscribeNewHead(ctx, ch)
	if err != nil {
		c.logger.Error("Failed to subscribe to new heads", "err", err)
		return nil, err
	}
	c.logger.Info("Subscribed to new heads")

	safeSub := NewSafeSubscription(sub)
	proxyC := make(chan *types.Header, 100)

	resub := func() error {
		newSub, err := c.Client.SubscribeNewHead(ctx, proxyC)
		if err != nil {
			c.logger.Error("Failed to resubscribe to new heads", "err", err)
			return err
		}
		c.logger.Info("Resubscribed to new heads")

		sub = newSub
		safeSub.SetUnderlyingSub(sub)

		return nil
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		headerTicker := time.NewTicker(c.headerTimeout)
		defer headerTicker.Stop()

		resubTicker := time.NewTicker(c.resubInterval)
		defer resubTicker.Stop()

		handleResub := func() {
			err := resub()
			if err != nil {
				c.logger.Error("Failed to resubscribe to heads", "err", err)
				resubTicker.Reset(c.resubInterval)
			} else {
				resubTicker.Stop()
			}
		}

		receivedBlock := false

		for {
			select {
			case header := <-proxyC:
				receivedBlock = true
				ch <- header
			case <-safeSub.Err():
				c.logger.Info("Safe subscription to new heads ended")
				return
			case <-sub.Err():
				c.logger.Info("Underlying subscription to new heads ended, resubscribing")
				handleResub()
			case <-headerTicker.C:
				c.logger.Info("Header ticker fired, ending subscription")
				if receivedBlock {
					receivedBlock = false
				} else {
					handleResub()
				}
			case <-c.closeC:
				c.logger.Info("Received close signal, ending new heads subscription")
				safeSub.Unsubscribe()
				return
			case <-ctx.Done():
				c.logger.Info("Context done, ending new heads subscription")
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
		c.logger.Error("Failed to dial client for version", "err", err)
		return "", err
	}

	var clientVersion string
	err = client.Client().Call(&clientVersion, "web3_clientVersion")
	if err != nil {
		c.logger.Error("Failed to get client version", "err", err)
		return "unavailable", nil
	}
	c.logger.Info("Got client version", "version", clientVersion)
	return clientVersion, nil
}
