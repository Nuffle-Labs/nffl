package integration_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/NethermindEth/near-sffl/core"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/NethermindEth/near-sffl/aggregator"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/NethermindEth/near-sffl/types"
)

const TEST_DATA_DIR = "../../test_data"

func TestIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
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

	stateRootHeight := uint64(10)
	stateRootUpdate, err := getStateRootUpdateAggregation(setup.aggregatorRestUrl, uint32(setup.rollupAnvils[0].ChainID.Uint64()), stateRootHeight)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}
	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message), core.FormatBlsAggregationRollup(&stateRootUpdate.Aggregation))
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	newOperatorConfig, _, _ := genOperatorConfig(t, ctx, setup.mainnetAnvil, setup.rollupAnvils, setup.rabbitMq)
	newOperator := startOperator(t, ctx, newOperatorConfig)

	time.Sleep(30 * time.Second)

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
	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message), core.FormatBlsAggregationRollup(&stateRootUpdate.Aggregation))
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

	operatorSetUpdate, err := getOperatorSetUpdateAggregation(setup.aggregatorRestUrl, operatorSetUpdateCount-1)
	if err != nil {
		t.Fatalf("Error getting operator set update: %s", err.Error())
	}

	expectedUpdatedOperators := []registryrollup.OperatorsOperator{
		{
			Pubkey: registryrollup.BN254G1Point{
				X: newOperator.BlsPubkeyG1().X.BigInt(big.NewInt(0)),
				Y: newOperator.BlsPubkeyG1().Y.BigInt(big.NewInt(0)),
			},
			Weight: big.NewInt(1000),
		},
	}
	assert.Equal(t, expectedUpdatedOperators, operatorSetUpdate.Message.Operators)

	<-ctx.Done()
}

type testEnv struct {
	mainnetAnvil        *AnvilInstance
	rollupAnvils        []*AnvilInstance
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
	rollupAnvils := []*AnvilInstance{
		startAnvilTestContainer(t, containersCtx, rollup0AnvilContainerName, "8546", "2", false, networkName),
		startAnvilTestContainer(t, containersCtx, rollup1AnvilContainerName, "8547", "3", false, networkName),
	}
	rabbitMq := startRabbitMqContainer(t, containersCtx, rmqContainerName, networkName)
	indexerContainer := startIndexer(t, containersCtx, indexerContainerName, rollupAnvils, rabbitMq, networkName)

	startRollupIndexing(t, ctx, rollupAnvils, indexerContainer)

	sfflDeploymentRaw := readSfflDeploymentRaw()

	configRaw := buildConfigRaw(mainnetAnvil, rollupAnvils)
	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

	nodeConfig, keyPair, _ := genOperatorConfig(t, ctx, mainnetAnvil, rollupAnvils, rabbitMq)

	avsReader, err := chainio.BuildAvsReader(common.HexToAddress(sfflDeploymentRaw.Addresses.RegistryCoordinatorAddr), common.HexToAddress(sfflDeploymentRaw.Addresses.OperatorStateRetrieverAddr), mainnetAnvil.HttpClient, logger)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	rollupInitialOperatorSet := []registryrollup.OperatorsOperator{
		{
			Pubkey: registryrollup.BN254G1Point{
				X: keyPair.PubKey.X.BigInt(big.NewInt(0)),
				Y: keyPair.PubKey.Y.BigInt(big.NewInt(0)),
			},
			Weight: big.NewInt(1000),
		},
	}

	addresses, registryRollups, registryRollupAuths := deployRegistryRollups(t, ctx, rollupInitialOperatorSet, 1, avsReader, rollupAnvils)

	operator := startOperator(t, ctx, nodeConfig)

	config := buildConfig(t, sfflDeploymentRaw, addresses, rollupAnvils, configRaw)
	aggregator := startAggregator(t, ctx, config, logger)

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

func startOperator(t *testing.T, ctx context.Context, nodeConfig types.NodeConfig) *operator.Operator {
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

func genOperatorConfig(t *testing.T, ctx context.Context, mainnetAnvil *AnvilInstance, rollupAnvils []*AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) (types.NodeConfig, *bls.KeyPair, *ecdsa.PrivateKey) {
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err := sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	log.Println("starting operator for integration tests")

	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")

	err = os.MkdirAll(getOperatorKeysPathPrefix(t), 0770)
	if err != nil {
		t.Fatalf("Failed to create operators keys dir: %s", err.Error())
	}

	keysPath, err := os.MkdirTemp(getOperatorKeysPathPrefix(t), "")
	if err != nil {
		t.Fatalf("Failed to create operator keys dir: %s", err.Error())
	}

	nodeConfig.BlsPrivateKeyStorePath, err = filepath.Abs(filepath.Join(keysPath, "test.bls.key.json"))
	if err != nil {
		t.Fatalf("Failed to get BLS key dir: %s", err.Error())
	}
	keyPair, err := bls.GenRandomBlsKeys()
	if err != nil {
		t.Fatalf("Failed to generate operator BLS keys: %s", err.Error())
	}
	err = keyPair.SaveToFile(nodeConfig.BlsPrivateKeyStorePath, "")
	if err != nil {
		t.Fatalf("Failed to save operator BLS keys: %s", err.Error())
	}

	nodeConfig.EcdsaPrivateKeyStorePath, err = filepath.Abs(filepath.Join(keysPath, "test.ecdsa.key.json"))
	if err != nil {
		t.Fatalf("Failed to get ECDSA key dir: %s", err.Error())
	}
	ecdsaKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to save generate operator ECDSA key: %s", err.Error())
	}
	sdkEcdsa.WriteKey(nodeConfig.EcdsaPrivateKeyStorePath, ecdsaKey, "")
	if err != nil {
		t.Fatalf("Failed to save operator ECDSA keys: %s", err.Error())
	}

	address := crypto.PubkeyToAddress(ecdsaKey.PublicKey)

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
	nodeConfig.EnableMetrics = false
	nodeConfig.EigenMetricsIpPortAddress = "0.0.0.0:0"

	nodeConfig.NearDaIndexerRmqIpPortAddress, err = rabbitMq.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Error getting AMQP URL: %s", err.Error())
	}

	mainnetAnvil.setBalance(address, big.NewInt(1e18))

	return nodeConfig, keyPair, ecdsaKey
}

func buildConfigRaw(mainnetAnvil *AnvilInstance, rollupAnvils []*AnvilInstance) config.ConfigRaw {
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

func buildConfig(t *testing.T, sfflDeploymentRaw config.SFFLDeploymentRaw, addresses []common.Address, rollupAnvils []*AnvilInstance, aggConfigRaw config.ConfigRaw) *config.Config {
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
		RegisterOperatorOnStartup:      aggConfigRaw.RegisterOperatorOnStartup,
		AggregatorAddress:              aggregatorAddr,
		RollupsInfo:                    rollupsInfo,
	}
}

type AnvilInstance struct {
	Container  testcontainers.Container
	HttpClient *eth.Client
	HttpUrl    string
	WsClient   *eth.Client
	WsUrl      string
	ChainID    *big.Int
}

func startAnvilTestContainer(t *testing.T, ctx context.Context, name, exposedPort, chainId string, isMainnet bool, networkName string) *AnvilInstance {
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

	anvil := &AnvilInstance{
		Container:  anvilC,
		HttpClient: httpClient,
		HttpUrl:    httpUrl,
		WsClient:   wsClient,
		WsUrl:      wsUrl,
		ChainID:    fetchedChainId,
	}

	if isMainnet {
		anvil.mine(big.NewInt(100), big.NewInt(1))
	}

	return anvil
}

func deployRegistryRollups(t *testing.T, ctx context.Context, initialOperatorSet []registryrollup.OperatorsOperator, nextOperatorSetUpdateId uint64, avsReader chainio.AvsReaderer, anvils []*AnvilInstance) ([]common.Address, []*registryrollup.ContractSFFLRegistryRollup, []*bind.TransactOpts) {
	var registryRollups []*registryrollup.ContractSFFLRegistryRollup
	var auths []*bind.TransactOpts
	var addresses []common.Address

	for _, anvil := range anvils {
		addr, registryRollup, auth := deployRegistryRollup(t, ctx, initialOperatorSet, nextOperatorSetUpdateId, avsReader, anvil)

		registryRollups = append(registryRollups, registryRollup)
		auths = append(auths, auth)
		addresses = append(addresses, addr)
	}

	return addresses, registryRollups, auths
}

func deployRegistryRollup(t *testing.T, ctx context.Context, initialOperatorSet []registryrollup.OperatorsOperator, nextOperatorSetUpdateId uint64, avsReader chainio.AvsReaderer, anvil *AnvilInstance) (common.Address, *registryrollup.ContractSFFLRegistryRollup, *bind.TransactOpts) {
	t.Logf("Deploying RegistryRollup to chain %s", anvil.ChainID.String())

	privateKeyString := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		t.Fatalf("Error generating private key: %s", err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, anvil.ChainID)
	if err != nil {
		t.Fatalf("Error generating transactor: %s", err.Error())
	}

	if len(initialOperatorSet) == 0 {
		t.Fatal("Operator set is empty")
	}

	t.Logf("RegistryRollup deployed with operators: %v", initialOperatorSet)

	addr, _, registryRollup, err := registryrollup.DeployContractSFFLRegistryRollup(auth, anvil.WsClient, initialOperatorSet, big.NewInt(66), nextOperatorSetUpdateId)
	if err != nil {
		t.Fatalf("Error deploying RegistryRollup: %s", err.Error())
	}

	return addr, registryRollup, auth
}

func startRollupIndexing(t *testing.T, ctx context.Context, rollupAnvils []*AnvilInstance, indexerContainer testcontainers.Container) {
	headers := make(chan *ethtypes.Header)

	indexerUrl, err := indexerContainer.Endpoint(ctx, "http")
	if err != nil {
		t.Fatalf("Error getting indexer endpoint: %s", err.Error())
	}

	for _, rollupAnvil := range rollupAnvils {
		anvil := rollupAnvil

		sub, err := anvil.WsClient.SubscribeNewHead(ctx, headers)
		if err != nil {
			t.Fatalf("Error subscribing to new rollup block headers: %s", err.Error())
		}

		go func() {
			for {
				select {
				case err := <-sub.Err():
					t.Errorf("Error on rollup block subscription: %s", err.Error())
					return
				case header := <-headers:
					t.Logf("Got rollup block header: #%s", header.Number.String())

					var block *ethtypes.Block

					for i := 0; i < 5; i++ {
						select {
						case <-ctx.Done():
							return
						default:
						}

						block, err = anvil.HttpClient.BlockByNumber(ctx, header.Number)

						if err != nil {
							t.Logf("Did not find rollup block: %s", err.Error())
							time.Sleep(1 * time.Second)
						} else {
							break
						}
					}

					if block == nil {
						t.Error("Could not fetch rollup block")
						return
					}

					submitBlock(t, ctx, getDaContractAccountId(anvil), block, indexerUrl)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func startIndexer(t *testing.T, ctx context.Context, name string, rollupAnvils []*AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer, networkName string) testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

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
		rollupArgs = append(rollupArgs, "--da-contract-ids", getDaContractAccountId(rollupAnvil))
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

	indexerContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Error starting indexer container: %s", err.Error())
	}

	indexerUrl, err := indexerContainer.Endpoint(ctx, "http")
	if err != nil {
		t.Fatalf("Error getting indexer endpoint: %s", err.Error())
	}

	hostNearCfgPath := getNearCliConfigPath(t)
	hostNearKeyPath := filepath.Join(hostNearCfgPath, "validator_key.json")
	containerNearCfgPath := "/root/.near"

	time.Sleep(5 * time.Second)

	copyFileFromContainer(ctx, indexerContainer, filepath.Join(containerNearCfgPath, "validator_key.json"), hostNearKeyPath, 0770)

	for _, rollupAnvil := range rollupAnvils {
		accountId := getDaContractAccountId(rollupAnvil)

		err = execCommand(t, "near",
			[]string{"create-account", accountId, "--masterAccount", "test.near"},
			append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+hostNearKeyPath, "NEAR_NODE_URL="+indexerUrl),
			true,
		)
		if err != nil {
			t.Fatalf("Error creating NEAR DA account: %s", err.Error())
		}

		err = execCommand(t, "near",
			[]string{"deploy", accountId, filepath.Join(integrationDir, "../near/near_da_blob_store.wasm"), "--initFunction", "new", "--initArgs", "{}", "--masterAccount", "test.near"},
			append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+hostNearKeyPath, "NEAR_NODE_URL="+indexerUrl),
			true,
		)
		if err != nil {
			t.Fatalf("Error deploying NEAR DA contract: %s", err.Error())
		}
	}

	return indexerContainer
}

func getDaContractAccountId(anvil *AnvilInstance) string {
	return fmt.Sprintf("da%s.test.near", anvil.ChainID.String())
}

func submitBlock(t *testing.T, ctx context.Context, accountId string, block *ethtypes.Block, indexerUrl string) error {
	t.Log("Submitting block to NEAR DA")

	encodedBlock, err := rlp.EncodeToBytes(block)
	if err != nil {
		return err
	}

	keyPath := filepath.Join(getNearCliConfigPath(t), "validator_key.json")

	err = execCommand(t, "near",
		[]string{"call", accountId, "submit", "--base64", base64.StdEncoding.EncodeToString(encodedBlock), "--accountId", accountId},
		append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+keyPath, "NEAR_NODE_URL="+indexerUrl),
		false,
	)
	if err != nil {
		return err
	}

	return nil
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

func getStateRootUpdateAggregation(addr string, rollupID uint32, blockHeight uint64) (*aggregator.GetStateRootUpdateAggregationResponse, error) {
	url := fmt.Sprintf("%s/aggregation/state-root-update?rollupId=%d&blockHeight=%d", addr, rollupID, blockHeight)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	var response aggregator.GetStateRootUpdateAggregationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func getOperatorSetUpdateAggregation(addr string, id uint64) (*aggregator.GetOperatorSetUpdateAggregationResponse, error) {
	url := fmt.Sprintf("%s/aggregation/operator-set-update?id=%d", addr, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	var response aggregator.GetOperatorSetUpdateAggregationResponse
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

func getOperatorKeysPathPrefix(t *testing.T) string {
	path, err := filepath.Abs(filepath.Join(TEST_DATA_DIR, "sffl_test_operators"))
	if err != nil {
		t.Fatalf("Error getting operator keys path prefix: %s", err.Error())
	}
	return path
}

func (ai *AnvilInstance) setBalance(address common.Address, balance *big.Int) error {
	return ai.WsClient.Client.Client().Call(nil, "anvil_setBalance", address.Hex(), "0x"+balance.Text(16))
}

func (ai *AnvilInstance) mine(blockCount, timestampInterval *big.Int) error {
	return ai.WsClient.Client.Client().Call(nil, "anvil_mine", "0x"+blockCount.Text(16), "0x"+timestampInterval.Text(16))
}
