package operator

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	regcoord "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryCoordinator"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/types"
)

type AvsManager struct {
	avsWriter     *chainio.AvsWriter
	avsReader     chainio.AvsReaderer
	avsSubscriber chainio.AvsSubscriberer
	operatorAddr  common.Address

	// TODO: agree on operatorSetUpdateC vs operatorSetUpdateChan
	// receive new tasks in this chan (typically from listening to onchain event)
	checkpointTaskCreatedChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated
	// receive operator set updates in this chan
	operatorSetUpdateChan           chan *regcoord.ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock
	signedCheckpointTaskCreatedChan chan *coretypes.SignedCheckpointTaskResponse
	signedOperatorSetUpdateChan     chan *coretypes.SignedOperatorSetUpdateMessage

	logger sdklogging.Logger
}

func NewAvsManager(config *types.NodeConfig, ethRpcClient eth.EthClient, ethWsClient eth.EthClient, signerV2 signerv2.SignerFn, logger sdklogging.Logger) (*AvsManager, error) {
	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(config.OperatorAddress))
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr, common.HexToAddress(config.AVSRegistryCoordinatorAddress),
		common.HexToAddress(config.OperatorStateRetrieverAddress), ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(config.AVSRegistryCoordinatorAddress),
		common.HexToAddress(config.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	avsSubscriber, err := chainio.BuildAvsSubscriber(common.HexToAddress(config.AVSRegistryCoordinatorAddress),
		common.HexToAddress(config.OperatorStateRetrieverAddress), ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	return &AvsManager{
		avsReader:                       avsReader,
		avsWriter:                       avsWriter,
		avsSubscriber:                   avsSubscriber,
		operatorAddr:                    common.HexToAddress(config.OperatorAddress),
		checkpointTaskCreatedChan:       make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:           make(chan *regcoord.ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock),
		signedCheckpointTaskCreatedChan: make(chan *coretypes.SignedCheckpointTaskResponse),
		signedOperatorSetUpdateChan:     make(chan *coretypes.SignedOperatorSetUpdateMessage),
	}, nil
}

func (avsManager *AvsManager) Start(ctx context.Context) error {
	operatorIsRegistered, err := avsManager.IsOperatorRegistered(&bind.CallOpts{}, avsManager.operatorAddr)
	if err != nil {
		avsManager.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	newTasksSub, err := avsManager.avsSubscriber.SubscribeToNewTasks(avsManager.checkpointTaskCreatedChan)
	if err != nil {
		avsManager.logger.Error("Error subscribing to new tasks", "err", err)
		return err
	}

	operatorSetUpdateSub, err := avsManager.avsSubscriber.SubscribeToOperatorSetUpdates(avsManager.operatorSetUpdateChan)
	if err != nil {
		avsManager.logger.Error("Error subscribing to operator set updates", "err", err)
		return err
	}

	for {
		select {
		case err := <-newTasksSub.Err():
			avsManager.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			newTasksSub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			newTasksSub, err = avsManager.avsSubscriber.SubscribeToNewTasks(avsManager.checkpointTaskCreatedChan)
			if err != nil {
				avsManager.logger.Error("Error re-subscribing to new tasks", "err", err)
				return
			}
		case checkpointTaskCreatedLog := <-avsManager.checkpointTaskCreatedChan:
			avsManager.metrics.IncNumTasksReceived()
			taskResponse := avsManager.ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog)
			signedCheckpointTaskResponse, err := avsManager.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}

			avsManager.signedCheckpointTaskCreatedChan <- signedCheckpointTaskResponse
		case err := <-operatorSetUpdateSub.Err():
			avsManager.logger.Error("Error in websocket subscription", "err", err)
			operatorSetUpdateSub.Unsubscribe()
			operatorSetUpdateSub = avsManager.avsSubscriber.SubscribeToOperatorSetUpdates(avsManager.operatorSetUpdateChan)
		case operatorSetUpdate := <-avsManager.operatorSetUpdateChan:
			go avsManager.handleOperatorSetUpdate(ctx, operatorSetUpdate)
		}
	}

}

// Takes a CheckpointTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (avsManager *AvsManager) ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) *taskmanager.CheckpointTaskResponse {
	avsManager.logger.Debug("Received new task", "task", checkpointTaskCreatedLog)
	avsManager.logger.Info("Received new task",
		"fromTimestamp", checkpointTaskCreatedLog.Task.FromTimestamp,
		"toTimestamp", checkpointTaskCreatedLog.Task.ToTimestamp,
		"taskIndex", checkpointTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", checkpointTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", checkpointTaskCreatedLog.Task.QuorumNumbers,
		"quorumThreshold", checkpointTaskCreatedLog.Task.QuorumThreshold,
	)

	// TODO: build SMT based on stored message agreements and update the test

	taskResponse := &taskmanager.CheckpointTaskResponse{
		ReferenceTaskIndex:     checkpointTaskCreatedLog.TaskIndex,
		StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *taskmanager.CheckpointTaskResponse) (*coretypes.SignedCheckpointTaskResponse, error) {
	taskResponseHash, err := core.GetCheckpointTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedCheckpointTaskResponse := &coretypes.SignedCheckpointTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedCheckpointTaskResponse", signedCheckpointTaskResponse)
	return signedCheckpointTaskResponse, nil
}

func (avsManager *AvsManager) handleOperatorSetUpdate(ctx context.Context, data *regcoord.ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock) error {
	operatorSetDelta, err := avsManager.avsReader.GetOperatorSetUpdateDelta(ctx, data.Id)
	if err != nil {
		avsManager.logger.Errorf("Couldn't get Operator set update delta: %v for block: %v", err, data.Id)
		return err
	}

	operators := make([]registryrollup.OperatorsOperator, len(operatorSetDelta))
	for i := 0; i < len(operatorSetDelta); i++ {
		operators[i] = registryrollup.OperatorsOperator{
			Pubkey: registryrollup.BN254G1Point{
				X: operatorSetDelta[i].Pubkey.X,
				Y: operatorSetDelta[i].Pubkey.Y,
			},
			Weight: operatorSetDelta[i].Weight,
		}
	}

	message := registryrollup.OperatorSetUpdateMessage{
		Id:        data.Id,
		Timestamp: data.Timestamp,
		Operators: operators,
	}

	signedMessage, err := SignOperatorSetUpdate(message, avsManager.blsKeypair, avsManager.operatorId)
	if err != nil {
		avsManager.logger.Error("Couldn't sign operator set update message", "err", err)
		return err
	}
	avsManager.logger.Debug("Signed operator set update response", "signedMessage", signedMessage)
	avsManager.signedOperatorSetUpdateChan <- signedMessage

	return nil
}

func SignOperatorSetUpdate(message registryrollup.OperatorSetUpdateMessage, blsKeyPair *bls.KeyPair, operatorId bls.OperatorId) (*coretypes.SignedOperatorSetUpdateMessage, error) {
	messageHash, err := core.GetOperatorSetUpdateMessageDigest(&message)
	if err != nil {
		return nil, err
	}
	signature := blsKeyPair.SignMessage(messageHash)
	signedOperatorSetUpdate := coretypes.SignedOperatorSetUpdateMessage{
		Message:      message,
		OperatorId:   operatorId,
		BlsSignature: *signature,
	}

	return &signedOperatorSetUpdate, nil
}

func (avsManager *AvsManager) IsOperatorRegistered(options *bind.CallOpts, address common.Address) (bool, error) {
	return avsManager.avsReader.IsOperatorRegistered(options, address)
}
