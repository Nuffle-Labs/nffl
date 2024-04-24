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

	"github.com/NethermindEth/near-sffl/relayer/config"
)

const (
	NAMESPACE_ID               = 1
	SUBMIT_BLOCK_INTERVAL      = 1500 * time.Millisecond
	SUBMIT_BLOCK_RETRY_TIMEOUT = 1 * time.Second
	SUBMIT_BLOCK_RETRIES       = 3
)

type Relayer struct {
	logger      sdklogging.Logger
	rpcClient   eth.Client
	rpcUrl      string
	daAccountId string

	nearClient *near.Config
}

func NewRelayerFromConfig(config *config.RelayerConfig, logger sdklogging.Logger) (*Relayer, error) {
	rpcClient, err := eth.NewClient(config.RpcUrl)
	if err != nil {
		return nil, err
	}

	nearClient, err := near.NewConfigFile(config.KeyPath, config.DaAccountId, config.Network, NAMESPACE_ID)
	if err != nil {
		return nil, err
	}

	return &Relayer{
		logger:      logger,
		rpcClient:   rpcClient,
		rpcUrl:      config.RpcUrl,
		daAccountId: config.DaAccountId,
		nearClient:  nearClient,
	}, nil
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
		return err
	}

	r.logger.Info(string(out))

	return nil
}

func (r *Relayer) submitEncodedBlocks(encodedBlocks []byte) ([]byte, error) {
	for i := 0; i < SUBMIT_BLOCK_RETRIES; i++ {
		out, err := r.nearClient.ForceSubmit(encodedBlocks)
		if err == nil {
			return out, nil
		}

		r.logger.Error("Error submitting blocks to NEAR, resubmitting", "err", err)

		if strings.Contains(err.Error(), "InvalidNonce") || strings.Contains(err.Error(), "Expired") {
			r.logger.Info("Invalid nonce or expired, resubmitting", "err", err)
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
			r.logger.Errorf("error on rollup block subscription: %s", err.Error())

			sub.Unsubscribe()

			r.rpcClient, err = eth.NewClient(r.rpcUrl)
			if err != nil {
				r.logger.Fatalf("Error reconnecting to RPC: %s", err.Error())
			}

			sub, err = r.rpcClient.SubscribeNewHead(ctx, headers)
			if err != nil {
				r.logger.Fatalf("Error resubscribing to new rollup block headers: %s", err.Error())
			}

			r.logger.Info("Resubscribed to rollup block headers")
		case header := <-headers:
			r.logger.Info("Received rollup block header", "number", header.Number.Uint64())
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
