package relayer

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	near "github.com/near/rollup-data-availability/gopkg/da-rpc"

	"github.com/NethermindEth/near-sffl/relayer/config"
)

const (
	NAMESPACE_ID = 1
)

type Relayer struct {
	logger      sdklogging.Logger
	rpcClient   eth.Client
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
		daAccountId: config.DaAccountId,
		nearClient:  nearClient,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) error {
	headers := make(chan *ethtypes.Header)
	sub, err := r.rpcClient.SubscribeNewHead(ctx, headers)
	if err != nil {
		r.logger.Fatalf("Error subscribing to new rollup block headers: %s", err.Error())
	}
	defer sub.Unsubscribe()

	ticker := time.NewTicker(1500 * time.Millisecond)
	defer ticker.Stop()

	var blocks []*ethtypes.Block
	for {
		select {
		case err := <-sub.Err():
			r.logger.Errorf("error on rollup block subscription: %s", err.Error())
			return err
		case header := <-headers:
			r.logger.Info("Received rollup block header", "number", header.Number.Uint64())
			blockWithNoTransactions := ethtypes.NewBlockWithHeader(header)
			blocks = append(blocks, blockWithNoTransactions)
			continue
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		select {
		case <-ticker.C:
			if len(blocks) == 0 {
				continue
			}

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

			blocks = blocks[:0]

			out, err := r.nearClient.ForceSubmit(encodedBlocks)
			if err != nil {
				r.logger.Error("Error submitting block to NEAR", "err", err)
				return err
			}

			r.logger.Info(string(out))
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
