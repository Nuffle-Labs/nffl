package aggregator

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/common"

	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/config"
)

const NUM_OF_RETRIES = 5
const TX_RETRY_INTERVAL = time.Millisecond * 200

type RollupWriter struct {
	txMgr              txmgr.TxManager
	client             eth.EthClient
	sfflRegistryRollup *registryrollup.ContractSFFLRegistryRollup

	logger logging.Logger
}

func NewRollupWriter(ctx context.Context, rollupInfo config.RollupInfo, signerConfig signerv2.Config, address common.Address, logger logging.Logger) (*RollupWriter, error) {
	client, err := eth.NewClient(rollupInfo.RpcUrl)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerConfig, chainId)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(client, logger, signerV2, address)

	sfflRegistryRollup, err := registryrollup.NewContractSFFLRegistryRollup(rollupInfo.SFFLRegistryRollupAddr, client)
	if err != nil {
		return nil, err
	}

	return &RollupWriter{
		txMgr:              txMgr,
		client:             client,
		sfflRegistryRollup: sfflRegistryRollup,
		logger:             logger,
	}, nil
}

func (w *RollupWriter) UpdateOperatorSet(ctx context.Context, message registryrollup.OperatorSetUpdateMessage, signatureInfo registryrollup.OperatorsSignatureInfo) error {
	operation := func() error {
		txOpts, err := w.txMgr.GetNoSendTxOpts()
		if err != nil {
			w.logger.Error("Error getting tx opts", "err", err)
			return err
		}

		tx, err := w.sfflRegistryRollup.UpdateOperatorSet(txOpts, message, signatureInfo)
		if err != nil {
			w.logger.Error("Error assembling UpdateOperatorSet tx", "err", err)
			return err
		}

		_, err = w.txMgr.Send(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	}

	var err error = nil
	for i := 0; i < NUM_OF_RETRIES; i++ {
		err = operation()
		if err == nil {
			return nil
		} else {
			// TODO: return on same tx err
			w.logger.Warn("Sending UpdateOperatorSet failed", "err", err)
		}

		select {
		case <-ctx.Done():
			w.logger.Info("Context canceled")
			return ctx.Err()

		case <-time.After(TX_RETRY_INTERVAL):
			continue
		}
	}

	return err
}

type RollupBroadcasterer interface {
	BroadcastOperatorSetUpdate(ctx context.Context, message registryrollup.OperatorSetUpdateMessage, signatureInfo registryrollup.OperatorsSignatureInfo)
	GetErrorChan() <-chan error
}

type RollupBroadcaster struct {
	writers   []*RollupWriter
	errorChan chan error
}

func NewRollupBroadcaster(ctx context.Context, rollupsInfo map[uint32]config.RollupInfo, signerConfig signerv2.Config, address common.Address, logger logging.Logger) (*RollupBroadcaster, error) {
	writers := make([]*RollupWriter, 0, len(rollupsInfo))
	for id, info := range rollupsInfo {
		writer, err := NewRollupWriter(ctx, info, signerConfig, address, logger)
		if err != nil {
			logger.Error("Couldn't create RollupWriter", "chainId", id, "err", err)
			return nil, err
		}

		writers = append(writers, writer)
	}

	return &RollupBroadcaster{
		writers:   writers,
		errorChan: make(chan error),
	}, nil
}

func (b *RollupBroadcaster) BroadcastOperatorSetUpdate(ctx context.Context, message registryrollup.OperatorSetUpdateMessage, signatureInfo registryrollup.OperatorsSignatureInfo) {
	go func() {
		for _, writer := range b.writers {
			select {
			case <-ctx.Done():
				return

			default:
				err := writer.UpdateOperatorSet(ctx, message, signatureInfo)
				if err != nil {
					b.errorChan <- err
				}
			}
		}
	}()
}

func (b *RollupBroadcaster) GetErrorChan() <-chan error {
	return b.errorChan
}
