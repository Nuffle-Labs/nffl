package aggregator

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

const NUM_OF_RETRIES = 5
const TX_RETRY_INTERVAL = time.Millisecond * 200
const OPERATOR_SET_RETRY_INTERVAL = time.Millisecond * 500

type RollupWriter struct {
	txMgr                 txmgr.TxManager
	client                eth.EthClient
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

	writer := &RollupWriter{
		txMgr:              txMgr,
		client:             client,
		sfflRegistryRollup: sfflRegistryRollup,
		rollupId:           rollupId,
		logger:             logger,
	}

	return writer, nil
}

func (w *RollupWriter) InitializeOperatorSet(ctx context.Context, operators []registryrollup.RollupOperatorsOperator, mainnetNextOperatorSetUpdateId uint64) error {
	w.operatorSetUpdateLock.Lock()
	defer w.operatorSetUpdateLock.Unlock()

	w.logger.Info("Initializing operator set")

	operation := func() error {
		txOpts, err := w.txMgr.GetNoSendTxOpts()
		if err != nil {
			w.logger.Error("Error getting tx opts", "err", err)
			return err
		}

		tx, err := w.sfflRegistryRollup.SetInitialOperatorSet(txOpts, operators, mainnetNextOperatorSetUpdateId)
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

	for i := 0; i < NUM_OF_RETRIES; i++ {
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

		case <-time.After(TX_RETRY_INTERVAL):
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

	for i := 0; i < NUM_OF_RETRIES; i++ {
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

		case <-time.After(TX_RETRY_INTERVAL):
			continue
		}
	}

	return errors.New("failed to update operator set after retries")
}

type RollupBroadcasterer interface {
	BroadcastOperatorSetUpdate(ctx context.Context, message messages.OperatorSetUpdateMessage, signatureInfo registryrollup.RollupOperatorsSignatureInfo)
	GetErrorChan() <-chan error
}

type RollupBroadcaster struct {
	writers   []*RollupWriter
	logger    logging.Logger
	errorChan chan error
}

func NewRollupBroadcaster(
	ctx context.Context,
	avsReader chainio.AvsReaderer,
	avsSubscriber chainio.AvsSubscriberer,
	rollupsInfo map[uint32]config.RollupInfo,
	signerConfig signerv2.Config,
	address common.Address,
	logger logging.Logger,
) (*RollupBroadcaster, error) {
	writers := make([]*RollupWriter, 0, len(rollupsInfo))

	for id, info := range rollupsInfo {
		writer, err := NewRollupWriter(ctx, id, info, signerConfig, address, logger)
		if err != nil {
			logger.Error("Couldn't create RollupWriter", "chainId", id, "err", err)
			return nil, err
		}

		writers = append(writers, writer)
	}

	broadcaster := &RollupBroadcaster{
		writers:   writers,
		logger:    logger,
		errorChan: make(chan error),
	}

	mainnetNextOperatorSetUpdateId, err := avsReader.GetNextOperatorSetUpdateId(ctx)
	if err != nil {
		logger.Error("Error fetching operator set update id", "err", err)
		return nil, err
	}

	if mainnetNextOperatorSetUpdateId == 0 {
		go broadcaster.initializeRollupOperatorSetsOnUpdate(ctx, avsReader, avsSubscriber)
	} else {
		for _, writer := range broadcaster.writers {
			go broadcaster.tryInitializeRollupOperatorSet(ctx, writer, mainnetNextOperatorSetUpdateId, avsReader, avsSubscriber)
		}
	}

	return broadcaster, nil
}

func (b *RollupBroadcaster) initializeRollupOperatorSetsOnUpdate(ctx context.Context, avsReader chainio.AvsReaderer, avsSubscriber chainio.AvsSubscriberer) {
	var operatorSetUpdatedChan chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock

	avsSubscriber.SubscribeToOperatorSetUpdates(operatorSetUpdatedChan)

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-operatorSetUpdatedChan:
			b.logger.Info("Received operator set update", "id", event.Id)

			operators, err := avsReader.GetOperatorSetById(ctx, event.Id)
			if err != nil {
				b.errorChan <- err
				continue
			}

			convertedOperators := make([]registryrollup.RollupOperatorsOperator, len(operators))
			for i, op := range operators {
				convertedOperators[i] = registryrollup.RollupOperatorsOperator{
					Pubkey: registryrollup.BN254G1Point{X: op.Pubkey.X, Y: op.Pubkey.Y},
					Weight: op.Weight,
				}
			}

			for _, writer := range b.writers {
				go func(writer *RollupWriter) {
					err := writer.InitializeOperatorSet(ctx, convertedOperators, event.Id)
					if err != nil {
						b.errorChan <- err
					}
				}(writer)
			}
		}
	}
}

func (b *RollupBroadcaster) tryInitializeRollupOperatorSet(ctx context.Context, writer *RollupWriter, mainnetNextOperatorSetUpdateId uint64, avsReader chainio.AvsReaderer, avsSubscriber chainio.AvsSubscriberer) {
	nextOperatorUpdateId, err := writer.sfflRegistryRollup.NextOperatorUpdateId(&bind.CallOpts{})
	if err != nil {
		b.logger.Error("Error fetching NextOperatorUpdateId", "err", err)
		b.errorChan <- err
		return
	}

	if nextOperatorUpdateId == 0 {
		b.logger.Info("Operator set not initialized yet", "rollupId", writer.rollupId, "mainnetNextOperatorSetUpdateId", mainnetNextOperatorSetUpdateId)

		operators, err := b.tryGetOperatorSetById(ctx, avsReader, mainnetNextOperatorSetUpdateId-1)
		if err != nil {
			b.logger.Error("Error fetching operator set", "err", err)
			b.errorChan <- err
			return
		}

		convertedOperators := make([]registryrollup.RollupOperatorsOperator, len(operators))
		for i, op := range operators {
			convertedOperators[i] = registryrollup.RollupOperatorsOperator{
				Pubkey: registryrollup.BN254G1Point{X: op.Pubkey.X, Y: op.Pubkey.Y},
				Weight: op.Weight,
			}
		}

		err = writer.InitializeOperatorSet(ctx, convertedOperators, mainnetNextOperatorSetUpdateId-1)
		if err != nil {
			b.logger.Error("Error initializing operator set", "err", err)
			b.errorChan <- err
		}
	}
}

func (b *RollupBroadcaster) BroadcastOperatorSetUpdate(ctx context.Context, message messages.OperatorSetUpdateMessage, signatureInfo registryrollup.RollupOperatorsSignatureInfo) {
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

func (b *RollupBroadcaster) tryGetOperatorSetById(ctx context.Context, avsReader chainio.AvsReaderer, operatorSetUpdateId uint64) ([]opsetupdatereg.RollupOperatorsOperator, error) {
	for i := 0; i < NUM_OF_RETRIES; i++ {
		operators, err := avsReader.GetOperatorSetById(ctx, operatorSetUpdateId)

		if err == nil {
			return operators, nil
		}

		b.errorChan <- err

		select {
		case <-ctx.Done():
			b.logger.Info("Context canceled")
			return nil, errors.New("failed to fetch operator set early")

		case <-time.After(OPERATOR_SET_RETRY_INTERVAL):
			continue
		}
	}

	return nil, errors.New("failed to fetch operator set after retries")
}
