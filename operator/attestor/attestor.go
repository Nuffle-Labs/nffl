package attestor

import (
	"context"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/safeclient"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/operator/consumer"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
)

const (
	MQ_WAIT_TIMEOUT        = 30 * time.Second
	MQ_REBROADCAST_TIMEOUT = 15 * time.Second
)

var (
	unknownRollupIdError = errors.New("notify: rollupId unknown")
)

type Attestorer interface {
	core.Metricable

	Start(ctx context.Context) error
	Close() error
	GetSignedRootC() <-chan messages.SignedStateRootUpdateMessage
}

// Attestor subscribes for RPCs block updates
// Also subscribes for MQ blocks from Consumer
// Each block from RPC waits for MQ_WAIT_TIMEOUT for MQ block
// In case same block doesn't arrive from MQ block is signed and sent
// If it arrives it is compared and then sent to Aggregator
type Attestor struct {
	signedRootC        chan messages.SignedStateRootUpdateMessage
	rollupIdsToUrls    map[uint32]string
	clients            map[uint32]eth.Client
	clientsLock        sync.Mutex
	rpcCallsCollectors map[uint32]*rpccalls.Collector
	notifier           Notifier
	consumer           *consumer.Consumer

	config     *optypes.NodeConfig
	blsKeypair *bls.KeyPair
	operatorId eigentypes.OperatorId

	logger   sdklogging.Logger
	listener EventListener
	// TODO(edwin): remove after https://github.com/Layr-Labs/eigensdk-go/pull/117 merged
	registry *prometheus.Registry
}

var _ core.Metricable = (*Attestor)(nil)

func NewAttestor(config *optypes.NodeConfig, blsKeypair *bls.KeyPair, operatorId eigentypes.OperatorId, registry *prometheus.Registry, logger sdklogging.Logger) (*Attestor, error) {
	consumer := consumer.NewConsumer(consumer.ConsumerConfig{
		RollupIds: config.NearDaIndexerRollupIds,
		Id:        hex.EncodeToString(operatorId[:]),
	}, logger)

	attestor := Attestor{
		signedRootC:        make(chan messages.SignedStateRootUpdateMessage),
		rollupIdsToUrls:    make(map[uint32]string),
		clients:            make(map[uint32]eth.Client),
		rpcCallsCollectors: make(map[uint32]*rpccalls.Collector),
		logger:             logger,
		notifier:           NewNotifier(),
		consumer:           consumer,
		blsKeypair:         blsKeypair,
		operatorId:         operatorId,
		registry:           registry,
		listener:           &SelectiveEventListener{},
		config:             config,
	}

	for rollupId, url := range config.RollupIdsToRpcUrls {
		var rpcCallsCollector *rpccalls.Collector
		if config.EnableMetrics {
			id := config.OperatorAddress + AttestorSubsystem
			rpcCallsCollector = rpccalls.NewCollector(id+url, registry)
		}

		clientOpts := make([]safeclient.SafeEthClientOption, 0)
		if rpcCallsCollector != nil {
			clientOpts = append(clientOpts, safeclient.WithInstrumentedCreateClient(rpcCallsCollector))
		}

		client, err := safeclient.NewSafeEthClient(url, logger, clientOpts...)
		if err != nil {
			return nil, err
		}

		attestor.rollupIdsToUrls[rollupId] = url

		attestor.clients[rollupId] = client
		attestor.rpcCallsCollectors[rollupId] = rpcCallsCollector
	}

	return &attestor, nil
}

func (attestor *Attestor) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeAttestorMetrics(registry)
	if err != nil {
		return err
	}
	attestor.listener = listener

	if err = attestor.consumer.EnableMetrics(registry); err != nil {
		return err
	}

	return nil
}

func (attestor *Attestor) Start(ctx context.Context) error {
	go attestor.consumer.Start(ctx, attestor.config.NearDaIndexerRmqIpPortAddress)

	subscriptions := make(map[uint32]ethereum.Subscription)
	headersCs := make(map[uint32]chan *ethtypes.Header)

	for rollupId, client := range attestor.clients {
		headersC := make(chan *ethtypes.Header, 100)
		subscription, err := client.SubscribeNewHead(ctx, headersC)
		if err != nil {
			attestor.logger.Fatalf("Failed to subscribe to new header: %v, for rollupId: %v", err, rollupId)
			return err
		}

		blockNumber, err := client.BlockNumber(ctx)
		if err != nil {
			attestor.logger.Fatalf("Failed to get block number: %v, for rollupId: %v", err, rollupId)
			return err
		}

		attestor.listener.ObserveInitializationInitialBlockNumber(rollupId, blockNumber)

		subscriptions[rollupId] = subscription
		headersCs[rollupId] = headersC
	}

	go attestor.processMQBlocks(ctx)

	for rollupId := range attestor.clients {
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
			attestor.logger.Info("Notifying", "rollupId", mqBlock.RollupId, "height", mqBlock.Block.Header().Number.Uint64())
			err := attestor.notifier.Notify(mqBlock.RollupId, mqBlock)
			if err != nil {
				attestor.logger.Errorf("Notifier: %v", err)
			}

			// Rebroadcast in case mq block arrives first
			go func(mqBlock consumer.BlockData) {
				select {
				case <-time.After(MQ_REBROADCAST_TIMEOUT):
					attestor.logger.Info("Renotifying", "rollupId", mqBlock.RollupId, "height", mqBlock.Block.Header().Number.Uint64())

					err := attestor.notifier.Notify(mqBlock.RollupId, mqBlock)
					if err != nil {
						attestor.logger.Error("Error while renotifying", "err", err)
					}

					return
				case <-ctx.Done():
					return
				}
			}(mqBlock)
		}
	}
}

// Spawns routines for new headers that die in one minute
func (attestor *Attestor) processRollupHeaders(rollupId uint32, headersC chan *ethtypes.Header, subscription ethereum.Subscription, ctx context.Context) {
	for {
		select {
		case <-subscription.Err():
			attestor.logger.Error("Header subscription error", "rollupId", rollupId)
			return
		case header, ok := <-headersC:
			if !ok {
				return
			}

			go attestor.processHeader(rollupId, header, ctx)
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
	attestor.logger.Info("Processing header", "rollupId", rollupId, "height", rollupHeader.Number.Uint64())

	attestor.listener.ObserveLastBlockReceived(rollupId, rollupHeader.Number.Uint64())
	attestor.listener.ObserveLastBlockReceivedTimestamp(rollupId, uint64(rollupHeader.Time))
	attestor.listener.OnBlockReceived(rollupId)

	predicate := func(mqBlock consumer.BlockData) bool {
		if mqBlock.RollupId != rollupId {
			attestor.logger.Warnf("Subscriber expected rollupId: %v, but got %v", rollupId, mqBlock.RollupId)
			return false
		}

		if rollupHeader.Number.Cmp(mqBlock.Block.Header().Number) != 0 {
			return false
		}

		if mqBlock.Block.Header().Root != rollupHeader.Root {
			attestor.logger.Warnf("StateRoot from MQ doesn't match one from Node")
			attestor.listener.OnBlockMismatch(rollupId)

			return false
		}

		return true
	}

	mqBlocksC, id := attestor.notifier.Subscribe(rollupId, predicate)
	defer attestor.notifier.Unsubscribe(rollupId, id)

	transactionId := [32]byte{0}
	daCommitment := [32]byte{0}

	timer := time.After(MQ_WAIT_TIMEOUT)
loop:
	for {
		select {
		case <-timer:
			attestor.logger.Info("MQ timeout", "rollupId", rollupId, "height", rollupHeader.Number.Uint64())
			attestor.listener.OnMissedMQBlock(rollupId)

			break loop

		case mqBlock := <-mqBlocksC:
			attestor.logger.Info("MQ block found", "height", mqBlock.Block.Header().Number.Uint64(), "rollupId", mqBlock.RollupId)

			daCommitment = mqBlock.Commitment
			transactionId = mqBlock.TransactionId

			break loop

		case <-ctx.Done():
			return
		}
	}

	message := messages.StateRootUpdateMessage{
		RollupId:            rollupId,
		BlockHeight:         rollupHeader.Number.Uint64(),
		Timestamp:           rollupHeader.Time,
		StateRoot:           rollupHeader.Root,
		NearDaTransactionId: transactionId,
		NearDaCommitment:    daCommitment,
	}
	signature, err := SignStateRootUpdateMessage(attestor.blsKeypair, &message)
	if err != nil {
		attestor.logger.Warn("StateRoot sign failed", "err", err)
		return
	}

	signedStateRootUpdateMessage := messages.SignedStateRootUpdateMessage{
		Message:      message,
		BlsSignature: *signature,
		OperatorId:   attestor.operatorId,
	}

	attestor.signedRootC <- signedStateRootUpdateMessage
}

func SignStateRootUpdateMessage(blsKeypair *bls.KeyPair, stateRootUpdateMessage *messages.StateRootUpdateMessage) (*bls.Signature, error) {
	messageDigest, err := stateRootUpdateMessage.Digest()
	if err != nil {
		return nil, err
	}

	blsSignature := blsKeypair.SignMessage(messageDigest)
	return blsSignature, nil
}

func (attestor *Attestor) GetSignedRootC() <-chan messages.SignedStateRootUpdateMessage {
	return attestor.signedRootC
}

func (attestor *Attestor) Close() error {
	if err := attestor.consumer.Close(); err != nil {
		return err
	}

	return nil
}
