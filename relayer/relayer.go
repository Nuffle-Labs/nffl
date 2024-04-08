package relayer

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	near "github.com/near/rollup-data-availability/gopkg/da-rpc"
)

const (
	NAMESPACE_ID = 1
)

type Relayer struct {
	logger      sdklogging.Logger
	rpcClient   eth.EthClient
	daAccountId string

	nearClient *near.Config
}

type RelayerConfig struct {
	Production  bool
	RpcUrl      string
	DaAccountId string
	KeyPath     string
	Network     string
}

func (c RelayerConfig) CompileCMD() []string {
	var cmd []string
	if c.Production {
		cmd = append(cmd, "--production")
	}

	cmd = append(cmd, "--key-path", c.KeyPath)
	cmd = append(cmd, "--rpc-url", c.RpcUrl)
	cmd = append(cmd, "--da-account-id", c.DaAccountId)
	cmd = append(cmd, "--network", c.Network)
	return cmd
}

func NewRelayerFromConfig(config *RelayerConfig, logger sdklogging.Logger) (*Relayer, error) {
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

	for {
		select {
		case err := <-sub.Err():
			r.logger.Errorf("error on rollup block subscription: %s", err.Error())
		case header := <-headers:
			block := r.getBlockByNumber(ctx, header.Number)
			if block == nil {
				return nil
			}

			encodedBlock, err := rlp.EncodeToBytes(block)
			if err != nil {
				r.logger.Errorf("error RLP encoding block: %s", err.Error())
				return err
			}

			out, err := r.nearClient.ForceSubmit(encodedBlock)
			if err != nil {
				r.logger.Error("Error submitting block to NEAR", "err", err)
				return err
			}

			r.logger.Info(string(out))
		case <-ctx.Done():
			return nil
		}
	}
}

func (r *Relayer) getBlockByNumber(ctx context.Context, number *big.Int) *ethtypes.Block {
	r.logger.Infof("Got rollup block: #%s", number.String())

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		block, err := r.rpcClient.BlockByNumber(ctx, number)
		if err != nil {
			r.logger.Errorf("Error fetching rollup block: %s", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		return block
	}

	panic("Could not fetch rollup block")
}
