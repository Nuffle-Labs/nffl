package integration_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/NethermindEth/near-sffl/aggregator"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	transparentproxy "github.com/NethermindEth/near-sffl/contracts/bindings/TransparentUpgradeableProxy"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/operator"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
	"github.com/NethermindEth/near-sffl/tests/integration/utils"
)

const (
	TEST_DATA_DIR  = "../../test_data"
	BLS_KEYS_DIR   = "../keys/bls"
	ECDSA_KEYS_DIR = "../keys/ecdsa"
)

func TestIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Second)
	setup := setupTestEnv(t, ctx)
	t.Cleanup(func() {
		cancel()

		setup.cleanup()
	})

	time.Sleep(30 * time.Second)

	taskHash, err := setup.avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskHashes(&bind.CallOpts{}, 0)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	taskResponseHash, err := setup.avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskResponses(&bind.CallOpts{}, 0)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
	}

	stateRootHeight, err := setup.rollupAnvils[1].HttpClient.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Cannot get current block height: %s", err.Error())
	}

	stateRootUpdate, err := getStateRootUpdateAggregation(setup.aggregatorRestUrl, uint32(setup.rollupAnvils[0].ChainID.Uint64()), stateRootHeight-1)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}
	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	newOperatorConfig, _, _ := genOperatorConfig(t, ctx, "4", setup.mainnetAnvil, setup.rollupAnvils, setup.rabbitMq)
	newOperator := startOperator(t, ctx, newOperatorConfig)

	time.Sleep(50 * time.Second)

	// Check if operator set was updated on rollups
	for _, registryRollup := range setup.registryRollups {
		nextOperatorSetUpdateId, err := registryRollup.NextOperatorUpdateId(&bind.CallOpts{})
		if err != nil {
			t.Fatalf("Error getting next operator set update ID: %s", err.Error())
		}

		if nextOperatorSetUpdateId != 2 {
			t.Fatalf("Wrong next operator set update ID: expected %d, got %d", 2, nextOperatorSetUpdateId)
		}
	}

	stateRootHeight = uint64(16)
	stateRootUpdate, err = getStateRootUpdateAggregation(setup.aggregatorRestUrl, uint32(setup.rollupAnvils[0].ChainID.Uint64()), stateRootHeight)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}
	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	operatorSetUpdateCount, err := setup.avsReader.AvsServiceBindings.OperatorSetUpdateRegistry.GetOperatorSetUpdateCount(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("Error getting operator set update count: %s", err.Error())
	}
	if operatorSetUpdateCount != 2 {
		t.Fatalf("Wrong operator set update count")
	}

	stateRootHeight, err = setup.rollupAnvils[1].HttpClient.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Cannot get current block height: %s", err.Error())
	}

	stateRootUpdate, err = getStateRootUpdateAggregation(setup.aggregatorRestUrl, uint32(setup.rollupAnvils[0].ChainID.Uint64()), stateRootHeight-1)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}

	// Check if operator sets are same on rollups
	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	operatorSetUpdate, err := getOperatorSetUpdateAggregation(setup.aggregatorRestUrl, operatorSetUpdateCount-1)
	if err != nil {
		t.Fatalf("Error getting operator set update: %s", err.Error())
	}

	expectedUpdatedOperators := []types.RollupOperator{
		{
			Pubkey: newOperator.BlsPubkeyG1(),
			Weight: big.NewInt(1000),
		},
	}
	assert.Equal(t, expectedUpdatedOperators, operatorSetUpdate.Message.Operators)

	t.Log("Done")
	<-ctx.Done()
}

type testEnv struct {
	mainnetAnvil        *utils.AnvilInstance
	rollupAnvils        []*utils.AnvilInstance
	rabbitMq            *rabbitmq.RabbitMQContainer
	indexerContainer    testcontainers.Container
	operator            *operator.Operator
	aggregator          *aggregator.Aggregator
	aggregatorRestUrl   string
	avsReader           *chainio.AvsReader
	registryRollups     []*registryrollup.ContractSFFLRegistryRollup
	registryRollupAuths []*bind.TransactOpts
	cleanup             func()
}

func setupTestEnv(t *testing.T, ctx context.Context) *testEnv {
	containersCtx, cancelContainersCtx := context.WithCancel(context.Background())

	networkName := "near-sffl"
	net, err := testcontainers.GenericNetwork(containersCtx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Driver:         "bridge",
			Name:           networkName,
			CheckDuplicate: true,
			Attachable:     true,
		},
	})
	if err != nil {
		t.Fatalf("Cannot create network: %s", err.Error())
	}

	indexerContainerName := "indexer"
	mainnetAnvilContainerName := "mainnet-anvil"
	rollup0AnvilContainerName := "rollup0-anvil"
	rollup1AnvilContainerName := "rollup1-anvil"
	rmqContainerName := "rmq"

	mainnetAnvil := startAnvilTestContainer(t, containersCtx, mainnetAnvilContainerName, "8545", "1", true, networkName)
	rollupAnvils := []*utils.AnvilInstance{
		startAnvilTestContainer(t, containersCtx, rollup0AnvilContainerName, "8546", "2", false, networkName),
		startAnvilTestContainer(t, containersCtx, rollup1AnvilContainerName, "8547", "3", false, networkName),
	}
	rabbitMq := startRabbitMqContainer(t, containersCtx, rmqContainerName, networkName)
	indexerContainer, relayers := startIndexer(t, containersCtx, indexerContainerName, rollupAnvils, rabbitMq, networkName)

	sfflDeploymentRaw := readSfflDeploymentRaw()

	configRaw := buildConfigRaw(mainnetAnvil, rollupAnvils)
	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

	nodeConfig, _, _ := genOperatorConfig(t, ctx, "3", mainnetAnvil, rollupAnvils, rabbitMq)

	addresses, registryRollups, registryRollupAuths, _ := deployRegistryRollups(t, rollupAnvils)
	operator := startOperator(t, ctx, nodeConfig)

	config := buildConfig(t, sfflDeploymentRaw, addresses, rollupAnvils, configRaw)
	aggregator := startAggregator(t, ctx, config, logger)

	avsReader, err := chainio.BuildAvsReader(common.HexToAddress(sfflDeploymentRaw.Addresses.RegistryCoordinatorAddr), common.HexToAddress(sfflDeploymentRaw.Addresses.OperatorStateRetrieverAddr), mainnetAnvil.HttpClient, logger)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	cleanup := func() {
		if err := os.RemoveAll(TEST_DATA_DIR); err != nil {
			t.Fatalf("Error cleaning test data dir: %s", err.Error())
		}

		time.Sleep(5 * time.Second)

		if err := mainnetAnvil.Container.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		for _, rollupAnvil := range rollupAnvils {
			if err := rollupAnvil.Container.Terminate(containersCtx); err != nil {
				t.Fatalf("Error terminating container: %s", err.Error())
			}
		}

		if err := rabbitMq.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		if err := indexerContainer.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		for _, relayer := range relayers {
			if err := relayer.Terminate(containersCtx); err != nil {
				t.Fatalf("Error terminating container: %s", err.Error())
			}
		}

		if err := net.Remove(containersCtx); err != nil {
			t.Fatalf("Error removing network: %s", err.Error())
		}

		cancelContainersCtx()
	}

	return &testEnv{
		mainnetAnvil:        mainnetAnvil,
		rollupAnvils:        rollupAnvils,
		rabbitMq:            rabbitMq,
		indexerContainer:    indexerContainer,
		operator:            operator,
		aggregator:          aggregator,
		aggregatorRestUrl:   "http://" + config.AggregatorRestServerIpPortAddr,
		avsReader:           avsReader,
		registryRollups:     registryRollups,
		registryRollupAuths: registryRollupAuths,
		cleanup:             cleanup,
	}
}

func startOperator(t *testing.T, ctx context.Context, nodeConfig optypes.NodeConfig) *operator.Operator {
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		t.Fatalf("Failed to create operator: %s", err.Error())
	}

	go operator.Start(ctx)

	t.Log("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	return operator
}

func startAggregator(t *testing.T, ctx context.Context, config *config.Config, logger sdklogging.Logger) *aggregator.Aggregator {
	t.Log("starting aggregator for integration tests")

	agg, err := aggregator.NewAggregator(ctx, config, logger)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}

	go agg.Start(ctx)

	t.Log("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	return agg
}

func startRabbitMqContainer(t *testing.T, ctx context.Context, name, networkName string) *rabbitmq.RabbitMQContainer {
	rabbitMqC, err := rabbitmq.RunContainer(
		ctx,
		testcontainers.WithImage("rabbitmq:latest"),
		func() testcontainers.CustomizeRequestOption {
			return func(req *testcontainers.GenericContainerRequest) {
				req.Name = name
				req.Networks = []string{networkName}
			}
		}(),
	)
	if err != nil {
		t.Fatalf("Error starting RMQ container: %s", err.Error())
	}

	return rabbitMqC
}

func readSfflDeploymentRaw() config.SFFLDeploymentRaw {
	var sfflDeploymentRaw config.SFFLDeploymentRaw
	sfflDeploymentFilePath := "../../contracts/evm/script/output/31337/sffl_avs_deployment_output.json"
	sdkutils.ReadJsonConfig(sfflDeploymentFilePath, &sfflDeploymentRaw)

	return sfflDeploymentRaw
}

func genOperatorConfig(t *testing.T, ctx context.Context, keyId string, mainnetAnvil *utils.AnvilInstance, rollupAnvils []*utils.AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) (optypes.NodeConfig, *bls.KeyPair, *ecdsa.PrivateKey) {
	nodeConfig := optypes.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err := sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	log.Println("starting operator for integration tests")

	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")

	nodeConfig.BlsPrivateKeyStorePath, err = filepath.Abs(filepath.Join(BLS_KEYS_DIR, keyId, "key.json"))
	if err != nil {
		t.Fatalf("Failed to get BLS key dir: %s", err.Error())
	}
	passwordPath := filepath.Join(BLS_KEYS_DIR, keyId, "password.txt")
	password, err := os.ReadFile(passwordPath)
	if err != nil {
		t.Fatalf("Failed to read BLS password: %s", err.Error())
	}
	if string(password) != "" {
		t.Fatalf("Password is not empty: '%s'", password)
	}
	keyPair, err := bls.ReadPrivateKeyFromFile(nodeConfig.BlsPrivateKeyStorePath, string(password))
	if err != nil {
		t.Fatalf("Failed to generate operator BLS keys: %s", err.Error())
	}

	nodeConfig.EcdsaPrivateKeyStorePath, err = filepath.Abs(filepath.Join(ECDSA_KEYS_DIR, keyId, "key.json"))
	if err != nil {
		t.Fatalf("Failed to get ECDSA key dir: %s", err.Error())
	}
	passwordPath = filepath.Join(ECDSA_KEYS_DIR, keyId, "password.txt")
	password, err = os.ReadFile(passwordPath)
	if err != nil {
		t.Fatalf("Failed to read ECDSA password: %s", err.Error())
	}
	if string(password) != "" {
		t.Fatalf("Password is not empty: '%s'", password)
	}
	ecdsaKey, err := sdkEcdsa.ReadKey(nodeConfig.EcdsaPrivateKeyStorePath, string(password))
	if err != nil {
		t.Fatalf("Failed to generate operator ECDSA keys: %s", err.Error())
	}

	address := crypto.PubkeyToAddress(ecdsaKey.PublicKey)

	t.Logf("Generated config for operator: %s", address.String())

	nodeConfig.OperatorAddress = address.String()
	nodeConfig.RegisterOperatorOnStartup = true
	nodeConfig.EthRpcUrl = mainnetAnvil.HttpUrl
	nodeConfig.EthWsUrl = mainnetAnvil.WsUrl
	nodeConfig.RollupIdsToRpcUrls = make(map[uint32]string)
	nodeConfig.NearDaIndexerRollupIds = make([]uint32, 0, len(rollupAnvils))
	for _, rollupAnvil := range rollupAnvils {
		nodeConfig.RollupIdsToRpcUrls[uint32(rollupAnvil.ChainID.Uint64())] = rollupAnvil.WsUrl
		nodeConfig.NearDaIndexerRollupIds = append(nodeConfig.NearDaIndexerRollupIds, uint32(rollupAnvil.ChainID.Uint64()))
	}
	nodeConfig.EnableNodeApi = false
	nodeConfig.NodeApiIpPortAddress = "0.0.0.0:0"
	if keyId == "4" {
		// TODO: fix, Now impossible because of eigensdk
		nodeConfig.EnableMetrics = false
		nodeConfig.EigenMetricsIpPortAddress = "0.0.0.0:0"
	}

	nodeConfig.NearDaIndexerRmqIpPortAddress, err = rabbitMq.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Error getting AMQP URL: %s", err.Error())
	}

	mainnetAnvil.SetBalance(address, big.NewInt(1e18))

	return nodeConfig, keyPair, ecdsaKey
}

func buildConfigRaw(mainnetAnvil *utils.AnvilInstance, rollupAnvils []*utils.AnvilInstance) config.ConfigRaw {
	var configRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	sdkutils.ReadYamlConfig(aggConfigFilePath, &configRaw)
	configRaw.EthRpcUrl = mainnetAnvil.HttpUrl
	configRaw.EthWsUrl = mainnetAnvil.WsUrl
	configRaw.AggregatorDatabasePath = ""

	configRaw.RollupIdsToRpcUrls = map[uint32]string{}
	for _, el := range rollupAnvils {
		cleanedUrl := strings.TrimPrefix(el.HttpUrl, "http://")
		configRaw.RollupIdsToRpcUrls[uint32(el.ChainID.Uint64())] = cleanedUrl
	}

	return configRaw
}

func buildConfig(t *testing.T, sfflDeploymentRaw config.SFFLDeploymentRaw, addresses []common.Address, rollupAnvils []*utils.AnvilInstance, aggConfigRaw config.ConfigRaw) *config.Config {
	aggregatorEcdsaPrivateKeyString := "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	if aggregatorEcdsaPrivateKeyString[:2] == "0x" {
		aggregatorEcdsaPrivateKeyString = aggregatorEcdsaPrivateKeyString[2:]
	}
	aggregatorEcdsaPrivateKey, err := crypto.HexToECDSA(aggregatorEcdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}
	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(aggregatorEcdsaPrivateKey)
	if err != nil {
		t.Fatalf("Cannot get operator address: %s", err.Error())
	}

	rollupsInfo := make(map[uint32]config.RollupInfo)
	for i, addr := range addresses {
		rollupsInfo[uint32(rollupAnvils[i].ChainID.Int64())] = config.RollupInfo{SFFLRegistryRollupAddr: addr, RpcUrl: rollupAnvils[i].WsUrl}
	}

	return &config.Config{
		EcdsaPrivateKey:                aggregatorEcdsaPrivateKey,
		EthHttpRpcUrl:                  aggConfigRaw.EthRpcUrl,
		EthWsRpcUrl:                    aggConfigRaw.EthWsUrl,
		OperatorStateRetrieverAddr:     common.HexToAddress(sfflDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		SFFLRegistryCoordinatorAddr:    common.HexToAddress(sfflDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr:     aggConfigRaw.AggregatorServerIpPortAddr,
		AggregatorRestServerIpPortAddr: aggConfigRaw.AggregatorRestServerIpPortAddr,
		AggregatorDatabasePath:         aggConfigRaw.AggregatorDatabasePath,
		AggregatorCheckpointInterval:   time.Duration(aggConfigRaw.AggregatorCheckpointInterval) * time.Millisecond,
		RegisterOperatorOnStartup:      aggConfigRaw.RegisterOperatorOnStartup,
		AggregatorAddress:              aggregatorAddr,
		RollupsInfo:                    rollupsInfo,
		MetricsIpPortAddress:           aggConfigRaw.MetricsIpPortAddress,
	}
}

func startAnvilTestContainer(t *testing.T, ctx context.Context, name, exposedPort, chainId string, isMainnet bool, networkName string) *utils.AnvilInstance {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest",
		Name:         name,
		Entrypoint:   []string{"anvil"},
		ExposedPorts: []string{exposedPort + "/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
		Networks:     []string{networkName},
	}

	if isMainnet {
		req.Mounts = testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: filepath.Join(integrationDir, "../anvil/data/avs-and-eigenlayer-deployed-anvil-state.json"),
				},
				Target: "/root/.anvil/state.json",
			},
		}
		req.Cmd = []string{"--host", "0.0.0.0", "--load-state", "/root/.anvil/state.json", "--port", exposedPort, "--chain-id", chainId}
	} else {
		req.Cmd = []string{"--host", "0.0.0.0", "--port", exposedPort, "--block-time", "10", "--chain-id", chainId}
	}

	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Error starting anvil container: %s", err.Error())
	}

	anvilEndpoint, err := anvilC.PortEndpoint(ctx, nat.Port(exposedPort), "")
	if err != nil {
		t.Fatalf("Error getting anvil endpoint: %s", err.Error())
	}

	httpUrl := "http://" + anvilEndpoint
	httpClient, err := eth.NewClient(httpUrl)
	if err != nil {
		t.Fatalf("Failed to create anvil HTTP client: %s", err.Error())
	}
	rpcClient, err := rpc.Dial(httpUrl)
	if err != nil {
		t.Fatalf("Failed to create anvil RPC client: %s", err.Error())
	}

	wsUrl := "ws://" + anvilEndpoint
	wsClient, err := eth.NewClient(wsUrl)
	if err != nil {
		t.Fatalf("Failed to create anvil WS client: %s", err.Error())
	}

	expectedChainId, ok := big.NewInt(0).SetString(chainId, 10)
	if !ok {
		t.Fatalf("Bad chain ID: %s", chainId)
	}

	fetchedChainId, err := httpClient.ChainID(ctx)
	if err != nil {
		t.Fatalf("Failed to get anvil chainId: %s", err.Error())
	}
	if fetchedChainId.Cmp(expectedChainId) != 0 {
		t.Fatalf("Anvil chainId is not the expected: expected %s, got %s", expectedChainId.String(), fetchedChainId.String())
	}

	anvil := &utils.AnvilInstance{
		Container:  anvilC,
		HttpClient: httpClient,
		HttpUrl:    httpUrl,
		WsClient:   wsClient,
		WsUrl:      wsUrl,
		RpcClient:  rpcClient,
		ChainID:    fetchedChainId,
	}

	if isMainnet {
		anvil.Mine(big.NewInt(100), big.NewInt(1))
	}

	return anvil
}

func deployRegistryRollups(t *testing.T, anvils []*utils.AnvilInstance) ([]common.Address, []*registryrollup.ContractSFFLRegistryRollup, []*bind.TransactOpts, []*bind.TransactOpts) {
	var registryRollups []*registryrollup.ContractSFFLRegistryRollup
	var ownerAuths []*bind.TransactOpts
	var proxyAdminAuths []*bind.TransactOpts
	var addresses []common.Address

	for _, anvil := range anvils {
		addr, registryRollup, ownerAuth, proxyAdminAuth := deployRegistryRollup(t, anvil)

		addresses = append(addresses, addr)
		registryRollups = append(registryRollups, registryRollup)
		ownerAuths = append(ownerAuths, ownerAuth)
		proxyAdminAuths = append(proxyAdminAuths, proxyAdminAuth)
	}

	return addresses, registryRollups, ownerAuths, proxyAdminAuths
}

func deployRegistryRollup(t *testing.T, anvil *utils.AnvilInstance) (common.Address, *registryrollup.ContractSFFLRegistryRollup, *bind.TransactOpts, *bind.TransactOpts) {
	t.Logf("Deploying RegistryRollup to chain %s", anvil.ChainID.String())

	ownerPrivateKeyString := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ownerKeyPair, err := crypto.HexToECDSA(ownerPrivateKeyString)
	if err != nil {
		t.Fatalf("Error generating private key: %s", err.Error())
	}
	ownerAddr := crypto.PubkeyToAddress(ownerKeyPair.PublicKey)

	aggregatorAddr := common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")

	ownerAuth, err := bind.NewKeyedTransactorWithChainID(ownerKeyPair, anvil.ChainID)
	if err != nil {
		t.Fatalf("Error generating transactor: %s", err.Error())
	}

	implAddr, _, _, err := registryrollup.DeployContractSFFLRegistryRollup(ownerAuth, anvil.WsClient)
	if err != nil {
		t.Fatalf("Error deploying RegistryRollup: %s", err.Error())
	}

	abi, err := registryrollup.ContractSFFLRegistryRollupMetaData.GetAbi()
	if err != nil {
		t.Fatalf("Error getting RegistryRollup ABI: %s", err.Error())
	}

	mockPauserRegistryAddr := common.HexToAddress("0x000000000000000000000000000000000000001")

	initCall, err := abi.Pack("initialize", big.NewInt(66), ownerAddr, aggregatorAddr, mockPauserRegistryAddr)
	if err != nil {
		t.Fatalf("Error encoding RegistryRollup initialize call: %s", err.Error())
	}

	// using a separate account as proxy admin since it cannot access the fallback
	proxyAdminPrivateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	proxyAdminKeyPair, err := crypto.HexToECDSA(proxyAdminPrivateKeyString)
	if err != nil {
		t.Fatalf("Error generating private key: %s", err.Error())
	}
	proxyAdminAddr := crypto.PubkeyToAddress(proxyAdminKeyPair.PublicKey)

	proxyAdminAuth, err := bind.NewKeyedTransactorWithChainID(proxyAdminKeyPair, anvil.ChainID)
	if err != nil {
		t.Fatalf("Error generating transactor: %s", err.Error())
	}

	proxyAddr, _, _, err := transparentproxy.DeployContractTransparentUpgradeableProxy(
		proxyAdminAuth,
		anvil.WsClient,
		implAddr,
		proxyAdminAddr,
		initCall,
	)
	if err != nil {
		t.Fatalf("Error deploying RegistryRollup proxy: %s", err.Error())
	}

	registry, err := registryrollup.NewContractSFFLRegistryRollup(proxyAddr, anvil.WsClient)
	if err != nil {
		t.Fatalf("Error creating RegistryRollup instance: %s", err.Error())
	}

	return proxyAddr, registry, ownerAuth, proxyAdminAuth
}

func startIndexer(t *testing.T, ctx context.Context, name string, rollupAnvils []*utils.AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer, networkName string) (testcontainers.Container, []testcontainers.Container) {
	rmqName, err := rabbitMq.Name(ctx)
	if err != nil {
		t.Fatalf("Error getting RabbitMQ container name: %s", err.Error())
	}
	rmqName = strings.TrimPrefix(rmqName, "/")

	amqpUrl, err := rabbitMq.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Error getting RabbitMQ container URL: %s", err.Error())
	}
	amqpUrl = strings.Split(amqpUrl, "@")[0] + "@" + rmqName + ":" + "5672"

	var rollupArgs []string
	for _, rollupAnvil := range rollupAnvils {
		rollupArgs = append(rollupArgs, "--da-contract-ids", utils.GetDaContractAccountId(rollupAnvil))
	}
	for _, rollupAnvil := range rollupAnvils {
		rollupArgs = append(rollupArgs, "--rollup-ids", rollupAnvil.ChainID.String())
	}

	req := testcontainers.ContainerRequest{
		Image:        "near-sffl-indexer",
		Name:         name,
		Cmd:          append([]string{"--rmq-address", amqpUrl}, rollupArgs...),
		ExposedPorts: []string{"3030/tcp"},
		WaitingFor:   wait.ForLog("Starting Streamer..."),
		Networks:     []string{networkName},
	}

	genericReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	indexerContainer, err := testcontainers.GenericContainer(ctx, genericReq)

	if err != nil {
		t.Fatalf("Error starting indexer container: %s", err.Error())
	}

	relayers := setupNearDa(t, ctx, indexerContainer, rollupAnvils)
	return indexerContainer, relayers
}

func setupNearDa(t *testing.T, ctx context.Context, indexerContainer testcontainers.Container, rollupAnvils []*utils.AnvilInstance) []testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	indexerUrl, err := indexerContainer.Endpoint(ctx, "http")
	if err != nil {
		t.Fatalf("Error getting indexer endpoint: %s", err.Error())
	}

	indexerContainerIp, err := indexerContainer.ContainerIP(ctx)
	if err != nil {
		t.Fatalf("Error getting indxer container IP: %s", err.Error())
	}

	hostNearCfgPath := getNearCliConfigPath(t)
	hostNearKeyPath := filepath.Join(hostNearCfgPath, "validator_key.json")
	containerNearCfgPath := "/root/.near"

	time.Sleep(5 * time.Second)

	copyFileFromContainer(ctx, indexerContainer, filepath.Join(containerNearCfgPath, "validator_key.json"), hostNearKeyPath, 0770)

	var relayers []testcontainers.Container
	nearCliEnv := []string{"NEAR_ENV=" + utils.NearNetworkId, "NEAR_CLI_LOCALNET_NETWORK_ID=" + utils.NearNetworkId, "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH=" + hostNearKeyPath, "NEAR_NODE_URL=" + indexerUrl}
	for _, rollupAnvil := range rollupAnvils {
		accountId := utils.GetDaContractAccountId(rollupAnvil)

		err := execCommand(t, "near",
			[]string{"create-account", accountId, "--masterAccount", "test.near"},
			append(os.Environ(), nearCliEnv...),
			true,
		)
		if err != nil {
			t.Fatalf("Error creating NEAR DA account: %s", err.Error())
		}

		relayer, err := utils.StartRelayer(t, ctx, accountId, indexerContainerIp, rollupAnvil)
		if err != nil {
			t.Fatalf("Error creating realayer: #%s", err.Error())
		}
		relayers = append(relayers, relayer)

		err = execCommand(t, "near",
			[]string{"deploy", accountId, filepath.Join(integrationDir, "../near/near_da_blob_store.wasm"), "--initFunction", "new", "--initArgs", "{}", "--masterAccount", "test.near"},
			append(os.Environ(), nearCliEnv...),
			true,
		)
		if err != nil {
			t.Fatalf("Error deploying NEAR DA contract: %s", err.Error())
		}
	}

	return relayers
}

func execCommand(t *testing.T, name string, arg, env []string, shouldLog bool) error {
	cmd := exec.Command(name, arg...)
	cmd.Env = env
	out, err := cmd.CombinedOutput()
	if shouldLog {
		t.Log(string(out))
	}
	return err
}

func getStateRootUpdateAggregation(addr string, rollupID uint32, blockHeight uint64) (*aggtypes.GetStateRootUpdateAggregationResponse, error) {
	url := fmt.Sprintf("%s/aggregation/state-root-update?rollupId=%d&blockHeight=%d", addr, rollupID, blockHeight)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			return nil, fmt.Errorf("failed to read response body for error: %v", errRead)
		}

		return nil, fmt.Errorf("error: %s, message: %s", resp.Status, string(bodyBytes))
	}

	var response aggtypes.GetStateRootUpdateAggregationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func getOperatorSetUpdateAggregation(addr string, id uint64) (*aggtypes.GetOperatorSetUpdateAggregationResponse, error) {
	url := fmt.Sprintf("%s/aggregation/operator-set-update?id=%d", addr, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			return nil, fmt.Errorf("error: %s, message: %s", resp.Status, string(body))
		}

		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var response aggtypes.GetOperatorSetUpdateAggregationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, err
}

func copyFileFromContainer(ctx context.Context, container testcontainers.Container, sourcePath, destinationPath string, destinationPermissions fs.FileMode) error {
	reader, err := container.CopyFileFromContainer(ctx, sourcePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	err = os.MkdirAll(filepath.Dir(destinationPath), destinationPermissions)
	if err != nil {
		return err
	}

	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	return nil
}

func getNearCliConfigPath(t *testing.T) string {
	path, err := filepath.Abs(filepath.Join(TEST_DATA_DIR, "sffl_test_localnet"))
	if err != nil {
		t.Fatalf("Error getting near-cli config path: %s", err.Error())
	}
	return path
}
