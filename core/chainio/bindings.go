package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	erc20mock "github.com/NethermindEth/near-sffl/contracts/bindings/ERC20Mock"
	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	regcoord "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryCoordinator"
	csservicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

type AvsManagersBindings struct {
	RegistryCoordinator       *regcoord.ContractSFFLRegistryCoordinator
	OperatorSetUpdateRegistry *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistry
	TaskManager               *taskmanager.ContractSFFLTaskManager
	ServiceManager            *csservicemanager.ContractSFFLServiceManager
	ethClient                 eth.EthClient
	logger                    logging.Logger
}

func NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethclient eth.EthClient, logger logging.Logger) (*AvsManagersBindings, error) {
	contractRegistryCoordinator, err := regcoord.NewContractSFFLRegistryCoordinator(registryCoordinatorAddr, ethclient)
	if err != nil {
		return nil, err
	}

	serviceManagerAddr, err := contractRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	contractServiceManager, err := csservicemanager.NewContractSFFLServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}

	taskManagerAddr, err := contractServiceManager.TaskManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}
	contractTaskManager, err := taskmanager.NewContractSFFLTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch SFFLTaskManager contract", "err", err)
		return nil, err
	}

	operatorSetUpdateRegistryAddr, err := contractRegistryCoordinator.OperatorSetUpdateRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	contractOperatorSetUpdateRegistry, err := opsetupdatereg.NewContractSFFLOperatorSetUpdateRegistry(operatorSetUpdateRegistryAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch OperatorSetUpdateRegistry contract", "err", err)
		return nil, err
	}

	return &AvsManagersBindings{
		RegistryCoordinator:       contractRegistryCoordinator,
		OperatorSetUpdateRegistry: contractOperatorSetUpdateRegistry,
		ServiceManager:            contractServiceManager,
		TaskManager:               contractTaskManager,
		ethClient:                 ethclient,
		logger:                    logger,
	}, nil
}

func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractERC20Mock, error) {
	contractErc20Mock, err := erc20mock.NewContractERC20Mock(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}

	return contractErc20Mock, nil
}
