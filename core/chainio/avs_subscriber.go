package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

type AvsSubscriberer interface {
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
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.EthClient, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
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
	s.logger.Infof("Subscribed to new TaskManager tasks")
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
	s.logger.Infof("Subscribed to CheckpointTaskResponded events")
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
	s.logger.Infof("Subscribed to OperatorSetUpdatedAtBlock events")
	return sub, nil
}
