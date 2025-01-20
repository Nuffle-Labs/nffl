package chainio

import (
	"errors"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)


func TypedErr(e interface{}) error {
	switch t := e.(type) {
	case error:
		return t
	case string:
		return errors.New(t)
	default:
		return nil
	}
}
type EigenlayerContractBindings struct {
	SlasherAddr           gethcommon.Address
	StrategyManagerAddr   gethcommon.Address
	DelegationManagerAddr gethcommon.Address
	AvsDirectoryAddr      gethcommon.Address
	Slasher               *slasher.ContractISlasher
	DelegationManager     *delegationmanager.ContractDelegationManager
	StrategyManager       *strategymanager.ContractStrategyManager
	AvsDirectory          *avsdirectory.ContractAVSDirectory
}

func WrapError(mainErr interface{}, subErr interface{}) error {
	var main, sub error
	main = TypedErr(mainErr)
	sub = TypedErr(subErr)
	// Some times the wrap will wrap a nil error
	if main == nil && sub == nil {
		return nil
	}

	if main == nil && sub != nil {
		return sub
	}

	if main != nil && sub == nil {
		return main
	}

	return fmt.Errorf("%w: %w", main, sub)
}

// Since v2.0.0, the slasher is not part of the bindings which is causing an issue
// Until we update the bindings, we need to mock the slasher

func NewEigenlayerContractBindings(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethclient eth.Client,
	logger logging.Logger,
) (*EigenlayerContractBindings, error) {
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		return nil, WrapError("Failed to create DelegationManager contract", err)
	}

	slasherAddr := gethcommon.Address{}
	contractSlasher, err := slasher.NewContractISlasher(slasherAddr, ethclient)
	if err != nil {
		return nil, WrapError("Failed to fetch Slasher contract", err)
	}

	strategyManagerAddr, err := contractDelegationManager.StrategyManager(&bind.CallOpts{})
	if err != nil {
		return nil, WrapError("Failed to fetch StrategyManager address", err)
	}
	contractStrategyManager, err := strategymanager.NewContractStrategyManager(strategyManagerAddr, ethclient)
	if err != nil {
		return nil, WrapError("Failed to fetch StrategyManager contract", err)
	}

	avsDirectory, err := avsdirectory.NewContractAVSDirectory(avsDirectoryAddr, ethclient)
	if err != nil {
		return nil, WrapError("Failed to fetch AVSDirectory contract", err)
	}

	return &EigenlayerContractBindings{
		SlasherAddr:           slasherAddr,
		StrategyManagerAddr:   strategyManagerAddr,
		DelegationManagerAddr: delegationManagerAddr,
		AvsDirectoryAddr:      avsDirectoryAddr,
		Slasher:               contractSlasher,
		StrategyManager:       contractStrategyManager,
		DelegationManager:     contractDelegationManager,
		AvsDirectory:          avsDirectory,
	}, nil
}



func BuildElReader(
	registryCoordinatorAddress common.Address,
	operatorStateRetrieverAddress common.Address,
	ethHttpClient eth.Client,
	logger logging.Logger,
) (*elcontracts.ELChainReader, error) {
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

	elContractBindings, err := NewEigenlayerContractBindings(
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

	return elChainReader, nil
}
