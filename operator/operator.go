package operator

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	taskmanager "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLTaskManager"
	"github.com/Nuffle-Labs/nffl/core"
	"github.com/Nuffle-Labs/nffl/core/chainio"
	"github.com/Nuffle-Labs/nffl/core/safeclient"
	"github.com/Nuffle-Labs/nffl/core/types/messages"
	"github.com/Nuffle-Labs/nffl/operator/attestor"
	optypes "github.com/Nuffle-Labs/nffl/operator/types"
)

const (
	AVS_NAME = "super-fast-finality-layer"
	SEM_VER  = "0.0.1"
)

type Operator struct {
	config    optypes.NodeConfig
	logger    sdklogging.Logger
	ethClient safeclient.SafeClient
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg *prometheus.Registry
	metrics    metrics.Metrics
	listener   OperatorEventListener

	nodeApi          *nodeapi.NodeApi
	blsKeypair       *bls.KeyPair
	operatorId       eigentypes.OperatorId
	operatorAddr     common.Address
	taskResponseWait time.Duration

	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	registryCoordinatorAddr common.Address
	// NEAR DA indexer consumer
	attestor attestor.Attestorer
	// Avs Manager
	avsManager *AvsManager
}

var _ core.Metricable = (*Operator)(nil)

func createLogger(config *optypes.NodeConfig) (sdklogging.Logger, error) {
	var logLevel sdklogging.LogLevel
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
func NewOperatorFromConfig(c optypes.NodeConfig) (*Operator, error) {
	logger, err := createLogger(&c)
	if err != nil {
		return nil, err
	}

	logger.Debug("Creating operator from config", "config", c)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)
	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warn("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}

	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Error("Cannot parse bls private key", "err", err)
		return nil, err
	}

	operatorId := eigentypes.OperatorIdFromG1Pubkey(blsKeyPair.GetPubKeyG1())

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warn("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	reg := prometheus.NewRegistry()
	id := c.OperatorAddress + OperatorSubsytem
	ethHttpClient, err := core.CreateEthClientWithCollector(id, c.EthRpcUrl, c.EnableMetrics, reg, logger)
	if err != nil {
		logger.Error("Cannot create http ethclient", "err", err)
		return nil, err
	}

	ethWsClient, err := core.CreateEthClientWithCollector(id, c.EthWsUrl, c.EnableMetrics, reg, logger)
	if err != nil {
		logger.Error("Cannot create ws ethclient", "err", err)
		return nil, err
	}

	// TODO(edwin): I agree with below.
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethHttpClient.ChainID(context.Background())
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

	txSender, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		logger.Error("Failed to create transaction sender", "err", err)
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(txSender, ethHttpClient, logger, common.HexToAddress(c.OperatorAddress)).WithGasLimitMultiplier(1.5)

	registryCoordinatorAddress := common.HexToAddress(c.AVSRegistryCoordinatorAddress)
	operatorStateRetrieverAddress := common.HexToAddress(c.OperatorStateRetrieverAddress)

	avsReader, err := chainio.BuildAvsReader(registryCoordinatorAddress, operatorStateRetrieverAddress, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create AvsReader", "err", err)
		return nil, err
	}

	elChainReader, err := chainio.BuildElReader(registryCoordinatorAddress, operatorStateRetrieverAddress, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create ElChainReader", "err", err)
		return nil, err
	}

	elChainWriter, err := chainio.BuildElWriter(registryCoordinatorAddress, operatorStateRetrieverAddress, txMgr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create ElChainWriter", "err", err)
		return nil, err
	}

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, operatorId, registryCoordinatorAddress, logger)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	avsManager, err := NewAvsManager(&c, ethHttpClient, ethWsClient, elChainReader, elChainWriter, txMgr, logger)
	if err != nil {
		logger.Error("Cannot create AvsManager", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[eigentypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		elChainReader,
		avsReader,
		AVS_NAME,
		logger,
		common.HexToAddress(c.OperatorAddress),
		quorumNames,
	)
	reg.MustRegister(economicMetricsCollector)

	var optionalMetrics metrics.Metrics
	if c.EnableMetrics {
		optionalMetrics = metrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	} else {
		optionalMetrics = metrics.NewNoopMetrics()
	}

	operator := &Operator{
		config:                     c,
		logger:                     logger,
		ethClient:                  ethHttpClient,
		metricsReg:                 reg,
		metrics:                    optionalMetrics,
		listener:                   &SelectiveOperatorListener{},
		nodeApi:                    nodeApi,
		avsManager:                 avsManager,
		blsKeypair:                 blsKeyPair,
		operatorAddr:               common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,
		registryCoordinatorAddr:    registryCoordinatorAddress,
		operatorId:                 operatorId,
		taskResponseWait:           time.Duration(c.TaskResponseWaitMs) * time.Millisecond,
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

	logger.Info("Operator info",
		"operatorId", operator.operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	attestor, err := attestor.NewAttestor(&c, blsKeyPair, operator.operatorId, reg, logger)
	if err != nil {
		return nil, err
	}
	operator.attestor = attestor

	if c.EnableMetrics {
		if err = operator.EnableMetrics(reg); err != nil {
			return nil, err
		}
	}

	return operator, nil
}

func (o *Operator) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeOperatorMetrics(registry)
	if err != nil {
		return err
	}
	o.listener = listener

	if err = o.attestor.EnableMetrics(registry); err != nil {
		return err
	}

	if err = o.aggregatorRpcClient.EnableMetrics(registry); err != nil {
		return err
	}

	return nil
}

func (o *Operator) Start(ctx context.Context) error {
	o.logger.Info("Starting operator")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}

	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	o.listener.IncInitializationCount()
	o.listener.ObserveLastInitializedTime()

	if err := o.avsManager.Start(ctx, o.operatorAddr); err != nil {
		return err
	}

	// TODO: hmm maybe remove start from attestor?
	if err := o.attestor.Start(ctx); err != nil {
		return err
	}
	// TODO: decide who have a right to sign
	signedRootsC := o.attestor.GetSignedRootC()
	checkpointTaskCreatedChan := o.avsManager.GetCheckpointTaskCreatedChan()
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

		case checkpointTaskCreatedEvent, ok := <-checkpointTaskCreatedChan:
			if !ok {
				o.logger.Info("Closing Operator")
				return o.Close()
			}

			go o.ProcessCheckpointTask(checkpointTaskCreatedEvent)
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

	o.ethClient.Close()

	return nil
}

func (o *Operator) SignTaskResponse(taskResponse *messages.CheckpointTaskResponse) (*messages.SignedCheckpointTaskResponse, error) {
	taskResponseHash, err := taskResponse.Digest()
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}

	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedCheckpointTaskResponse := &messages.SignedCheckpointTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}

	o.logger.Debug("Signed task response", "signedCheckpointTaskResponse", signedCheckpointTaskResponse)
	return signedCheckpointTaskResponse, nil
}

func SignOperatorSetUpdate(message messages.OperatorSetUpdateMessage, blsKeyPair *bls.KeyPair, operatorId eigentypes.OperatorId) (*messages.SignedOperatorSetUpdateMessage, error) {
	messageHash, err := message.Digest()
	if err != nil {
		return nil, err
	}
	signature := blsKeyPair.SignMessage(messageHash)
	signedOperatorSetUpdate := messages.SignedOperatorSetUpdateMessage{
		Message:      message,
		OperatorId:   operatorId,
		BlsSignature: *signature,
	}

	return &signedOperatorSetUpdate, nil
}

func (o *Operator) ProcessCheckpointTask(event *taskmanager.ContractSFFLTaskManagerCheckpointTaskCreated) {
	o.listener.OnTasksReceived()

	if o.taskResponseWait > 0 {
		<-time.After(o.taskResponseWait)
	}

	checkpointMessages, err := o.aggregatorRpcClient.GetAggregatedCheckpointMessages(
		event.Task.FromTimestamp,
		event.Task.ToTimestamp,
	)
	if err != nil {
		o.logger.Error("Failed to get checkpoint messages", "err", err)
		return
	}

	checkpointTaskResponse, err := messages.NewCheckpointTaskResponseFromMessages(
		event.TaskIndex,
		checkpointMessages,
	)
	if err != nil {
		o.logger.Error("Failed to get create checkpoint response", "err", err)
		return
	}

	signedCheckpointTaskResponse, err := o.SignTaskResponse(&checkpointTaskResponse)
	if err != nil {
		o.logger.Error("Failed to sign checkpoint task response", "checkpointTaskResponse", checkpointTaskResponse)
		return
	}

	go o.aggregatorRpcClient.SendSignedCheckpointTaskResponseToAggregator(signedCheckpointTaskResponse)
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
		o.logger.Info("Registered operator with eigenlayer")
	}

	if mockTokenStrategyAddr.Cmp(common.Address{}) != 0 {
		// TODO(samlaf): shouldn't hardcode number here
		amount := big.NewInt(1000)
		err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
		if err != nil {
			o.logger.Fatal("Error depositing into strategy", "err", err)
		}
		o.logger.Info("Deposited into strategy", "amount", amount, "strategy", mockTokenStrategyAddr)
	}

	isOperatorRegistered, err := o.avsManager.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		o.logger.Fatal("Error checking if operator is registered", "err", err)
	}

	if !isOperatorRegistered {
		err = o.avsManager.RegisterOperatorWithAvs(o.ethClient, operatorEcdsaPrivateKey, o.blsKeypair)
		if err != nil {
			o.logger.Fatal("Error registering operator with avs", "err", err)
		}
		o.logger.Info("Registered operator with avs")
	} else {
		o.logger.Info("Operator already registered with avs")
	}
}

func (o *Operator) BlsPubkeyG1() *bls.G1Point {
	return o.blsKeypair.GetPubKeyG1()
}
