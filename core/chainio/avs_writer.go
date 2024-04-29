package chainio

import (
	"context"
	"time"

	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/core/types/messages"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewCheckpointTask(
		ctx context.Context,
		fromTimestamp uint64,
		toTimestamp uint64,
		quorumThreshold eigentypes.QuorumThresholdPercentage,
		quorumNumbers eigentypes.QuorumNums,
	) (taskmanager.CheckpointTask, uint32, error)
	RaiseChallenge(
		ctx context.Context,
		task taskmanager.CheckpointTask,
		taskResponse messages.CheckpointTaskResponse,
		taskResponseMetadata taskmanager.CheckpointTaskResponseMetadata,
		pubkeysOfNonSigningOperators []taskmanager.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task taskmanager.CheckpointTask,
		taskResponse messages.CheckpointTaskResponse,
		aggregation messages.MessageBlsAggregation,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(txMgr txmgr.TxManager, config *config.Config, client eth.Client, logger logging.Logger) (*AvsWriter, error) {
	return BuildAvsWriter(txMgr, config.SFFLRegistryCoordinatorAddr, config.OperatorStateRetrieverAddr, client, logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, ethHttpClient, logger, txMgr), nil
}

func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, client eth.Client, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		TxMgr:               txMgr,
		client:              client,
		logger:              logger,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewCheckpointTask(ctx context.Context, fromTimestamp uint64, toTimestamp uint64, quorumThreshold eigentypes.QuorumThresholdPercentage, quorumNumbers eigentypes.QuorumNums) (taskmanager.CheckpointTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return taskmanager.CheckpointTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateCheckpointTask(txOpts, fromTimestamp, toTimestamp, uint32(quorumThreshold), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Errorf("Error assembling CreateCheckpointTask tx")
		return taskmanager.CheckpointTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateCheckpointTask tx")
		return taskmanager.CheckpointTask{}, 0, err
	}
	checkpointTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractSFFLTaskManagerFilterer.ParseCheckpointTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return taskmanager.CheckpointTask{}, 0, err
	}
	return checkpointTaskCreatedEvent.Task, checkpointTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task taskmanager.CheckpointTask,
	taskResponse messages.CheckpointTaskResponse,
	aggregation messages.MessageBlsAggregation,
) (*types.Receipt, error) {
	// Wait a block if the task TaskCreatedBlock is the same as the current block
	currentBlock, err := w.client.BlockNumber(ctx)
	if err != nil {
		w.logger.Errorf("Error getting current block number")
		return nil, err
	}

	if uint64(task.TaskCreatedBlock) == currentBlock {
		w.logger.Info("Waiting roughly a block before sending aggregated response...")
		time.Sleep(20 * time.Second)
	}

	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}

	tx, err := w.AvsContractBindings.TaskManager.RespondToCheckpointTask(txOpts, task, taskResponse.ToBinding(), aggregation.ExtractBindingMainnet())
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}

	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateCheckpointTask tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task taskmanager.CheckpointTask,
	taskResponse messages.CheckpointTaskResponse,
	taskResponseMetadata taskmanager.CheckpointTaskResponseMetadata,
	pubkeysOfNonSigningOperators []taskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveCheckpointChallenge(txOpts, task, taskResponse.ToBinding(), taskResponseMetadata, pubkeysOfNonSigningOperators)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateCheckpointTask tx")
		return nil, err
	}
	return receipt, nil
}
