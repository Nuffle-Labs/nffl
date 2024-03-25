package operator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core/chainio"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
)

type AvsManagerer interface {
	Start(ctx context.Context, operatorAddr common.Address) error
	DepositIntoStrategy(operatorAddr common.Address, strategyAddr common.Address, amount *big.Int) error
	RegisterOperatorWithEigenlayer(operatorAddr common.Address) error
	RegisterOperatorWithAvs(
		client eth.EthClient,
		operatorEcdsaKeyPair *ecdsa.PrivateKey,
		blsKeyPair *bls.KeyPair,
	) error
	ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) messages.CheckpointTaskResponse

	GetOperatorId(options *bind.CallOpts, address common.Address) ([32]byte, error)
	GetCheckpointTaskResponseChan() <-chan messages.CheckpointTaskResponse
	GetOperatorSetUpdateChan() <-chan messages.OperatorSetUpdateMessage
}

type AvsManager struct {
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader elcontracts.ELReader
	eigenlayerWriter elcontracts.ELWriter

	// receive new tasks in this chan (typically from listening to onchain event)
	checkpointTaskCreatedChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated
	// receive operator set updates in this chan
	operatorSetUpdateChan chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock

	// Sends message for operator to sign
	checkpointTaskResponseCreatedChan chan messages.CheckpointTaskResponse
	operatorSetUpdateMessageChan      chan messages.OperatorSetUpdateMessage

	logger sdklogging.Logger
}

func NewAvsManager(config *optypes.NodeConfig, ethRpcClient eth.EthClient, ethWsClient eth.EthClient, sdkClients *clients.Clients, txManager *txmgr.SimpleTxManager, logger sdklogging.Logger) (*AvsManager, error) {
	avsWriter, err := chainio.BuildAvsWriter(
		txManager, common.HexToAddress(config.AVSRegistryCoordinatorAddress),
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
		avsReader:                         avsReader,
		avsWriter:                         avsWriter,
		avsSubscriber:                     avsSubscriber,
		eigenlayerReader:                  sdkClients.ElChainReader,
		eigenlayerWriter:                  sdkClients.ElChainWriter,
		checkpointTaskCreatedChan:         make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:             make(chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock),
		checkpointTaskResponseCreatedChan: make(chan messages.CheckpointTaskResponse),
		operatorSetUpdateMessageChan:      make(chan messages.OperatorSetUpdateMessage),
		logger:                            logger,
	}, nil
}

func (avsManager *AvsManager) GetCheckpointTaskResponseChan() <-chan messages.CheckpointTaskResponse {
	return avsManager.checkpointTaskResponseCreatedChan
}

func (avsManager *AvsManager) GetOperatorSetUpdateChan() <-chan messages.OperatorSetUpdateMessage {
	return avsManager.operatorSetUpdateMessageChan
}

func (avsManager *AvsManager) Start(ctx context.Context, operatorAddr common.Address) error {
	operatorIsRegistered, err := avsManager.IsOperatorRegistered(&bind.CallOpts{}, operatorAddr)
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

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case err := <-newTasksSub.Err():
				avsManager.logger.Error("Error in websocket subscription", "err", err)
				// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
				newTasksSub.Unsubscribe()
				// TODO(samlaf): wrap this call with increase in avs-node-spec metric
				newTasksSub, err = avsManager.avsSubscriber.SubscribeToNewTasks(avsManager.checkpointTaskCreatedChan)
				if err != nil {
					avsManager.logger.Error("Error re-subscribing to new tasks", "err", err)
					close(avsManager.checkpointTaskResponseCreatedChan)
					return
				}

				continue

			case checkpointTaskCreatedLog := <-avsManager.checkpointTaskCreatedChan:
				taskResponse := avsManager.ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog)
				avsManager.checkpointTaskResponseCreatedChan <- taskResponse

			case err := <-operatorSetUpdateSub.Err():
				avsManager.logger.Error("Error in websocket subscription", "err", err)
				operatorSetUpdateSub.Unsubscribe()
				operatorSetUpdateSub, err = avsManager.avsSubscriber.SubscribeToOperatorSetUpdates(avsManager.operatorSetUpdateChan)
				if err != nil {
					avsManager.logger.Error("Error re-subscribing to operator set updates", "err", err)
					close(avsManager.checkpointTaskResponseCreatedChan)
					return
				}

				continue

			case operatorSetUpdate := <-avsManager.operatorSetUpdateChan:
				go avsManager.handleOperatorSetUpdate(ctx, operatorSetUpdate)
				continue
			}
		}
	}()

	return nil
}

// Takes a CheckpointTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (avsManager *AvsManager) ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) messages.CheckpointTaskResponse {
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

	taskResponse := messages.CheckpointTaskResponse{
		ReferenceTaskIndex:     checkpointTaskCreatedLog.TaskIndex,
		StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	return taskResponse
}

func (avsManager *AvsManager) handleOperatorSetUpdate(ctx context.Context, data *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock) error {
	operatorSetDelta, err := avsManager.avsReader.GetOperatorSetUpdateDelta(ctx, data.Id)
	if err != nil {
		avsManager.logger.Errorf("Couldn't get Operator set update delta: %v for block: %v", err, data.Id)
		return err
	}

	operators := make([]coretypes.RollupOperator, 0, len(operatorSetDelta))
	for i := 0; i < len(operatorSetDelta); i++ {
		operators = append(operators, coretypes.RollupOperator{
			Pubkey: bls.NewG1Point(operatorSetDelta[i].Pubkey.X, operatorSetDelta[i].Pubkey.Y),
			Weight: operatorSetDelta[i].Weight,
		})
	}

	message := messages.OperatorSetUpdateMessage{
		Id:        data.Id,
		Timestamp: data.Timestamp,
		Operators: operators,
	}

	avsManager.operatorSetUpdateMessageChan <- message
	return nil
}

func (avsManager *AvsManager) DepositIntoStrategy(operatorAddr common.Address, strategyAddr common.Address, amount *big.Int) error {
	_, tokenAddr, err := avsManager.eigenlayerReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
	if err != nil {
		avsManager.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	contractErc20Mock, err := avsManager.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		avsManager.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := avsManager.avsWriter.TxMgr.GetNoSendTxOpts()
	tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, amount)
	if err != nil {
		avsManager.logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = avsManager.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		avsManager.logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = avsManager.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
	if err != nil {
		avsManager.logger.Errorf("Error depositing into strategy", "err", err)
		return err
	}
	return nil
}

func (avsManager *AvsManager) RegisterOperatorWithEigenlayer(operatorAddr common.Address) error {
	operator := eigenSdkTypes.Operator{
		Address:                 operatorAddr.String(),
		EarningsReceiverAddress: operatorAddr.String(),
	}
	_, err := avsManager.eigenlayerWriter.RegisterAsOperator(context.Background(), operator)
	if err != nil {
		avsManager.logger.Errorf("Error registering operator with eigenlayer")
		return err
	}

	return nil
}

// RegisterOperatorWithAvs Registration specific functions
func (avsManager *AvsManager) RegisterOperatorWithAvs(
	client eth.EthClient,
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
) error {
	// hardcode these things for now
	quorumNumbers := []byte{0}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{123}
	curBlockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		avsManager.logger.Errorf("Unable to get current block number")
		return err
	}

	curBlock, err := client.BlockByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
	if err != nil {
		avsManager.logger.Errorf("Unable to get current block")
		return err
	}

	sigValidForSeconds := int64(1_000_000)
	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time()) + sigValidForSeconds)
	_, err = avsManager.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		context.Background(),
		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
		blsKeyPair, quorumNumbers, socket,
	)

	if err != nil {
		avsManager.logger.Errorf("Unable to register operator with avs registry coordinator")
		return err
	}
	avsManager.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}

func (avsManager *AvsManager) IsOperatorRegistered(options *bind.CallOpts, address common.Address) (bool, error) {
	return avsManager.avsReader.IsOperatorRegistered(options, address)
}

func (avsManager *AvsManager) GetOperatorId(options *bind.CallOpts, address common.Address) ([32]byte, error) {
	return avsManager.avsReader.GetOperatorId(options, address)
}
