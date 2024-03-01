package aggregator

import (
	"context"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	regrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

const NUM_OF_RETRIES = 5
const TX_RETRY_INTERVAL = time.Millisecond * 200

type RollupWriter struct {
	txMgr              txmgr.TxManager
	client             eth.EthClient
	sfflRegistryRollup *regrollup.ContractSFFLRegistryRollup

	logger logging.Logger
}

func NewRollupWriter(rollupInfo config.RollupInfo, signerConfig signerv2.Config, address common.Address, logger logging.Logger, ctx context.Context) (*RollupWriter, error) {
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

	sfflRegistryRollup, err := regrollup.NewContractSFFLRegistryRollup(rollupInfo.SFFLRegistryRollupAddr, client)
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

func (writer *RollupWriter) UpdateOperatorSet(ctx context.Context, message regrollup.OperatorSetUpdateMessage, signatureInfo regrollup.OperatorsSignatureInfo) error {
	operation := func() error {
		txOpts, err := writer.txMgr.GetNoSendTxOpts()
		if err != nil {
			writer.logger.Error("Error getting tx opts", "err", err)
			return err
		}

		tx, err := writer.sfflRegistryRollup.UpdateOperatorSet(txOpts, message, signatureInfo)
		if err != nil {
			writer.logger.Error("Error assembling UpdateOperatorSet tx", "err", err)
			return err
		}

		_, err = writer.txMgr.Send(ctx, tx)
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
		}

		select {
		case <-ctx.Done():
			writer.logger.Info("Context canceled")
			return ctx.Err()

		case <-time.After(TX_RETRY_INTERVAL):
			continue
		}
	}

	return err
}

type RollupBroadcaster struct {
	writers   []*RollupWriter
	errorChan chan error
}

func NewRollupBroadcaster(rollupsInfo map[uint32]config.RollupInfo, signerConfig signerv2.Config, address common.Address, logger logging.Logger, ctx context.Context) (*RollupBroadcaster, error) {
	writers := make([]*RollupWriter, 0, len(rollupsInfo))
	for id, info := range rollupsInfo {
		writer, err := NewRollupWriter(info, signerConfig, address, logger, ctx)
		if err != nil {
			logger.Errorf("Couldn't create RollupWriter for chainId: %d, error: %s", id, err.Error())
			return nil, err
		}

		writers = append(writers, writer)
	}

	return &RollupBroadcaster{
		writers:   writers,
		errorChan: make(chan error),
	}, nil
}

func (broadcaster *RollupBroadcaster) BroadcastOperatorSetUpdate(ctx context.Context, message regrollup.OperatorSetUpdateMessage, signatureInfo regrollup.OperatorsSignatureInfo) {
	go func() {
		for _, writer := range broadcaster.writers {
			select {
			case <-ctx.Done():
				return

			default:
				err := writer.UpdateOperatorSet(ctx, message, signatureInfo)
				if err != nil {
					broadcaster.errorChan <- err
				}
			}
		}
	}()
}

func (broadcaster *RollupBroadcaster) GetErrorChan() <-chan error {
	return broadcaster.errorChan
}
