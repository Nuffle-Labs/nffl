package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	opsetupdatereg "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	taskmanager "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLTaskManager"
)

type AvsSubscriberer interface {
	avsregistry.AvsRegistrySubscriber
	SubscribeToNewTasks(checkpointTaskCreatedChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) (event.Subscription, error)
	SubscribeToTaskResponses(taskResponseLogs chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskResponded) (event.Subscription, error)
	SubscribeToOperatorSetUpdates(operatorSetUpdateChan chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock) (event.Subscription, error)
	ParseCheckpointTaskResponded(rawLog types.Log) (*taskmanager.ContractSFFLTaskManagerCheckpointTaskResponded, error)
}

// Subscribers use a ws connection instead of http connection like Readers
// kind of stupid that the geth client doesn't have a unified interface for both...
// it takes a single url, so the bindings, even though they have watcher functions, those can't be used
// with the http connection... seems very very stupid. Am I missing something?
type AvsSubscriber struct {
	avsregistry.AvsRegistrySubscriber
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

var _ (AvsSubscriberer) = (*AvsSubscriber)(nil)

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.Client, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}

	avsRegistrySubscriber, err := avsregistry.NewAvsRegistryChainSubscriber(logger, avsContractBindings.RegistryCoordinator, avsContractBindings.BlsApkRegistry)
	if err != nil {
		logger.Error("Failed to create chain registry subscriber", "err", err)
		return nil, err
	}

	return NewAvsSubscriber(avsContractBindings, avsRegistrySubscriber, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, avsRegistrySubscriber avsregistry.AvsRegistrySubscriber, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsRegistrySubscriber: avsRegistrySubscriber,
		AvsContractBindings:   avsContractBindings,
		logger:                logger,
	}
}

func (s *AvsSubscriber) SubscribeToNewTasks(checkpointTaskCreatedChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) (event.Subscription, error) {
	sub, err := s.AvsContractBindings.TaskManager.WatchCheckpointTaskCreated(
		&bind.WatchOpts{}, checkpointTaskCreatedChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new TaskManager tasks", "err", err)
		return nil, err
	}
	s.logger.Info("Subscribed to new TaskManager tasks")
	return sub, nil
}

func (s *AvsSubscriber) SubscribeToTaskResponses(taskResponseChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskResponded) (event.Subscription, error) {
	sub, err := s.AvsContractBindings.TaskManager.WatchCheckpointTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to CheckpointTaskResponded events", "err", err)
		return nil, err
	}
	s.logger.Info("Subscribed to CheckpointTaskResponded events")
	return sub, nil
}

func (s *AvsSubscriber) ParseCheckpointTaskResponded(rawLog types.Log) (*taskmanager.ContractSFFLTaskManagerCheckpointTaskResponded, error) {
	return s.AvsContractBindings.TaskManager.ContractSFFLTaskManagerFilterer.ParseCheckpointTaskResponded(rawLog)
}

func (s *AvsSubscriber) SubscribeToOperatorSetUpdates(operatorSetUpdateChan chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock) (event.Subscription, error) {
	sub, err := s.AvsContractBindings.OperatorSetUpdateRegistry.WatchOperatorSetUpdatedAtBlock(
		&bind.WatchOpts{}, operatorSetUpdateChan, nil, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to OperatorSetUpdatedAtBlock events", "err", err)
		return nil, err
	}
	s.logger.Info("Subscribed to OperatorSetUpdatedAtBlock events")
	return sub, nil
}
