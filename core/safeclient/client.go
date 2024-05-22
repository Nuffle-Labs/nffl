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
	"github.com/ethereum/go-ethereum/common"
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
				err := resub()
				if err != nil {
					c.logger.Error("Failed to resubscribe to logs", "err", err)
					ticker.Reset(c.resubInterval)
				} else {
					ticker.Stop()
				}
			case <-sub.Err():
				c.logger.Info("Underlying subscription ended, resubscribing")
				err := resub()
				if err != nil {
					c.logger.Error("Failed to resubscribe to logs", "err", err)
					ticker.Reset(c.resubInterval)
				} else {
					ticker.Stop()
				}
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
				err := resub()
				if err != nil {
					c.logger.Error("Failed to resubscribe to heads", "err", err)
					headerTicker.Reset(c.resubInterval)
				} else {
					headerTicker.Stop()
				}
			case <-headerTicker.C:
				c.logger.Info("Header ticker fired, ending subscription")
				if receivedBlock {
					receivedBlock = false
				} else {
					err := resub()
					if err != nil {
						c.logger.Error("Failed to resubscribe to heads", "err", err)
						resubTicker.Reset(c.resubInterval)
					} else {
						resubTicker.Stop()
					}
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

// Wrapped eth.Client.ChainID
func (c *SafeEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	result, err := c.Client.ChainID(ctx)

	return result, err
}

// Wrapped eth.Client.BalanceAt
func (c *SafeEthClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	result, err := c.Client.BalanceAt(ctx, account, blockNumber)

	return result, err
}

// Wrapped eth.Client.BlockByHash
func (c *SafeEthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	result, err := c.Client.BlockByHash(ctx, hash)

	return result, err
}

// Wrapped eth.Client.BlockByNumber
func (c *SafeEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	result, err := c.Client.BlockByNumber(ctx, number)

	return result, err
}

// Wrapped eth.Client.BlockNumber
func (c *SafeEthClient) BlockNumber(ctx context.Context) (uint64, error) {
	result, err := c.Client.BlockNumber(ctx)

	return result, err
}

// Wrapped eth.Client.CallContract
func (c *SafeEthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	result, err := c.Client.CallContract(ctx, msg, blockNumber)

	return result, err
}

// Wrapped eth.Client.CallContractAtHash
func (c *SafeEthClient) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error) {
	result, err := c.Client.CallContractAtHash(ctx, msg, blockHash)

	return result, err
}

// Wrapped eth.Client.CodeAt
func (c *SafeEthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	result, err := c.Client.CodeAt(ctx, account, blockNumber)

	return result, err
}

// Wrapped eth.Client.EstimateGas
func (c *SafeEthClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	result, err := c.Client.EstimateGas(ctx, msg)

	return result, err
}

// Wrapped eth.Client.FilterLogs
func (c *SafeEthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	result, err := c.Client.FilterLogs(ctx, q)

	return result, err
}

// Wrapped eth.Client.HeaderByHash
func (c *SafeEthClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	result, err := c.Client.HeaderByHash(ctx, hash)

	return result, err
}

// Wrapped eth.Client.HeaderByNumber
func (c *SafeEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	result, err := c.Client.HeaderByNumber(ctx, number)

	return result, err
}

// Wrapped eth.Client.NetworkID
func (c *SafeEthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	result, err := c.Client.NetworkID(ctx)

	return result, err
}

// Wrapped eth.Client.NonceAt
func (c *SafeEthClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	result, err := c.Client.NonceAt(ctx, account, blockNumber)

	return result, err
}

// Wrapped eth.Client.PeerCount
func (c *SafeEthClient) PeerCount(ctx context.Context) (uint64, error) {
	result, err := c.Client.PeerCount(ctx)

	return result, err
}

// Wrapped eth.Client.PendingBalanceAt
func (c *SafeEthClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	result, err := c.Client.PendingBalanceAt(ctx, account)

	return result, err
}

// Wrapped eth.Client.PendingCallContract
func (c *SafeEthClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	result, err := c.Client.PendingCallContract(ctx, msg)

	return result, err
}

// Wrapped eth.Client.PendingCodeAt
func (c *SafeEthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	result, err := c.Client.PendingCodeAt(ctx, account)

	return result, err
}

// Wrapped eth.Client.PendingNonceAt
func (c *SafeEthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	result, err := c.Client.PendingNonceAt(ctx, account)

	return result, err
}

// Wrapped eth.Client.PendingStorageAt
func (c *SafeEthClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	result, err := c.Client.PendingStorageAt(ctx, account, key)

	return result, err
}

// Wrapped eth.Client.PendingTransactionCount
func (c *SafeEthClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	result, err := c.Client.PendingTransactionCount(ctx)

	return result, err
}

// Wrapped eth.Client.SendTransaction
func (c *SafeEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := c.Client.SendTransaction(ctx, tx)

	return err
}

// Wrapped eth.Client.StorageAt
func (c *SafeEthClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	result, err := c.Client.StorageAt(ctx, account, key, blockNumber)

	return result, err
}

// Wrapped eth.Client.SuggestGasPrice
func (c *SafeEthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	result, err := c.Client.SuggestGasPrice(ctx)

	return result, err
}

// Wrapped eth.Client.SuggestGasTipCap
func (c *SafeEthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	result, err := c.Client.SuggestGasTipCap(ctx)

	return result, err
}

// Wrapped eth.Client.SyncProgress
func (c *SafeEthClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	result, err := c.Client.SyncProgress(ctx)

	return result, err
}

// Wrapped eth.Client.TransactionByHash
func (c *SafeEthClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	tx, isPending, err = c.Client.TransactionByHash(ctx, hash)

	return tx, isPending, err
}

// Wrapped eth.Client.TransactionCount
func (c *SafeEthClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	result, err := c.Client.TransactionCount(ctx, blockHash)

	return result, err
}

// Wrapped eth.Client.TransactionInBlock
func (c *SafeEthClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	result, err := c.Client.TransactionInBlock(ctx, blockHash, index)

	return result, err
}

// Wrapped eth.Client.TransactionReceipt
func (c *SafeEthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	result, err := c.Client.TransactionReceipt(ctx, txHash)

	return result, err
}

// Wrapped eth.Client.TransactionSender
func (c *SafeEthClient) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	result, err := c.Client.TransactionSender(ctx, tx, block, index)

	return result, err
}

// Wrapped eth.Client.FeeHistory
func (c *SafeEthClient) FeeHistory(
	ctx context.Context,
	blockCount uint64,
	lastBlock *big.Int,
	rewardPercentiles []float64,
) (*ethereum.FeeHistory, error) {
	result, err := c.Client.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)

	return result, err
}
