package relayer

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	near "github.com/near/rollup-data-availability/gopkg/da-rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/relayer/config"
)

const (
	NAMESPACE_ID               = 1
	SUBMIT_BLOCK_INTERVAL      = 2500 * time.Millisecond
	SUBMIT_BLOCK_RETRY_TIMEOUT = 2 * time.Second
	SUBMIT_BLOCK_RETRIES       = 3
)

type Relayer struct {
	rpcClient   eth.Client
	rpcUrl      string
	daAccountId string
	logger      sdklogging.Logger
	listener    EventListener

	nearClient *near.Config
}

var _ core.Metricable = (*Relayer)(nil)

func NewRelayerFromConfig(config *config.RelayerConfig, logger sdklogging.Logger) (*Relayer, error) {
	rpcClient, err := core.NewSafeEthClient(config.RpcUrl, logger)
	if err != nil {
		return nil, err
	}

	nearClient, err := near.NewConfigFile(config.KeyPath, config.DaAccountId, config.Network, NAMESPACE_ID)
	if err != nil {
		return nil, err
	}

	return &Relayer{
		rpcClient:   rpcClient,
		rpcUrl:      config.RpcUrl,
		daAccountId: config.DaAccountId,
		nearClient:  nearClient,
		logger:      logger,
		listener:    &SelectiveListener{},
	}, nil
}

func (r *Relayer) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeRelayerMetrics(registry)
	if err != nil {
		return err
	}

	r.listener = listener
	return nil
}

func (r *Relayer) Start(ctx context.Context) error {
	blocksToSubmit := make(chan []*ethtypes.Block)

	ticker := time.NewTicker(SUBMIT_BLOCK_INTERVAL)
	defer ticker.Stop()

	go r.listenToBlocks(ctx, blocksToSubmit, ticker)

	for {
		select {
		case blocks := <-blocksToSubmit:
			err := r.handleBlocks(blocks, ticker)
			if err != nil {
				r.logger.Error("Error handling blocks", "err", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (r *Relayer) handleBlocks(blocks []*ethtypes.Block, ticker *time.Ticker) error {
	defer ticker.Reset(SUBMIT_BLOCK_INTERVAL)

	blockNumbers := make([]uint64, len(blocks))
	for i, block := range blocks {
		blockNumbers[i] = block.Number().Uint64()
	}
	r.logger.Info("Submitting blocks to NEAR", "numbers", blockNumbers)

	encodedBlocks, err := rlp.EncodeToBytes(blocks)
	if err != nil {
		r.logger.Error("Error RLP encoding block", "err", err.Error())
		return err
	}

	out, err := r.submitEncodedBlocks(encodedBlocks)
	if err != nil {
		r.logger.Error("Error submitting encoded blocks", "err", err)
		r.listener.OnDaSubmissionFailed()

		return err
	}

	r.logger.Info(string(out))

	return nil
}

func (r *Relayer) submitEncodedBlocks(encodedBlocks []byte) ([]byte, error) {
	startTime := time.Now()
	for i := 0; i < SUBMIT_BLOCK_RETRIES; i++ {
		out, err := r.nearClient.ForceSubmit(encodedBlocks)
		if err == nil {
			r.listener.OnDaSubmitted(time.Since(startTime))
			r.listener.OnRetriesRequired(i)

			return out, nil
		}

		r.logger.Error("Error submitting blocks to NEAR, resubmitting", "err", err)

		if strings.Contains(err.Error(), "InvalidNonce") {
			r.logger.Info("Invalid nonce, resubmitting", "err", err)
			r.listener.OnInvalidNonce()
			time.Sleep(SUBMIT_BLOCK_RETRY_TIMEOUT)
		} else if strings.Contains(err.Error(), "Expired") {
			r.logger.Info("Expired, resubmitting", "err", err)
			r.listener.OnExpiredTx()
			time.Sleep(SUBMIT_BLOCK_RETRY_TIMEOUT)
		} else {
			return nil, errors.New("unknown error while submitting block to NEAR")
		}
	}

	return nil, errors.New("failed to submit blocks to NEAR after retries")
}

func (r *Relayer) listenToBlocks(ctx context.Context, blockBatchC chan []*ethtypes.Block, ticker *time.Ticker) {
	headers := make(chan *ethtypes.Header)

	sub, err := r.rpcClient.SubscribeNewHead(ctx, headers)
	if err != nil {
		r.logger.Fatalf("Error subscribing to new rollup block headers: %s", err.Error())
	}
	defer sub.Unsubscribe()

	var blocks []*ethtypes.Block

	for {
		select {
		case err := <-sub.Err():
			r.logger.Error("Error on rollup block subscription", "err", err)
			return
		case header := <-headers:
			r.logger.Info("Received rollup block header", "number", header.Number.Uint64())
			r.listener.OnBlockReceived()

			blockWithNoTransactions := ethtypes.NewBlockWithHeader(header)
			blocks = append(blocks, blockWithNoTransactions)

		case <-ticker.C:
			if len(blocks) > 0 {
				blockBatchC <- blocks
				blocks = nil
				ticker.Stop()
			}

		case <-ctx.Done():
			return
		}
	}
}
