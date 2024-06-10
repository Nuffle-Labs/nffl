package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func BuildElWriter(
	registryCoordinatorAddress common.Address,
	operatorStateRetrieverAddress common.Address,
	txMgr *txmgr.SimpleTxManager,
	ethHttpClient eth.Client,
	logger logging.Logger,
) (*elcontracts.ELChainWriter, error) {
	avsRegistryContractBindings, err := utils.NewAVSRegistryContractBindings(registryCoordinatorAddress, operatorStateRetrieverAddress, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}

	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	avsDirectoryAddr, err := avsRegistryContractBindings.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	elContractBindings, err := utils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, err
	}

	elChainReader := elcontracts.NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		logger,
		ethHttpClient,
	)

	elChainWriter := elcontracts.NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethHttpClient,
		logger,
		nil,
		txMgr,
	)

	return elChainWriter, nil
}
