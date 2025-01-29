package chainio

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Nuffle-Labs/nffl/core/config"
	"github.com/Nuffle-Labs/nffl/core/types/messages"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	smbase "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	taskmanager "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLTaskManager"
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
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsRegistryChainWriter(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	logger logging.Logger,
	ethClient eth.Client,
	txMgr txmgr.TxManager,
) (*avsregistry.AvsRegistryChainWriter, error) {
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create RegistryCoordinator contract", err)
	}
	operatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create OperatorStateRetriever contract", err)
	}
	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get ServiceManager address", err)
	}
	serviceManager, err := smbase.NewContractServiceManagerBase(serviceManagerAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create ServiceManager contract", err)
	}
	blsApkRegistryAddr, err := registryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get BLSApkRegistry address", err)
	}
	blsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(blsApkRegistryAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create BLSApkRegistry contract", err)
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get StakeRegistry address", err)
	}
	stakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create StakeRegistry contract", err)
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get DelegationManager address", err)
	}
	avsDirectoryAddr, err := serviceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to get AvsDirectory address", err)
	}
	elReader, err := BuildELChainReader(delegationManagerAddr, avsDirectoryAddr, ethClient, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create ELChainReader", err)
	}
	return avsregistry.NewAvsRegistryChainWriter(
		serviceManagerAddr,
		registryCoordinator,
		operatorStateRetriever,
		stakeRegistry,
		blsApkRegistry,
		elReader,
		logger,
		ethClient,
		txMgr,
	)
}

func BuildAvsWriterFromConfig(txMgr txmgr.TxManager, config *config.Config, client eth.Client, logger logging.Logger) (*AvsWriter, error) {
	return BuildAvsWriter(txMgr, config.SFFLRegistryCoordinatorAddr, config.OperatorStateRetrieverAddr, client, logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	// avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)

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
		logger:              logger,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewCheckpointTask(ctx context.Context, fromTimestamp uint64, toTimestamp uint64, quorumThreshold eigentypes.QuorumThresholdPercentage, quorumNumbers eigentypes.QuorumNums) (taskmanager.CheckpointTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Error("Error getting tx opts")
		return taskmanager.CheckpointTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateCheckpointTask(txOpts, fromTimestamp, toTimestamp, uint32(quorumThreshold), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Error("Error assembling CreateCheckpointTask tx")
		return taskmanager.CheckpointTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Error("Error submitting CreateCheckpointTask tx")
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
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Error("Error getting tx opts")
		return nil, err
	}

	tx, err := w.AvsContractBindings.TaskManager.RespondToCheckpointTask(txOpts, task, taskResponse.ToBinding(), aggregation.ExtractBindingMainnet())
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}

	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Error("Error submitting CreateCheckpointTask tx")
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
		w.logger.Error("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveCheckpointChallenge(txOpts, task, taskResponse.ToBinding(), taskResponseMetadata, pubkeysOfNonSigningOperators)
	if err != nil {
		w.logger.Error("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Error("Error submitting CreateCheckpointTask tx")
		return nil, err
	}
	return receipt, nil
}
