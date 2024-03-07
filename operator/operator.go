package operator

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/metrics"
	"github.com/NethermindEth/near-sffl/operator/attestor"
	"github.com/NethermindEth/near-sffl/types"
)

const AVS_NAME = "super-fast-finality-layer"
const SEM_VER = "0.0.1"

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.EthClient
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg   *prometheus.Registry
	metrics      metrics.Metrics
	nodeApi      *nodeapi.NodeApi
	blsKeypair   *bls.KeyPair
	operatorId   bls.OperatorId
	operatorAddr common.Address

	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	sfflServiceManagerAddr common.Address
	// NEAR DA indexer consumer
	attestor attestor.Attestorer
	// Avs Manager
	avsManager *AvsManager
}

func createEthClients(config *types.NodeConfig, registry *prometheus.Registry, logger sdklogging.Logger) (eth.EthClient, eth.EthClient, error) {
	if config.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, registry)
		ethRpcClient, err := eth.NewInstrumentedClient(config.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, nil, err
		}
		ethWsClient, err := eth.NewInstrumentedClient(config.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, nil, err
		}

		return ethRpcClient, ethWsClient, nil
	}

	ethRpcClient, err := eth.NewClient(config.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, nil, err
	}
	ethWsClient, err := eth.NewClient(config.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return nil, nil, err
	}

	return ethRpcClient, ethWsClient, nil
}

func createLogger(config *types.NodeConfig) (sdklogging.Logger, error) {
	var logLevel logging.LogLevel
	if config.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}

	return sdklogging.NewZapLogger(logLevel)
}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	logger, err := createLogger(&c)
	if err != nil {
		return nil, err
	}

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)
	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}

	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	reg := prometheus.NewRegistry()
	ethRpcClient, ethWsClient, err := createEthClients(&c, reg, logger)
	if err != nil {
		return nil, err
	}

	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, common.HexToAddress(c.OperatorAddress))
	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 c.EthRpcUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	}
	sdkClients, err := clients.BuildAll(chainioConfig, common.HexToAddress(c.OperatorAddress), signerV2, logger)
	if err != nil {
		panic(err)
	}

	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, sdkClients.Metrics, reg)
	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	avsManager, err := NewAvsManager(&c, ethRpcClient, ethWsClient, sdkClients, txMgr, logger)
	if err != nil {
		logger.Error("Cannot create AvsManager", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	operator := &Operator{
		config:                     c,
		logger:                     logger,
		ethClient:                  ethRpcClient,
		metricsReg:                 reg,
		metrics:                    avsAndEigenMetrics,
		nodeApi:                    nodeApi,
		avsManager:                 avsManager,
		blsKeypair:                 blsKeyPair,
		operatorAddr:               common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,
		sfflServiceManagerAddr:     common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                 [32]byte{0}, // this is set below
	}

	if c.RegisterOperatorOnStartup {
		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			c.EcdsaPrivateKeyStorePath,
			ecdsaKeyPassword,
		)
		if err != nil {
			return nil, err
		}

		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey, common.HexToAddress(c.TokenStrategyAddr))
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := avsManager.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}

	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	attestor, err := attestor.NewAttestor(&c, blsKeyPair, operatorId, logger)
	if err != nil {
		return nil, err
	}
	operator.attestor = attestor

	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	if err := o.avsManager.Start(ctx, o.operatorAddr); err != nil {
		return err
	}

	// TODO: hmm maybe remove start from attestor?
	if err := o.attestor.Start(ctx); err != nil {
		return err
	}

	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}

	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO: decide who have a right to sign
	signedRootsC := o.attestor.GetSignedRootC()
	checkpointTaskResponseChan := o.avsManager.GetCheckpointTaskResponseChan()
	operatorSetUpdateChan := o.avsManager.GetOperatorSetUpdateChan()

	for {
		select {
		case <-ctx.Done():
			return o.Close()

		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.logger.Fatal("Error in metrics server", "err", err)
			continue

		case signedStateRootUpdateMessage := <-signedRootsC:
			go o.aggregatorRpcClient.SendSignedStateRootUpdateToAggregator(&signedStateRootUpdateMessage)
			continue

		case checkpointTaskResponse, ok := <-checkpointTaskResponseChan:
			if !ok {
				o.logger.Info("Closing Operator")
				return o.Close()
			}

			o.metrics.IncNumTasksReceived()
			signedCheckpointTaskResponse, err := o.SignTaskResponse(&checkpointTaskResponse)
			if err != nil {
				o.logger.Error("Failed to sign checkpoint task response", "checkpointTaskResponse", checkpointTaskResponse)
				continue
			}

			go o.aggregatorRpcClient.SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse)
			continue

		case operatorSetUpdate, ok := <-operatorSetUpdateChan:
			if !ok {
				o.logger.Info("Closing Operator")
				return o.Close()
			}

			signedOperatorSetUpdate, err := SignOperatorSetUpdate(operatorSetUpdate, o.blsKeypair, o.operatorId)
			if err != nil {
				o.logger.Error("Failed to sign operator set update", "signedOperatorSetUpdate", signedOperatorSetUpdate)
				continue
			}

			go o.aggregatorRpcClient.SendSignedOperatorSetUpdateToAggregator(signedOperatorSetUpdate)
			continue
		}
	}
}

func (o *Operator) Close() error {
	if err := o.attestor.Close(); err != nil {
		return err
	}

	return nil
}

func (o *Operator) SignTaskResponse(taskResponse *taskmanager.CheckpointTaskResponse) (*coretypes.SignedCheckpointTaskResponse, error) {
	taskResponseHash, err := core.GetCheckpointTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}

	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedCheckpointTaskResponse := &coretypes.SignedCheckpointTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}

	o.logger.Debug("Signed task response", "signedCheckpointTaskResponse", signedCheckpointTaskResponse)
	return signedCheckpointTaskResponse, nil
}

func SignOperatorSetUpdate(message registryrollup.OperatorSetUpdateMessage, blsKeyPair *bls.KeyPair, operatorId bls.OperatorId) (*coretypes.SignedOperatorSetUpdateMessage, error) {
	messageHash, err := core.GetOperatorSetUpdateMessageDigest(&message)
	if err != nil {
		return nil, err
	}
	signature := blsKeyPair.SignMessage(messageHash)
	signedOperatorSetUpdate := coretypes.SignedOperatorSetUpdateMessage{
		Message:      message,
		OperatorId:   operatorId,
		BlsSignature: *signature,
	}

	return &signedOperatorSetUpdate, nil
}

func (o *Operator) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	return o.avsManager.RegisterOperatorWithAvs(o.ethClient, operatorEcdsaKeyPair, o.blsKeypair)
}

func (o *Operator) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
	return o.avsManager.DepositIntoStrategy(o.operatorAddr, strategyAddr, amount)
}

func (o *Operator) RegisterOperatorWithEigenlayer() error {
	return o.avsManager.RegisterOperatorWithEigenlayer(o.operatorAddr)
}

type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func (o *Operator) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	operatorId, err := o.avsManager.GetOperatorId(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		return err
	}

	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := o.operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(o.operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(operatorStatusJson))
	return nil
}

func (o *Operator) registerOperatorOnStartup(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	mockTokenStrategyAddr common.Address,
) {
	err := o.RegisterOperatorWithEigenlayer()
	if err != nil {
		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
	} else {
		o.logger.Infof("Registered operator with eigenlayer")
	}

	// TODO(samlaf): shouldn't hardcode number here
	amount := big.NewInt(1000)
	err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
	if err != nil {
		o.logger.Fatal("Error depositing into strategy", "err", err)
	}
	o.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

	err = o.avsManager.RegisterOperatorWithAvs(o.ethClient, operatorEcdsaPrivateKey, o.blsKeypair)
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
	o.logger.Infof("Registered operator with avs")
}

func (o *Operator) BlsPubkeyG1() *bls.G1Point {
	return o.blsKeypair.GetPubKeyG1()
}
