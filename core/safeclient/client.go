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
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	BLOCK_CHUNK_SIZE   = 2000
	BLOCK_MAX_RANGE    = 10000
	LOG_RESUB_INTERVAL = 5 * time.Minute
	HEADER_TIMEOUT     = 30 * time.Second
)

type SafeClient interface {
	eth.Client

	Close()
}

type SafeEthClient struct {
	eth.Client

	wg               sync.WaitGroup
	logger           logging.Logger
	rpcUrl           string
	closeC           chan struct{}
	closed           bool
	headerTimeout    time.Duration
	logResubInterval time.Duration
	blockChunkSize   uint64
	blockMaxRange    uint64

	createClient func(string, logging.Logger) (eth.Client, error)
}

func NewSafeEthClient(rpcUrl string, logger logging.Logger, opts ...SafeEthClientOption) (*SafeEthClient, error) {
	safeClient := &SafeEthClient{
		logger:           logger,
		rpcUrl:           rpcUrl,
		logResubInterval: LOG_RESUB_INTERVAL,
		headerTimeout:    HEADER_TIMEOUT,
		blockChunkSize:   BLOCK_CHUNK_SIZE,
		blockMaxRange:    BLOCK_MAX_RANGE,
		closeC:           make(chan struct{}),
		createClient:     createDefaultClient,
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

func WithLogResubInterval(interval time.Duration) SafeEthClientOption {
	return func(c *SafeEthClient) {
		c.logResubInterval = interval
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
	underlying   ethereum.Subscription
	lock         sync.Mutex
	errC         chan error
	unsubscribed bool
}

func NewSafeSubscription(sub ethereum.Subscription) *SafeSubscription {
	return &SafeSubscription{
		underlying: sub,
		errC:       make(chan error, 1),
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

	s.underlying.Unsubscribe()
	s.unsubscribed = true
	s.lock.Unlock()

	close(s.errC)
}

func (s *SafeSubscription) SetUnderlyingSub(sub ethereum.Subscription) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.underlying.Unsubscribe()
	s.underlying = sub
}

func (c *SafeEthClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	logCache, err := lru.New[[32]byte, any](100)
	if err != nil {
		c.logger.Error("Failed to create log cache", "err", err)
		return nil, err
	}

	tryCacheLog := func(log *types.Log) bool {
		hash := hashLog(log)
		ok, _ := logCache.ContainsOrAdd(hash, nil)
		return !ok
	}

	currentBlock, err := c.Client.BlockNumber(ctx)
	if err != nil {
		c.logger.Error("Failed to get current block number", "err", err)
		return nil, err
	}
	c.logger.Debug("Got current block number", "block", currentBlock)

	proxyC := make(chan types.Log, 100)

	newSub, err := c.Client.SubscribeFilterLogs(ctx, q, proxyC)
	if err != nil {
		c.logger.Error("Failed to subscribe to logs", "err", err)
		return nil, err
	}
	c.logger.Info("Subscribed to logs")

	safeSub := NewSafeSubscription(newSub)
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

		c.logger.Debug("Comparing last log block with current block", "lastBlock", lastBlock, "currentBlock", currentBlock)

		missedLogs := make([]types.Log, 0)

		rangeStartBlock := currentBlock - c.blockMaxRange
		if c.blockMaxRange > currentBlock {
			rangeStartBlock = 0
		}

		fromBlock := max(lastBlock, rangeStartBlock+1)

		for ; fromBlock < currentBlock; fromBlock += (c.blockChunkSize + 1) {
			toBlock := min(fromBlock+c.blockChunkSize, currentBlock)

			c.logger.Debug("Getting past logs", "fromBlock", fromBlock, "toBlock", toBlock)

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

		safeSub.SetUnderlyingSub(newSub)

		for _, log := range missedLogs {
			if tryCacheLog(&log) {
				lastBlock = max(lastBlock, log.BlockNumber)
				ch <- log
			}
		}

		return nil
	}

	lastBlock = currentBlock

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		ticker := time.NewTicker(c.logResubInterval)
		defer ticker.Stop()

		handleResub := func() {
			err := resub()
			if err != nil {
				c.logger.Error("Failed to resubscribe to logs", "err", err)
				ticker.Reset(c.logResubInterval)
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
				if tryCacheLog(&log) {
					lastBlock = max(lastBlock, log.BlockNumber)
					ch <- log
				}
			case <-ticker.C:
				c.logger.Debug("Resub ticker fired")
				handleResub()
			case <-safeSub.underlying.Err():
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
	proxyC := make(chan *types.Header, 100)

	newSub, err := c.Client.SubscribeNewHead(ctx, proxyC)
	if err != nil {
		c.logger.Error("Failed to subscribe to new heads", "err", err)
		return nil, err
	}
	c.logger.Info("Subscribed to new heads")

	safeSub := NewSafeSubscription(newSub)

	resub := func() error {
		newSub, err := c.Client.SubscribeNewHead(ctx, proxyC)
		if err != nil {
			c.logger.Error("Failed to resubscribe to new heads", "err", err)
			return err
		}
		c.logger.Info("Resubscribed to new heads")

		safeSub.SetUnderlyingSub(newSub)

		return nil
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		headerTicker := time.NewTicker(c.headerTimeout)
		defer headerTicker.Stop()

		handleResub := func() {
			err := resub()
			if err != nil {
				c.logger.Error("Failed to resubscribe to heads", "err", err)
				headerTicker.Reset(c.headerTimeout)
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
			case <-safeSub.underlying.Err():
				c.logger.Info("Underlying subscription to new heads ended, resubscribing")
				handleResub()
			case <-headerTicker.C:
				c.logger.Debug("Header ticker fired")
				if receivedBlock {
					receivedBlock = false
				} else {
					c.logger.Info("No block received, resubscribing")
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
