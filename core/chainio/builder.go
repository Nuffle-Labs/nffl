package chainio

import (
	chainioavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewElContractBindings(
	avsRegistryContractBindings *chainioutils.AvsRegistryContractBindings,
	ethHttpClient eth.Client,
	logger logging.Logger,
) (*chainioutils.EigenlayerContractBindings, error) {
	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	avsDirectoryAddr, err := avsRegistryContractBindings.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, err
	}

	return elContractBindings, nil
}

func NewAvsRegistryChainReaderFromContract(
	avsRegistryContractBindings *chainioutils.AvsRegistryContractBindings,
	ethHttpClient eth.Client,
	logger logging.Logger,
) *chainioavsregistry.AvsRegistryChainReader {
	avsRegistryChainReader := chainioavsregistry.NewAvsRegistryChainReader(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		avsRegistryContractBindings.BlsApkRegistryAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		logger,
		ethHttpClient,
	)

	return avsRegistryChainReader
}

func NewELChainReaderFromContract(
	elContractBindings *chainioutils.EigenlayerContractBindings,
	ethHttpClient eth.Client,
	logger logging.Logger,
) *elcontracts.ELChainReader {
	elChainReader := elcontracts.NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		logger,
		ethHttpClient,
	)

	return elChainReader
}

func NewElChainWriterFromBindings(
	elContractBindings *chainioutils.EigenlayerContractBindings,
	elChainReader *elcontracts.ELChainReader,
	ethHttpClient eth.Client,
	txMgr *txmgr.SimpleTxManager,
	logger logging.Logger,
) *elcontracts.ELChainWriter {
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

	return elChainWriter
}

func NewEigenlayerContractBindingsFromContract(
	avsRegistryContractBindings *chainioutils.AvsRegistryContractBindings,
	ethHttpClient eth.Client,
	logger logging.Logger,
) (*chainioutils.EigenlayerContractBindings, error) {
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

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, err
	}

	return elContractBindings, nil
}
