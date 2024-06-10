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
	elChainReader *elcontracts.ELChainReader,
	txMgr *txmgr.SimpleTxManager,
	ethHttpClient eth.Client,
	logger logging.Logger,
) (*elcontracts.ELChainWriter, error) {
	avsRegistryContractBindings, err := utils.NewAVSRegistryContractBindings(registryCoordinatorAddress, operatorStateRetrieverAddress, ethHttpClient, logger)
	if err != nil {
		logger.Error("Cannot create AVSRegistryContractBindings", "err", err)
		return nil, err
	}

	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
		return nil, err
	}

	avsDirectoryAddr, err := avsRegistryContractBindings.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
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