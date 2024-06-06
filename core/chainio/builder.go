package chainio

// Rewrite of 'github.com/Layr-Labs/eigensdk-go/chainio/clients/builder.go' to be more flexible
// Supports passing in `eth.Client` rather than building it internally
import (
	"context"
	"crypto/ecdsa"
	"errors"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type Clients struct {
	AvsRegistryChainReader     *avsregistry.AvsRegistryChainReader
	AvsRegistryChainSubscriber *avsregistry.AvsRegistryChainSubscriber
	AvsRegistryChainWriter     *avsregistry.AvsRegistryChainWriter
	ElChainReader              *elcontracts.ELChainReader
	ElChainWriter              *elcontracts.ELChainWriter
}

func BuildAll(
	avsName string,
	registryCoordinatorAddr string,
	operatorStateRetrieverAddr string,
	ethHttpClient eth.Client,
	ethWsClient eth.Client,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	logger logging.Logger,
) (*Clients, error) {
	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	if err != nil {
		return nil, utils.WrapError(errors.New("cannot get chain id"), err)
	}

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError(errors.New("cannot create signer from config"), err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError(errors.New("failed to create transaction sender"), err)
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)

	// creating EL clients: Reader, Writer and Subscriber
	elChainReader, elChainWriter, err := buildElClients(
		registryCoordinatorAddr,
		operatorStateRetrieverAddr,
		ethHttpClient,
		txMgr,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError(errors.New("failed to create EL Reader, Writer and Subscriber"), err)
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryChainWriter, err := buildAvsClients(
		registryCoordinatorAddr,
		operatorStateRetrieverAddr,
		elChainReader,
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError(errors.New("failed to create AVS Registry Reader and Writer"), err)
	}

	return &Clients{
		ElChainReader:              elChainReader,
		ElChainWriter:              elChainWriter,
		AvsRegistryChainReader:     avsRegistryChainReader,
		AvsRegistryChainSubscriber: avsRegistryChainSubscriber,
		AvsRegistryChainWriter:     avsRegistryChainWriter,
	}, nil

}

func buildElClients(
	registryCoordinatorAddr string,
	operatorStateRetrieverAddr string,
	ethHttpClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*elcontracts.ELChainReader, *elcontracts.ELChainWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(registryCoordinatorAddr),
		gethcommon.HexToAddress(operatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, nil, utils.WrapError(errors.New("failed to create AVSRegistryContractBindings"), err)
	}

	delegationManagerAddr, err := avsRegistryContractBindings.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}

	avsDirectoryAddr, err := avsRegistryContractBindings.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch Slasher contract", "err", err)
	}

	elContractBindings, err := chainioutils.NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, nil, utils.WrapError(errors.New("failed to create EigenlayerContractBindings"), err)
	}

	// get the Reader for the EL contracts
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

	return elChainReader, elChainWriter, nil
}

func buildAvsClients(
	registryCoordinatorAddr string,
	operatorStateRetrieverAddr string,
	elReader elcontracts.ELReader,
	ethHttpClient eth.Client,
	ethWsClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*avsregistry.AvsRegistryChainReader, *avsregistry.AvsRegistryChainSubscriber, *avsregistry.AvsRegistryChainWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(registryCoordinatorAddr),
		gethcommon.HexToAddress(operatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, nil, nil, utils.WrapError(errors.New("failed to create AVSRegistryContractBindings"), err)
	}

	avsRegistryChainReader := avsregistry.NewAvsRegistryChainReader(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		avsRegistryContractBindings.BlsApkRegistryAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		logger,
		ethHttpClient,
	)

	avsRegistryChainWriter, err := avsregistry.NewAvsRegistryChainWriter(
		avsRegistryContractBindings.ServiceManagerAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		avsRegistryContractBindings.BlsApkRegistry,
		elReader,
		logger,
		ethHttpClient,
		txMgr,
	)
	if err != nil {
		return nil, nil, nil, utils.WrapError(errors.New("failed to create AvsRegistryChainWriter"), err)
	}

	avsRegistrySubscriber, err := avsregistry.BuildAvsRegistryChainSubscriber(
		avsRegistryContractBindings.BlsApkRegistryAddr,
		ethWsClient, // note that the subscriber needs a ws connection instead of http
		logger,
	)
	if err != nil {
		return nil, nil, nil, utils.WrapError(errors.New("failed to create AvsRegistryChainSubscriber"), err)
	}

	return avsRegistryChainReader, avsRegistrySubscriber, avsRegistryChainWriter, nil
}
