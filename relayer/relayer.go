package relayer

import (
	"context"
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
			blockNumbers := make([]uint64, len(blocks))
			for i, block := range blocks {
				blockNumbers[i] = block.Number().Uint64()
			}
			r.logger.Info("Submitting blocks to NEAR", "numbers", blockNumbers)

			encodedBlocks, err := rlp.EncodeToBytes(blocks)
			if err != nil {
				r.logger.Error("Error RLP encoding block", "err", err.Error())
				continue
			}

			out, err := r.nearClient.ForceSubmit(encodedBlocks)
			if err != nil {
				r.logger.Error("Error submitting block to NEAR", "err", err)

				if strings.Contains(err.Error(), "InvalidNonce") {
					r.logger.Info("Invalid nonce, resubmitting")
					time.Sleep(SUBMIT_BLOCK_RETRY_TIMEOUT)

					out, err = r.nearClient.ForceSubmit(encodedBlocks)
					if err != nil {
						r.logger.Error("Error resubmitting block to NEAR", "err", err)
					} else {
						r.logger.Info(string(out))
					}
				}

				return err
			} else {
				r.logger.Info(string(out))
			}

			ticker.Reset(SUBMIT_BLOCK_INTERVAL)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
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
