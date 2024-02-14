package integration_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/NethermindEth/near-sffl/aggregator"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/NethermindEth/near-sffl/types"
)

func TestIntegration(t *testing.T) {
	t.Log("This test takes ~100 seconds to run...")

	containersCtx, cancelContainersCtx := context.WithCancel(context.Background())

	mainnetAnvil := startAnvilTestContainer(t, containersCtx, "8545", "1", true)
	rollupAnvil := startAnvilTestContainer(t, containersCtx, "8547", "2", false)
	rabbitMq := startRabbitMqContainer(t, containersCtx)

	time.Sleep(4 * time.Second)

	sfflDeploymentRaw := readSfflDeploymentRaw()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	nodeConfig := buildOperatorConfig(t, ctx, mainnetAnvil, rollupAnvil, rabbitMq)
	config := buildConfig(t, sfflDeploymentRaw, mainnetAnvil)

	_ = startOperator(t, ctx, nodeConfig)
	_ = startAggregator(t, ctx, config)

	t.Cleanup(func() {
		cancel()

		time.Sleep(5 * time.Second)

		if err := mainnetAnvil.Container.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		if err := rollupAnvil.Container.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		if err := rabbitMq.Terminate(containersCtx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}

		cancelContainersCtx()
	})

	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	taskHash, err := avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskHashes(&bind.CallOpts{}, 1)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	taskResponseHash, err := avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskResponses(&bind.CallOpts{}, 1)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
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

func startAggregator(t *testing.T, ctx context.Context, config *config.Config) *aggregator.Aggregator {
	t.Log("starting aggregator for integration tests")

	agg, err := aggregator.NewAggregator(ctx, config)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}

	go agg.Start(ctx)

	t.Log("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	return agg
}

func startRabbitMqContainer(t *testing.T, ctx context.Context) *rabbitmq.RabbitMQContainer {
	rabbitMqC, err := rabbitmq.RunContainer(
		ctx,
		testcontainers.WithImage("rabbitmq:latest"),
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

func buildOperatorConfig(t *testing.T, ctx context.Context, mainnetAnvil, rollupAnvil *AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) types.NodeConfig {
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err := sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")
	nodeConfig.BlsPrivateKeyStorePath = "../keys/test.bls.key.json"
	nodeConfig.EcdsaPrivateKeyStorePath = "../keys/test.ecdsa.key.json"
	nodeConfig.RegisterOperatorOnStartup = true
	nodeConfig.EthRpcUrl = mainnetAnvil.HttpUrl
	nodeConfig.EthWsUrl = mainnetAnvil.WsUrl
	nodeConfig.RollupIdsToRpcUrls = make(map[uint32]string)
	nodeConfig.RollupIdsToRpcUrls[uint32(rollupAnvil.ChainID.Uint64())] = rollupAnvil.WsUrl
	nodeConfig.NearDaIndexerRollupIds = []uint32{uint32(rollupAnvil.ChainID.Uint64())}

	amqpUrl, err := rabbitMq.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Error getting AMQP URL: %s", err.Error())
	}

	nodeConfig.NearDaIndexerRmqIpPortAddress = amqpUrl

	return nodeConfig
}

func buildConfig(t *testing.T, sfflDeploymentRaw config.SFFLDeploymentRaw, mainnetAnvil *AnvilInstance) *config.Config {
	var aggConfigRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	sdkutils.ReadYamlConfig(aggConfigFilePath, &aggConfigRaw)
	aggConfigRaw.EthRpcUrl = mainnetAnvil.HttpUrl
	aggConfigRaw.EthWsUrl = mainnetAnvil.WsUrl
	aggConfigRaw.AggregatorDatabasePath = ""

	logger, err := sdklogging.NewZapLogger(aggConfigRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

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

	privateKeySigner, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: aggregatorEcdsaPrivateKey}, mainnetAnvil.ChainID)
	if err != nil {
		t.Fatalf("Cannot create signer: %s", err.Error())
	}
	txMgr := txmgr.NewSimpleTxManager(mainnetAnvil.HttpClient, logger, privateKeySigner, aggregatorAddr)

	return &config.Config{
		EcdsaPrivateKey:                aggregatorEcdsaPrivateKey,
		Logger:                         logger,
		EthHttpRpcUrl:                  aggConfigRaw.EthRpcUrl,
		EthHttpClient:                  mainnetAnvil.HttpClient,
		EthWsRpcUrl:                    aggConfigRaw.EthWsUrl,
		EthWsClient:                    mainnetAnvil.WsClient,
		OperatorStateRetrieverAddr:     common.HexToAddress(sfflDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		SFFLRegistryCoordinatorAddr:    common.HexToAddress(sfflDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr:     aggConfigRaw.AggregatorServerIpPortAddr,
		AggregatorRestServerIpPortAddr: aggConfigRaw.AggregatorRestServerIpPortAddr,
		AggregatorDatabasePath:         aggConfigRaw.AggregatorDatabasePath,
		RegisterOperatorOnStartup:      aggConfigRaw.RegisterOperatorOnStartup,
		TxMgr:                          txMgr,
		AggregatorAddress:              aggregatorAddr,
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

func startAnvilTestContainer(t *testing.T, ctx context.Context, exposedPort, chainId string, isMainnet bool) *AnvilInstance {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest",
		Entrypoint:   []string{"anvil"},
		ExposedPorts: []string{exposedPort + "/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
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
		req.Cmd = []string{"--host", "0.0.0.0", "--port", exposedPort, "--chain-id", chainId}
	}

	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Error starting anvil container: %s", err.Error())
	}

	if isMainnet {
		advanceChain(t, anvilC)
	}

	anvilEndpoint, err := anvilC.Endpoint(ctx, "")
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

	return &AnvilInstance{
		Container:  anvilC,
		HttpClient: httpClient,
		HttpUrl:    httpUrl,
		WsClient:   wsClient,
		WsUrl:      wsUrl,
		ChainID:    fetchedChainId,
	}
}

func advanceChain(t *testing.T, anvilC testcontainers.Container) {
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		panic(err)
	}
	rpcUrl := "http://" + anvilEndpoint
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(
			`forge script script/utils/Utils.sol --sig "advanceChainByNBlocks(uint256)" 100 --rpc-url %s --private-key %s --broadcast`,
			rpcUrl, privateKey),
	)
	cmd.Dir = "../../contracts/evm"

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		t.Fatalf("Error advancing chain: %s", stderr.String())
	}
}
