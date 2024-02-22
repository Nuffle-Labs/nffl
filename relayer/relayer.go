package relayer

import (
	"context"
	"encoding/base64"
	"os"
	"os/exec"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Relayer struct {
	logger      sdklogging.Logger
	rpcClient   eth.EthClient
	daAccountId string
}

type RelayerConfig struct {
	Production  bool
	RpcUrl      string
	DaAccountId string
}

func NewRelayerFromConfig(config *RelayerConfig) (*Relayer, error) {
	var logLevel sdklogging.LogLevel
	if config.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}

	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}

	rpcClient, err := eth.NewClient(config.RpcUrl)
	if err != nil {
		return nil, err
	}

	return &Relayer{
		logger:      logger,
		rpcClient:   rpcClient,
		daAccountId: config.DaAccountId,
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
			r.logger.Infof("Got rollup block: #%s", header.Number.String())

			block, err := r.rpcClient.BlockByNumber(ctx, header.Number)
			if err != nil {
				r.logger.Errorf("error getting rollup block: %s", err.Error())
				return err
			}

			encodedBlock, err := rlp.EncodeToBytes(block)
			if err != nil {
				r.logger.Errorf("error RLP encoding block: %s", err.Error())
				return err
			}

			cmd := exec.Command("near", "call", r.daAccountId, "submit", "--base64", base64.StdEncoding.EncodeToString(encodedBlock), "--accountId", r.daAccountId)
			cmd.Env = os.Environ()
			out, err := cmd.CombinedOutput()

			r.logger.Info(string(out))
			if err != nil {
				r.logger.Errorf("error calling NEAR account: %s", err.Error())
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}
