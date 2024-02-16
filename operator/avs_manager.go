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
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	regcoord "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryCoordinator"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/types"
)

type AvsManager struct {
	ethClient        eth.EthClient
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader elcontracts.ELReader
	eigenlayerWriter elcontracts.ELWriter
	operatorAddr     common.Address

	// receive new tasks in this chan (typically from listening to onchain event)
	checkpointTaskCreatedChan chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated
	// receive operator set updates in this chan
	operatorSetUpdateChan chan *regcoord.ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock

	// Prepare message for operator to sign
	checkpointTaskResponseCreatedChan chan taskmanager.CheckpointTaskResponse
	operatorSetUpdateMessageChan      chan registryrollup.OperatorSetUpdateMessage

	logger sdklogging.Logger
}

func NewAvsManager(config *types.NodeConfig, ethRpcClient eth.EthClient, ethWsClient eth.EthClient, signerV2 signerv2.SignerFn, registry *prometheus.Registry, logger sdklogging.Logger) (*AvsManager, error) {
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

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 config.EthRpcUrl,
		EthWsUrl:                   config.EthWsUrl,
		RegistryCoordinatorAddr:    config.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: config.OperatorStateRetrieverAddress,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   config.EigenMetricsIpPortAddress,
	}
	sdkClients, err := clients.BuildAll(chainioConfig, common.HexToAddress(config.OperatorAddress), signerV2, logger)
	if err != nil {
		panic(err)
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(config.OperatorAddress), quorumNames)
	registry.MustRegister(economicMetricsCollector)

	return &AvsManager{
		ethClient:                         ethRpcClient,
		avsReader:                         avsReader,
		avsWriter:                         avsWriter,
		avsSubscriber:                     avsSubscriber,
		operatorAddr:                      common.HexToAddress(config.OperatorAddress),
		eigenlayerReader:                  sdkClients.ElChainReader,
		eigenlayerWriter:                  sdkClients.ElChainWriter,
		checkpointTaskCreatedChan:         make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:             make(chan *regcoord.ContractSFFLRegistryCoordinatorOperatorSetUpdatedAtBlock),
		checkpointTaskResponseCreatedChan: make(chan taskmanager.CheckpointTaskResponse),
		operatorSetUpdateMessageChan:      make(chan registryrollup.OperatorSetUpdateMessage),
		logger:                            logger,
	}, nil
}

func (avsManager *AvsManager) GetCheckpointTaskResponseChan() <-chan taskmanager.CheckpointTaskResponse {
	return avsManager.checkpointTaskResponseCreatedChan
}

func (avsManager *AvsManager) GetOperatorSetUpdateChan() <-chan registryrollup.OperatorSetUpdateMessage {
	return avsManager.operatorSetUpdateMessageChan
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

	go func(ctx context.Context) {
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
			case <-ctx.Done():
				return

			case err := <-newTasksSub.Err():
				avsManager.logger.Error("Error in websocket subscription", "err", err)
				// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
				newTasksSub.Unsubscribe()
				// TODO(samlaf): wrap this call with increase in avs-node-spec metric
				newTasksSub = avsManager.avsSubscriber.SubscribeToNewTasks(avsManager.checkpointTaskCreatedChan)
				continue

			case checkpointTaskCreatedLog := <-avsManager.checkpointTaskCreatedChan:
				//avsManager.metrics.IncNumTasksReceived()
				taskResponse := avsManager.ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog)
				avsManager.checkpointTaskResponseCreatedChan <- taskResponse

			case err := <-operatorSetUpdateSub.Err():
				avsManager.logger.Error("Error in websocket subscription", "err", err)
				operatorSetUpdateSub.Unsubscribe()
				operatorSetUpdateSub = avsManager.avsSubscriber.SubscribeToOperatorSetUpdates(avsManager.operatorSetUpdateChan)
				continue

			case operatorSetUpdate := <-avsManager.operatorSetUpdateChan:
				go avsManager.handleOperatorSetUpdate(ctx, operatorSetUpdate)
				continue
			}
		}
	}(ctx)

	return nil
}

// Takes a CheckpointTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (avsManager *AvsManager) ProcessCheckpointTaskCreatedLog(checkpointTaskCreatedLog *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) taskmanager.CheckpointTaskResponse {
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

	taskResponse := taskmanager.CheckpointTaskResponse{
		ReferenceTaskIndex:     checkpointTaskCreatedLog.TaskIndex,
		StateRootUpdatesRoot:   [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		OperatorSetUpdatesRoot: [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	return taskResponse
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

	avsManager.operatorSetUpdateMessageChan <- message
	return nil
}

func (avsManager *AvsManager) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
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
	tx, err := contractErc20Mock.Mint(txOpts, avsManager.operatorAddr, amount)
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

func (avsManager *AvsManager) RegisterOperatorWithEigenlayer() error {
	// TODO: could pass as param
	op := eigenSdkTypes.Operator{
		Address:                 avsManager.operatorAddr.String(),
		EarningsReceiverAddress: avsManager.operatorAddr.String(),
	}
	_, err := avsManager.eigenlayerWriter.RegisterAsOperator(context.Background(), op)
	if err != nil {
		avsManager.logger.Errorf("Error registering operator with eigenlayer")
		return err
	}
	return nil
}

func (avsManager *AvsManager) registerOperatorOnStartup(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	mockTokenStrategyAddr common.Address,
	blsKeyPair *bls.KeyPair,
) {
	err := avsManager.RegisterOperatorWithEigenlayer()
	if err != nil {
		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
		avsManager.logger.Error("Error registering operator with eigenlayer", "err", err)
	} else {
		avsManager.logger.Infof("Registered operator with eigenlayer")
	}

	// TODO(samlaf): shouldn't hardcode number here
	amount := big.NewInt(1000)
	err = avsManager.DepositIntoStrategy(mockTokenStrategyAddr, amount)
	if err != nil {
		avsManager.logger.Fatal("Error depositing into strategy", "err", err)
	}
	avsManager.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

	err = avsManager.RegisterOperatorWithAvs(operatorEcdsaPrivateKey, blsKeyPair)
	if err != nil {
		avsManager.logger.Fatal("Error registering operator with avs", "err", err)
	}
	avsManager.logger.Infof("Registered operator with avs")
}

// Registration specific functions
func (avsManager *AvsManager) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
) error {
	// hardcode these things for now
	quorumNumbers := []byte{0}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{123}
	curBlockNum, err := avsManager.ethClient.BlockNumber(context.Background())
	if err != nil {
		avsManager.logger.Errorf("Unable to get current block number")
		return err
	}

	curBlock, err := avsManager.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
	if err != nil {
		avsManager.logger.Errorf("Unable to get current block")
		return err
	}

	sigValidForSeconds := int64(1_000_000)
	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time()) + sigValidForSeconds)
	_, err = avsManager.avsWriter.RegisterOperatorWithAVSRegistryCoordinator(
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
