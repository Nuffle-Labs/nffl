package safeclient

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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BLOCK_CHUNK_SIZE = 2000
	BLOCK_MAX_RANGE  = 10000
	RESUB_INTERVAL   = 5 * time.Minute
	REINIT_INTERVAL  = 1 * time.Minute
	HEADER_TIMEOUT   = 1 * time.Minute
)

type SafeEthClient struct {
	eth.Client

	clientLock            sync.RWMutex
	reinitSubscribersLock sync.Mutex
	wg                    sync.WaitGroup
	logger                logging.Logger
	rpcUrl                string
	isReinitializing      bool
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
		logger.Error("Failed to create client", "err", err)
		return nil, err
	}
	safeClient.Client = client

	safeClient.wg.Add(1)
	go safeClient.handleReinit()

	logger.Info("Created new SafeEthClient", "rpcUrl", rpcUrl)
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
		c.logger.Debug("Created new eth client without collector")
		return client, nil
	} else {
		client, err := eth.NewInstrumentedClient(c.rpcUrl, c.collector)
		if err != nil {
			return nil, err
		}
		c.logger.Debug("Created new instrumented eth client with collector")
		return client, nil
	}
}

func (c *SafeEthClient) reinit() bool {
	c.clientLock.Lock()
	defer c.clientLock.Unlock()

	client, err := c.createClient()
	if err != nil {
		c.logger.Error("Failed to reinitialize client", "err", err)
		return false
	}

	c.Client = client
	c.logger.Info("Successfully reinitialized client")
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
	c.logger.Debug("Notified subscribers of reinit", "success", success)
}

func (c *SafeEthClient) handleReinitEvent() {
	success := c.reinit()
	c.notifySubscribers(success)
	if success {
		return
	}

	defer c.wg.Done()

	c.isReinitializing = true

	c.wg.Add(1)
	go func() {
		defer func() {
			c.isReinitializing = false
			c.wg.Done()
		}()

		reinitTicker := time.NewTicker(c.reinitInterval)
		defer reinitTicker.Stop()

		for {
			select {
			case <-reinitTicker.C:
				c.logger.Debug("Reinit ticker fired")

				success := c.reinit()
				c.notifySubscribers(success)

				if success {
					return
				}
			case <-c.closeC:
				c.logger.Info("Received close signal, stopping reinit handler")
				return
			}
		}
	}()
}

func (c *SafeEthClient) handleReinit() {
	defer c.wg.Done()

	c.isReinitializing = false

	for {
		select {
		case <-c.closeC:
			c.logger.Info("Received close signal, stopping reinit handler")
			return
		case <-c.reinitC:
			if c.isReinitializing {
				c.logger.Debug("Already reinitializing, ignoring reinit signal")
				continue
			}

			c.logger.Debug("Received reinit signal")
			c.handleReinitEvent()
		}
	}
}

func (c *SafeEthClient) WatchReinit() <-chan bool {
	ch := make(chan bool, 1)
	c.reinitSubscribersLock.Lock()
	c.reinitSubscribers = append(c.reinitSubscribers, ch)
	c.reinitSubscribersLock.Unlock()
	c.logger.Debug("Added reinit watcher")
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
		fromBlock := max(lastBlock, currentBlock-BLOCK_MAX_RANGE) + 1

		for ; fromBlock < currentBlock; fromBlock += (BLOCK_CHUNK_SIZE + 1) {
			toBlock := min(fromBlock+BLOCK_CHUNK_SIZE, currentBlock)

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
		c.clientLock.RLock()
		defer c.clientLock.RUnlock()

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

		ticker := time.NewTicker(RESUB_INTERVAL)
		defer ticker.Stop()

		reinitC := c.WatchReinit()

		for {
			select {
			case <-safeSub.Err():
				c.logger.Debug("Safe subscription ended")
				return
			case success := <-reinitC:
				reinitC = c.WatchReinit()
				if success {
					err := resub()
					c.handleClientError(err)
				}
			case log := <-proxyC:
				// since resub pushes the missed blocks directly to the channel and updates lastBlock, this is ordered
				if lastBlock < log.BlockNumber {
					continue
				}

				lastBlock = log.BlockNumber
				ch <- log
			case <-ticker.C:
				c.logger.Debug("Resub ticker fired")
				err := resub()
				c.handleClientError(err)
			case <-sub.Err():
				c.logger.Info("Underlying subscription ended, resubscribing")
				err := resub()
				c.handleClientError(err)
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
	c.clientLock.Lock()
	defer c.clientLock.Unlock()

	close(c.closeC)
	c.wg.Wait()
	c.logger.Info("SafeEthClient closed")
}

func (c *SafeEthClient) isConnectionError(err error) bool {
	if err == nil {
		return false
	}

	isConnectionReset := strings.Contains(err.Error(), "connection reset")
	isConnectionRefused := strings.Contains(err.Error(), "connection refused")
	isAbnormalClosure := strings.Contains(err.Error(), "abnormal closure")

	return isConnectionReset || isConnectionRefused || isAbnormalClosure
}

func (c *SafeEthClient) handleClientError(err error) {
	if err == nil {
		return
	}

	if c.isConnectionError(err) {
		c.logger.Error("Connection error detected, triggering reinit", "err", err)
		c.triggerReinit()
	} else {
		c.logger.Error("Client error detected", "err", err)
	}
}

func (c *SafeEthClient) triggerReinit() {
	c.reinitC <- struct{}{}
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
		c.clientLock.RLock()
		defer c.clientLock.RUnlock()

		sub, err = c.Client.SubscribeNewHead(ctx, proxyC)
		if err != nil {
			c.logger.Error("Failed to resubscribe to new heads", "err", err)
			return err
		}
		c.logger.Info("Resubscribed to new heads")

		safeSub.SetUnderlyingSub(sub)

		return nil
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		headerTicker := time.NewTicker(HEADER_TIMEOUT)
		defer headerTicker.Stop()

		reinitC := c.WatchReinit()
		receivedBlock := false

		for {
			select {
			case header := <-proxyC:
				receivedBlock = true
				ch <- header
			case <-safeSub.Err():
				c.logger.Info("Safe subscription to new heads ended")
				return
			case success := <-reinitC:
				reinitC = c.WatchReinit()
				if success {
					err := resub()
					c.handleClientError(err)
				}
			case <-sub.Err():
				c.logger.Info("Underlying subscription to new heads ended, resubscribing")
				err := resub()
				c.handleClientError(err)
			case <-headerTicker.C:
				c.logger.Info("Header ticker fired, ending subscription")
				if receivedBlock {
					receivedBlock = false
				} else {
					err := resub()
					c.handleClientError(err)
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
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.ChainID(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.BalanceAt
func (c *SafeEthClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.BalanceAt(ctx, account, blockNumber)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.BlockByHash
func (c *SafeEthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.BlockByHash(ctx, hash)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.BlockByNumber
func (c *SafeEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.BlockByNumber(ctx, number)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.BlockNumber
func (c *SafeEthClient) BlockNumber(ctx context.Context) (uint64, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.BlockNumber(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.CallContract
func (c *SafeEthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.CallContract(ctx, msg, blockNumber)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.CallContractAtHash
func (c *SafeEthClient) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.CallContractAtHash(ctx, msg, blockHash)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.CodeAt
func (c *SafeEthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.CodeAt(ctx, account, blockNumber)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.EstimateGas
func (c *SafeEthClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.EstimateGas(ctx, msg)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.FilterLogs
func (c *SafeEthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.FilterLogs(ctx, q)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.HeaderByHash
func (c *SafeEthClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.HeaderByHash(ctx, hash)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.HeaderByNumber
func (c *SafeEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.HeaderByNumber(ctx, number)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.NetworkID
func (c *SafeEthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.NetworkID(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.NonceAt
func (c *SafeEthClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.NonceAt(ctx, account, blockNumber)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PeerCount
func (c *SafeEthClient) PeerCount(ctx context.Context) (uint64, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PeerCount(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingBalanceAt
func (c *SafeEthClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingBalanceAt(ctx, account)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingCallContract
func (c *SafeEthClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingCallContract(ctx, msg)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingCodeAt
func (c *SafeEthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingCodeAt(ctx, account)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingNonceAt
func (c *SafeEthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingNonceAt(ctx, account)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingStorageAt
func (c *SafeEthClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingStorageAt(ctx, account, key)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.PendingTransactionCount
func (c *SafeEthClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.PendingTransactionCount(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.SendTransaction
func (c *SafeEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	err := c.Client.SendTransaction(ctx, tx)
	c.handleClientError(err)

	return err
}

// Wrapped eth.Client.StorageAt
func (c *SafeEthClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.StorageAt(ctx, account, key, blockNumber)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.SuggestGasPrice
func (c *SafeEthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.SuggestGasPrice(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.SuggestGasTipCap
func (c *SafeEthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.SuggestGasTipCap(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.SyncProgress
func (c *SafeEthClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.SyncProgress(ctx)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.TransactionByHash
func (c *SafeEthClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	tx, isPending, err = c.Client.TransactionByHash(ctx, hash)
	c.handleClientError(err)

	return tx, isPending, err
}

// Wrapped eth.Client.TransactionCount
func (c *SafeEthClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.TransactionCount(ctx, blockHash)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.TransactionInBlock
func (c *SafeEthClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.TransactionInBlock(ctx, blockHash, index)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.TransactionReceipt
func (c *SafeEthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.TransactionReceipt(ctx, txHash)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.TransactionSender
func (c *SafeEthClient) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.TransactionSender(ctx, tx, block, index)
	c.handleClientError(err)

	return result, err
}

// Wrapped eth.Client.FeeHistory
func (c *SafeEthClient) FeeHistory(
	ctx context.Context,
	blockCount uint64,
	lastBlock *big.Int,
	rewardPercentiles []float64,
) (*ethereum.FeeHistory, error) {
	c.clientLock.RLock()
	defer c.clientLock.RUnlock()

	result, err := c.Client.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	c.handleClientError(err)

	return result, err
}
