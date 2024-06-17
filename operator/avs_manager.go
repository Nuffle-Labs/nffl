package operator

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	eigenutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

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
		client eth.Client,
		operatorEcdsaKeyPair *ecdsa.PrivateKey,
		blsKeyPair *bls.KeyPair,
	) error
	DeregisterOperator(blsKeyPair *bls.KeyPair) error
	GetOperatorId(options *bind.CallOpts, address common.Address) ([32]byte, error)
	GetCheckpointTaskCreatedChan() <-chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated
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

	operatorSetUpdateMessageChan chan messages.OperatorSetUpdateMessage

	logger sdklogging.Logger
}

var _ AvsManagerer = (*AvsManager)(nil)

func NewAvsManager(config *optypes.NodeConfig, ethRpcClient eth.Client, ethWsClient eth.Client, elChainReader *elcontracts.ELChainReader, elChainWriter *elcontracts.ELChainWriter, txManager txmgr.TxManager, logger sdklogging.Logger) (*AvsManager, error) {
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
		avsReader:                    avsReader,
		avsWriter:                    avsWriter,
		avsSubscriber:                avsSubscriber,
		eigenlayerReader:             elChainReader,
		eigenlayerWriter:             elChainWriter,
		checkpointTaskCreatedChan:    make(chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated),
		operatorSetUpdateChan:        make(chan *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock),
		operatorSetUpdateMessageChan: make(chan messages.OperatorSetUpdateMessage),
		logger:                       logger,
	}, nil
}

func (avsManager *AvsManager) GetCheckpointTaskCreatedChan() <-chan *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated {
	return avsManager.checkpointTaskCreatedChan
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
				avsManager.logger.Error("New tasks subscription error", "err", err)
				newTasksSub.Unsubscribe()
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case err := <-operatorSetUpdateSub.Err():
				avsManager.logger.Error("Operator set update subscription error", "err", err)
				operatorSetUpdateSub.Unsubscribe()
				return

			case operatorSetUpdate := <-avsManager.operatorSetUpdateChan:
				go avsManager.handleOperatorSetUpdate(ctx, operatorSetUpdate)
				continue
			}
		}
	}()

	return nil
}

func (avsManager *AvsManager) handleOperatorSetUpdate(ctx context.Context, data *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock) error {
	operatorSetDelta, err := avsManager.avsReader.GetOperatorSetUpdateDelta(ctx, data.Id)
	if err != nil {
		avsManager.logger.Error("Couldn't get Operator set update delta", "err", err, "block", data.Id)
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
	if err != nil {
		avsManager.logger.Error("Error getting tx options")
		return err
	}
	tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, amount)
	if err != nil {
		avsManager.logger.Error("Error assembling Mint tx")
		return err
	}
	_, err = avsManager.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		avsManager.logger.Error("Error submitting Mint tx")
		return err
	}

	_, err = avsManager.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
	if err != nil {
		avsManager.logger.Error("Error depositing into strategy", "err", err)
		return err
	}
	return nil
}

func (avsManager *AvsManager) RegisterOperatorWithEigenlayer(operatorAddr common.Address) error {
	operator := eigentypes.Operator{
		Address:                 operatorAddr.String(),
		EarningsReceiverAddress: operatorAddr.String(),
	}
	_, err := avsManager.eigenlayerWriter.RegisterAsOperator(context.Background(), operator)
	if err != nil {
		avsManager.logger.Error("Error registering operator with eigenlayer")
		return err
	}

	return nil
}

// RegisterOperatorWithAvs Registration specific functions
func (avsManager *AvsManager) RegisterOperatorWithAvs(
	client eth.Client,
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
) error {
	// hardcode these things for now
	quorumNumbers := eigentypes.QuorumNums{0}
	socket := "Not Needed"
	curBlockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		avsManager.logger.Error("Unable to get current block number")
		return err
	}

	curBlock, err := client.BlockByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
	if err != nil {
		avsManager.logger.Error("Unable to get current block")
		return err
	}

	operatorId := eigentypes.OperatorIdFromG1Pubkey(blsKeyPair.GetPubKeyG1())

	sigValidForSeconds := int64(1_000_000)
	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time()) + sigValidForSeconds)
	operatorToAvsRegistrationSigSalt := [32]byte{}
	copy(operatorToAvsRegistrationSigSalt[:], crypto.Keccak256([]byte("sffl"), operatorId[:], quorumNumbers.UnderlyingType(), []byte(time.Now().String())))
	_, err = avsManager.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		context.Background(),
		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
		blsKeyPair, quorumNumbers, socket,
	)

	if err != nil {
		avsManager.logger.Error("Unable to register operator with avs registry coordinator")
		return err
	}
	avsManager.logger.Info("Registered operator with avs registry coordinator.")

	return nil
}

func (avsManager *AvsManager) DeregisterOperator(blsKeyPair *bls.KeyPair) error {
	// TODO: 'QuorumNums' is hardcoded for now
	quorumNumbers := eigentypes.QuorumNums{0}
	pubKey := eigenutils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())

	_, err := avsManager.avsWriter.DeregisterOperator(context.Background(), quorumNumbers, pubKey)
	if err != nil {
		avsManager.logger.Error("Unable to deregister operator with avs registry coordinator", "err", err)
		return err
	}
	avsManager.logger.Info("Deregistered operator with avs registry coordinator")

	return nil
}

func (avsManager *AvsManager) IsOperatorRegistered(options *bind.CallOpts, address common.Address) (bool, error) {
	return avsManager.avsReader.IsOperatorRegistered(options, address)
}

func (avsManager *AvsManager) GetOperatorId(options *bind.CallOpts, address common.Address) ([32]byte, error) {
	return avsManager.avsReader.GetOperatorId(options, address)
}
