package attestor

import (
	"context"
	"errors"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"

	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	"github.com/NethermindEth/near-sffl/types"
)

const (
	MQ_WAIT_TIMEOUT       = time.Second
	MQ_REBROADCAST_DELAY  = 10 * time.Second
	RECONNECTION_ATTEMPTS = 5
	RECONNECTION_DELAY    = time.Second
)

var (
	unknownRollupIdError = errors.New("notify: rollupId unknown")
)

func createEthClient(rpcUrl string, enableMetrics bool, registry *prometheus.Registry, logger sdklogging.Logger) (eth.EthClient, error) {
	if enableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(rpcUrl, registry)
		ethClient, err := eth.NewInstrumentedClient(rpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Error("Cannot create http ethclient", "err", err)
			return nil, err
		}

		return ethClient, nil
	}

	ethClient, err := eth.NewClient(rpcUrl)
	if err != nil {
		logger.Error("Cannot create http ethclient", "err", err)
		return nil, err
	}

	return ethClient, nil
}

type Attestorer interface {
	Start(ctx context.Context) error
	Close() error
	GetSignedRootC() <-chan coretypes.SignedStateRootUpdateMessage
}

// Attestor subscribes for RPCs block updates
// Also subscribes for MQ blocks from Consumer
// Each block from RPC waits for MQ_WAIT_TIMEOUT for MQ block
// In case same block doesn't arrive from MQ block is signed and sent
// If it arrives it is compared and then sent to Aggregator
type Attestor struct {
	signedRootC     chan coretypes.SignedStateRootUpdateMessage
	rollupIdsToUrls map[uint32]string
	clients         map[uint32]eth.EthClient
	notifier        Notifier
	consumer        *consumer.Consumer

	registry   *prometheus.Registry
	config     *types.NodeConfig
	blsKeypair *bls.KeyPair
	operatorId bls.OperatorId

	logger sdklogging.Logger
}

func NewAttestor(config *types.NodeConfig, blsKeypair *bls.KeyPair, operatorId bls.OperatorId, logger sdklogging.Logger) (*Attestor, error) {
	registry := prometheus.NewRegistry()

	consumer := consumer.NewConsumer(consumer.ConsumerConfig{
		Addr:      config.NearDaIndexerRmqIpPortAddress,
		RollupIds: config.NearDaIndexerRollupIds,
	}, logger)

	attestor := Attestor{
		signedRootC: make(chan coretypes.SignedStateRootUpdateMessage),
		clients:     make(map[uint32]eth.EthClient),
		logger:      logger,
		notifier:    NewNotifier(),
		consumer:    consumer,
		blsKeypair:  blsKeypair,
		operatorId:  operatorId,
		registry:    registry,
		config:      config,
	}

	for rollupId, url := range config.RollupIdsToRpcUrls {
		client, err := createEthClient(url, config.EnableMetrics, registry, logger)
		if err != nil {
			return nil, err
		}

		attestor.clients[rollupId] = client
	}

	return &attestor, nil
}

func (attestor *Attestor) Start(ctx context.Context) error {
	subscriptions := make(map[uint32]ethereum.Subscription)
	headersCs := make(map[uint32]chan *ethtypes.Header)

	for rollupId, client := range attestor.clients {
		headersC := make(chan *ethtypes.Header)
		subscription, err := client.SubscribeNewHead(ctx, headersC)
		if err != nil {
			attestor.logger.Fatalf("Failed to subscribe to new header: %v, for rollupId: %v", err, rollupId)
			return err
		}

		subscriptions[rollupId] = subscription
		headersCs[rollupId] = headersC
	}

	go attestor.processMQBlocks(ctx)

	for rollupId, _ := range attestor.clients {
		go attestor.processRollupHeaders(rollupId, headersCs[rollupId], subscriptions[rollupId], ctx)
	}

	return nil
}

// Receives MQ blocks and broadcasts them for a particular rollup
func (attestor *Attestor) processMQBlocks(ctx context.Context) {
	mqBlockC := attestor.consumer.GetBlockStream()

	for {
		select {
		case <-ctx.Done():
			return
		case mqBlock := <-mqBlockC:
			err := attestor.notifier.Notify(mqBlock.RollupId, mqBlock)
			if err != nil {
				attestor.logger.Errorf("Notifier: %v", err)
			}

			// Rebroadcast in case mq block arrives first
			go func(mqBlock consumer.BlockData) {
				select {
				case <-time.After(MQ_REBROADCAST_DELAY):
					attestor.notifier.Notify(mqBlock.RollupId, mqBlock)
					return

				case <-ctx.Done():
					return
				}
			}(mqBlock)
		}
	}
}

func (attestor *Attestor) reconnectClient(rollupId uint32) (eth.EthClient, error) {
	var err error
	var client eth.EthClient
	for i := 0; i < RECONNECTION_ATTEMPTS; i++ {
		<-time.After(RECONNECTION_DELAY)

		client, err = createEthClient(attestor.rollupIdsToUrls[rollupId], attestor.config.EnableMetrics, attestor.registry, attestor.logger)
		if err == nil {
			return client, nil
		}
	}

	return nil, err
}

// Spawns routines for new headers that die in one minute
func (attestor *Attestor) processRollupHeaders(rollupId uint32, headersC chan *ethtypes.Header, subscription ethereum.Subscription, ctx context.Context) {
	for {
		select {
		case <-subscription.Err():
			subscription.Unsubscribe()

			client, err := attestor.reconnectClient(rollupId)
			if err != nil {
				return
			}
			attestor.clients[rollupId] = client

			subscription, err = client.SubscribeNewHead(ctx, headersC)
			if err != nil {
				return
			}

			continue

		case header, ok := <-headersC:
			if !ok {
				continue
			}

			go attestor.processHeader(rollupId, header, ctx)
			continue

		case <-ctx.Done():
			subscription.Unsubscribe()
			close(headersC)

			return
		}
	}
}

// Waits for MQ block for 1 minute. Then signs off and sends
// Filters until receives one having same height
func (attestor *Attestor) processHeader(rollupId uint32, rollupHeader *ethtypes.Header, ctx context.Context) {
	mqBlocksC, id := attestor.notifier.Subscribe(rollupId)
	defer attestor.notifier.Unsubscribe(rollupId, id)

loop:
	for {
		select {
		case <-time.After(MQ_WAIT_TIMEOUT):
			break loop

		case mqBlock := <-mqBlocksC:
			if mqBlock.RollupId != rollupId {
				attestor.logger.Warnf("Subsriber expected rollupId: %v, but got %v", rollupId, mqBlock.RollupId)
				continue loop
			}

			// Filter notifications
			if rollupHeader.Number != mqBlock.Block.Header().Number {
				continue loop
			}

			if mqBlock.Block.Header().Root != rollupHeader.Root {
				// TODO: Do smth here
				attestor.logger.Warnf("StateRoot from MQ doesn't match one from Node")
			}

			break loop

		case <-ctx.Done():
			return
		}
	}

	message := servicemanager.StateRootUpdateMessage{
		RollupId:    rollupId,
		BlockHeight: rollupHeader.Number.Uint64(),
		Timestamp:   rollupHeader.Time,
		StateRoot:   rollupHeader.Root,
	}
	signature, err := SignStateRootUpdateMessage(attestor.blsKeypair, &message)
	if err != nil {
		attestor.logger.Warn("StateRoot sign failed", "err", err)
		return
	}

	signedStateRootUpdateMessage := coretypes.SignedStateRootUpdateMessage{
		Message:      message,
		BlsSignature: *signature,
		OperatorId:   attestor.operatorId,
	}

	attestor.signedRootC <- signedStateRootUpdateMessage
}

func SignStateRootUpdateMessage(blsKeypair *bls.KeyPair, stateRootUpdateMessage *servicemanager.StateRootUpdateMessage) (*bls.Signature, error) {
	messageDigest, err := core.GetStateRootUpdateMessageDigest(stateRootUpdateMessage)
	if err != nil {
		return nil, err
	}

	blsSignature := blsKeypair.SignMessage(messageDigest)
	return blsSignature, nil
}

func (attestor *Attestor) GetSignedRootC() <-chan coretypes.SignedStateRootUpdateMessage {
	return attestor.signedRootC
}

func (attestor *Attestor) Close() error {
	if err := attestor.consumer.Close(); err != nil {
		return err
	}

	return nil
}
