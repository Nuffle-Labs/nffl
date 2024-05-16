package aggregator

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/core/safeclient"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

const (
	INITIALIZE_OPERATOR_SET_RETRIES        = 5
	INITIALIZE_OPERATOR_SET_RETRY_INTERVAL = 500 * time.Millisecond
	UPDATE_OPERATOR_SET_RETRIES            = 5
	UPDATE_OPERATOR_SET_RETRY_INTERVAL     = 500 * time.Millisecond
)

type RollupWriter struct {
	txMgr                 txmgr.TxManager
	client                safeclient.SafeClient
	sfflRegistryRollup    *registryrollup.ContractSFFLRegistryRollup
	rollupId              uint32
	operatorSetUpdateLock sync.Mutex

	logger logging.Logger
}

func NewRollupWriter(
	ctx context.Context,
	rollupId uint32,
	rollupInfo config.RollupInfo,
	signerConfig signerv2.Config,
	address common.Address,
	logger logging.Logger,
) (*RollupWriter, error) {
	client, err := safeclient.NewSafeEthClient(rollupInfo.RpcUrl, logger)
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

	txSender, err := wallet.NewPrivateKeyWallet(client, signerV2, address, logger)
	if err != nil {
		logger.Error("Failed to create transaction sender", "err", err)
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(txSender, client, logger, address).WithGasLimitMultiplier(1.5)

	sfflRegistryRollup, err := registryrollup.NewContractSFFLRegistryRollup(rollupInfo.SFFLRegistryRollupAddr, client)
	if err != nil {
		return nil, err
	}

	writer := &RollupWriter{
		txMgr:              txMgr,
		client:             client,
		sfflRegistryRollup: sfflRegistryRollup,
		rollupId:           rollupId,
		logger:             logger,
	}

	return writer, nil
}

func (w *RollupWriter) InitializeOperatorSet(ctx context.Context, operators []registryrollup.RollupOperatorsOperator, operatorSetUpdateId uint64) error {
	w.operatorSetUpdateLock.Lock()
	defer w.operatorSetUpdateLock.Unlock()

	w.logger.Info("Initializing operator set")

	operation := func() error {
		txOpts, err := w.txMgr.GetNoSendTxOpts()
		if err != nil {
			w.logger.Error("Error getting tx opts", "err", err)
			return err
		}

		tx, err := w.sfflRegistryRollup.SetInitialOperatorSet(txOpts, operators, operatorSetUpdateId+1)
		if err != nil {
			w.logger.Error("Error assembling SetInitialOperatorSet tx", "err", err)
			return err
		}

		_, err = w.txMgr.Send(ctx, tx)
		if err != nil {
			w.logger.Error("Error sending SetInitialOperatorSet tx", "err", err)
			return err
		}

		return nil
	}

	for i := 0; i < INITIALIZE_OPERATOR_SET_RETRIES; i++ {
		err := operation()
		if err == nil {
			w.logger.Info("Rollup Operator set initialized")

			return nil
		} else {
			// TODO: return on some tx errors
			w.logger.Warn("Sending SetInitialOperatorSet failed", "err", err)
		}

		select {
		case <-ctx.Done():
			w.logger.Info("Context canceled")
			return ctx.Err()

		case <-time.After(INITIALIZE_OPERATOR_SET_RETRY_INTERVAL):
			continue
		}
	}

	return errors.New("failed to initialize operator set after retries")
}

func (w *RollupWriter) UpdateOperatorSet(ctx context.Context, message messages.OperatorSetUpdateMessage, signatureInfo registryrollup.RollupOperatorsSignatureInfo) error {
	w.operatorSetUpdateLock.Lock()
	defer w.operatorSetUpdateLock.Unlock()

	operation := func() error {
		txOpts, err := w.txMgr.GetNoSendTxOpts()
		if err != nil {
			w.logger.Error("Error getting tx opts", "err", err)
			return err
		}

		nextOperatorUpdateId, err := w.sfflRegistryRollup.NextOperatorUpdateId(&bind.CallOpts{})
		if err != nil {
			w.logger.Error("Error fetching NextOperatorUpdateId", "err", err)
			return err
		}

		// TODO: queue in case message.id > nextOperatorUpdateId
		if message.Id != nextOperatorUpdateId {
			w.logger.Warn("MessageId didn't match nextOperatorUpdateId", "id", message.Id, "nextOperatorUpdateId", nextOperatorUpdateId)
			return nil
		}

		tx, err := w.sfflRegistryRollup.UpdateOperatorSet(txOpts, message.ToBinding(), signatureInfo)
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

	for i := 0; i < UPDATE_OPERATOR_SET_RETRIES; i++ {
		err := operation()
		if err == nil {
			w.logger.Info("Rollup Operator set updated")

			return nil
		} else {
			// TODO: return on some tx errors
			w.logger.Warn("Sending UpdateOperatorSet failed", "err", err)
		}

		select {
		case <-ctx.Done():
			w.logger.Info("Context canceled")
			return ctx.Err()

		case <-time.After(UPDATE_OPERATOR_SET_RETRY_INTERVAL):
			continue
		}
	}

	return errors.New("failed to update operator set after retries")
}

func (w *RollupWriter) Close() {
	w.client.Close()
}
