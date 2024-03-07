package chainio

import (
	"context"
	"github.com/NethermindEth/near-sffl/core/config"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewCheckpointTask(
		ctx context.Context,
		fromTimestamp uint64,
		toTimestamp uint64,
		quorumThreshold uint32,
		quorumNumbers []byte,
	) (taskmanager.CheckpointTask, uint32, error)
	RaiseChallenge(
		ctx context.Context,
		task taskmanager.CheckpointTask,
		taskResponse taskmanager.CheckpointTaskResponse,
		taskResponseMetadata taskmanager.CheckpointTaskResponseMetadata,
		pubkeysOfNonSigningOperators []taskmanager.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task taskmanager.CheckpointTask,
		taskResponse taskmanager.CheckpointTaskResponse,
		nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.EthClient
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(txMgr txmgr.TxManager, config *config.Config, client eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	return BuildAvsWriter(txMgr, config.SFFLRegistryCoordinatorAddr, config.OperatorStateRetrieverAddr, client, logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}

func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewCheckpointTask(ctx context.Context, fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (taskmanager.CheckpointTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return taskmanager.CheckpointTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateCheckpointTask(txOpts, fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
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
	taskResponse taskmanager.CheckpointTaskResponse,
	nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToCheckpointTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
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
	taskResponse taskmanager.CheckpointTaskResponse,
	taskResponseMetadata taskmanager.CheckpointTaskResponseMetadata,
	pubkeysOfNonSigningOperators []taskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveCheckpointChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
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
