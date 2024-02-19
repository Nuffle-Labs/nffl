package integration_test

import (
	"bytes"
	"context"
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
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/NethermindEth/near-sffl/aggregator"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/NethermindEth/near-sffl/types"
)

func TestIntegration(t *testing.T) {
	setup := setupTestEnv(t)

	time.Sleep(10 * time.Second)

	taskHash, err := setup.avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskHashes(&bind.CallOpts{}, 1)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	taskResponseHash, err := setup.avsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskResponses(&bind.CallOpts{}, 1)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
	}

	stateRootUpdate, err := getStateRootUpdateAggregation(setup.aggregatorRestUrl, uint32(setup.rollupAnvils[0].ChainID.Uint64()), 7)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}

	_, err = setup.registryRollups[1].UpdateStateRoot(setup.registryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message), formatBlsAggregationRollup(t, &stateRootUpdate.Aggregation))
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}
}

type TestEnv struct {
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
}

func setupTestEnv(t *testing.T) *TestEnv {
	t.Log("This test takes ~100 seconds to run...")

	containersCtx, cancelContainersCtx := context.WithCancel(context.Background())

	mainnetAnvil := startAnvilTestContainer(t, containersCtx, "8545", "1", true)
	rollupAnvils := []*AnvilInstance{
		startAnvilTestContainer(t, containersCtx, "8546", "2", false),
		startAnvilTestContainer(t, containersCtx, "8547", "3", false),
	}
	rabbitMq := startRabbitMqContainer(t, containersCtx)
	indexerContainer := startIndexer(t, containersCtx, rollupAnvils, rabbitMq)

	startRollupIndexing(t, containersCtx, rollupAnvils)

	sfflDeploymentRaw := readSfflDeploymentRaw()

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	nodeConfig := buildOperatorConfig(t, ctx, mainnetAnvil, rollupAnvils, rabbitMq)
	config := buildConfig(t, sfflDeploymentRaw, mainnetAnvil)

	operator := startOperator(t, ctx, nodeConfig)
	aggregator := startAggregator(t, ctx, config)

	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	registryRollups, registryRollupAuths := deployRegistryRollups(t, ctx, avsReader, rollupAnvils)

	t.Cleanup(func() {
		cancel()

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

		cancelContainersCtx()
	})

	return &TestEnv{
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
		func() testcontainers.CustomizeRequestOption {
			return func(req *testcontainers.GenericContainerRequest) {
				req.HostConfigModifier = func(hc *container.HostConfig) {
					hc.NetworkMode = "host"
				}
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

func buildOperatorConfig(t *testing.T, ctx context.Context, mainnetAnvil *AnvilInstance, rollupAnvils []*AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) types.NodeConfig {
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
	nodeConfig.NearDaIndexerRollupIds = make([]uint32, 0, len(rollupAnvils))
	for _, rollupAnvil := range rollupAnvils {
		nodeConfig.RollupIdsToRpcUrls[uint32(rollupAnvil.ChainID.Uint64())] = rollupAnvil.WsUrl
		nodeConfig.NearDaIndexerRollupIds = append(nodeConfig.NearDaIndexerRollupIds, uint32(rollupAnvil.ChainID.Uint64()))
	}

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
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.NetworkMode = "host"
		},
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

	if isMainnet {
		advanceChain(t, httpUrl)
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

func deployRegistryRollups(t *testing.T, ctx context.Context, avsReader chainio.AvsReaderer, anvils []*AnvilInstance) ([]*registryrollup.ContractSFFLRegistryRollup, []*bind.TransactOpts) {
	var registryRollups []*registryrollup.ContractSFFLRegistryRollup
	var auths []*bind.TransactOpts

	for _, anvil := range anvils {
		registryRollup, auth := deployRegistryRollup(t, ctx, avsReader, anvil)

		registryRollups = append(registryRollups, registryRollup)
		auths = append(auths, auth)
	}

	return registryRollups, auths
}

func deployRegistryRollup(t *testing.T, ctx context.Context, avsReader chainio.AvsReaderer, anvil *AnvilInstance) (*registryrollup.ContractSFFLRegistryRollup, *bind.TransactOpts) {
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

	operatorSetUpdateId := uint64(0)

	operatorsDelta, err := avsReader.GetOperatorSetUpdateDelta(ctx, operatorSetUpdateId)
	if err != nil {
		t.Fatalf("Error getting operator set: %s", err.Error())
	}

	var operators []registryrollup.OperatorsOperator
	for _, operator := range operatorsDelta {
		operators = append(operators, registryrollup.OperatorsOperator{
			Pubkey: registryrollup.BN254G1Point(operator.Pubkey),
			Weight: operator.Weight,
		})
	}

	if len(operators) == 0 {
		t.Fatal("Operator set is empty")
	}

	t.Logf("RegistryRollup deployed with operators: %v", operators)

	_, _, registryRollup, err := registryrollup.DeployContractSFFLRegistryRollup(auth, anvil.WsClient, operators, big.NewInt(66), operatorSetUpdateId+1)
	if err != nil {
		t.Fatalf("Error deploying RegistryRollup: %s", err.Error())
	}

	return registryRollup, auth
}

func startRollupIndexing(t *testing.T, ctx context.Context, rollupAnvils []*AnvilInstance) {
	headers := make(chan *ethtypes.Header)

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
					panic(fmt.Errorf("Error on rollup block subscription: %s", err.Error()))
				case header := <-headers:
					t.Logf("Got rollup block: #%s", header.Number.String())
					block, err := anvil.WsClient.BlockByNumber(ctx, header.Number)
					if err != nil {
						panic(fmt.Errorf("Error getting rollup block: %s", err.Error()))
					}
					submitBlock(t, getDaContractAccountId(anvil), block)
					return
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func startIndexer(t *testing.T, ctx context.Context, rollupAnvils []*AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	amqpUrl, err := rabbitMq.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Error getting RabbitMQ container URL: %s", err.Error())
	}

	var rollupArgs []string
	for _, rollupAnvil := range rollupAnvils {
		rollupArgs = append(rollupArgs, "--da-contract-ids", getDaContractAccountId(rollupAnvil))
	}
	for _, rollupAnvil := range rollupAnvils {
		rollupArgs = append(rollupArgs, "--rollup-ids", rollupAnvil.ChainID.String())
	}

	req := testcontainers.ContainerRequest{
		Image: "near-sffl-indexer",
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.NetworkMode = "host"
		},
		Cmd:        append([]string{"--rmq-address", amqpUrl}, rollupArgs...),
		WaitingFor: wait.ForLog("Starting Streamer..."),
	}

	indexerContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Error starting indexer container: %s", err.Error())
	}

	hostNearCfgPath := getNearConfigPath()
	hostNearKeyPath := filepath.Join(hostNearCfgPath, "validator_key.json")
	containerNearCfgPath := "/root/.near"

	time.Sleep(5 * time.Second)

	copyFileFromContainer(ctx, indexerContainer, filepath.Join(containerNearCfgPath, "validator_key.json"), hostNearKeyPath, 0770)

	for _, rollupAnvil := range rollupAnvils {
		accountId := getDaContractAccountId(rollupAnvil)

		err = execCommand(t, "near",
			[]string{"create-account", accountId, "--masterAccount", "test.near"},
			append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+hostNearKeyPath),
			true,
		)
		if err != nil {
			t.Fatalf("Error creating NEAR DA account: %s", err.Error())
		}

		err = execCommand(t, "near",
			[]string{"deploy", accountId, filepath.Join(integrationDir, "../near/near_da_blob_store.wasm"), "--initFunction", "new", "--initArgs", "{}", "--masterAccount", "test.near"},
			append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+hostNearKeyPath),
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

func getNearConfigPath() string {
	return "/tmp/sffl_test_localnet"
}

func submitBlock(t *testing.T, accountId string, block *ethtypes.Block) error {
	t.Log("Submitting block to NEAR DA")

	encodedBlock, err := rlp.EncodeToBytes(block)
	if err != nil {
		return err
	}

	err = execCommand(t, "near",
		[]string{"call", accountId, "submit", "--base64", base64.StdEncoding.EncodeToString(encodedBlock), "--accountId", accountId},
		append(os.Environ(), "NEAR_ENV=localnet", "NEAR_HELPER_ACCOUNT=near", "NEAR_CLI_LOCALNET_KEY_PATH="+filepath.Join(getNearConfigPath(), "validator_key.json")),
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

func advanceChain(t *testing.T, rpcUrl string) {
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(
			`forge script script/utils/Utils.sol --sig "advanceChainByNBlocks(uint256)" 100 --rpc-url %s --private-key %s --broadcast`,
			rpcUrl, privateKey),
	)
	cmd.Dir = "../../contracts/evm"

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Error advancing chain: %s", stderr.String())
	}
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

func formatBlsAggregationRollup(t *testing.T, agg *aggtypes.MessageBlsAggregationServiceResponse) registryrollup.OperatorsSignatureInfo {
	var nonSignerPubkeys []registryrollup.BN254G1Point

	for _, pubkey := range agg.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, registryrollup.BN254G1Point{
			X: pubkey.X.BigInt(big.NewInt(0)),
			Y: pubkey.Y.BigInt(big.NewInt(0)),
		})
	}

	apkG2 := registryrollup.BN254G2Point{
		X: [2]*big.Int{agg.SignersApkG2.X.A1.BigInt(big.NewInt(0)), agg.SignersApkG2.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{agg.SignersApkG2.Y.A1.BigInt(big.NewInt(0)), agg.SignersApkG2.Y.A0.BigInt(big.NewInt(0))},
	}

	sigma := registryrollup.BN254G1Point{
		X: agg.SignersAggSigG1.X.BigInt(big.NewInt(0)),
		Y: agg.SignersAggSigG1.Y.BigInt(big.NewInt(0)),
	}

	return registryrollup.OperatorsSignatureInfo{
		NonSignerPubkeys: nonSignerPubkeys,
		ApkG2:            apkG2,
		Sigma:            sigma,
	}
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
