package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	erc20mock "github.com/NethermindEth/near-sffl/contracts/bindings/ERC20Mock"
	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	sfflregcoord "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryCoordinator"
	csservicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
)

type AvsManagersBindings struct {
	RegistryCoordinator       *regcoord.ContractRegistryCoordinator
	SFFLRegistryCoordinator   *sfflregcoord.ContractSFFLRegistryCoordinator
	OperatorSetUpdateRegistry *opsetupdatereg.ContractSFFLOperatorSetUpdateRegistry
	TaskManager               *taskmanager.ContractSFFLTaskManager
	ServiceManager            *csservicemanager.ContractSFFLServiceManager
	BlsApkRegistry            blsapkreg.ContractBLSApkRegistryFilters
	ethClient                 eth.Client
	logger                    logging.Logger
}

func NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr common.Address, ethclient eth.Client, logger logging.Logger) (*AvsManagersBindings, error) {
	contractSfflRegistryCoordinator, err := sfflregcoord.NewContractSFFLRegistryCoordinator(registryCoordinatorAddr, ethclient)
	if err != nil {
		return nil, err
	}

	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethclient)
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

	operatorSetUpdateRegistryAddr, err := contractSfflRegistryCoordinator.OperatorSetUpdateRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	contractOperatorSetUpdateRegistry, err := opsetupdatereg.NewContractSFFLOperatorSetUpdateRegistry(operatorSetUpdateRegistryAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch OperatorSetUpdateRegistry contract", "err", err)
		return nil, err
	}

	blsApkRegistryAddr, err := contractRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	blsApkRegistry, err := blsapkreg.NewContractBLSApkRegistry(blsApkRegistryAddr, ethclient)
	if err != nil {
		return nil, err
	}

	return &AvsManagersBindings{
		RegistryCoordinator:       contractRegistryCoordinator,
		SFFLRegistryCoordinator:   contractSfflRegistryCoordinator,
		OperatorSetUpdateRegistry: contractOperatorSetUpdateRegistry,
		ServiceManager:            contractServiceManager,
		TaskManager:               contractTaskManager,
		BlsApkRegistry:            blsApkRegistry,
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
