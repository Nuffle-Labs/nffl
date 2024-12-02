package integration

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Nuffle-Labs/nffl/aggregator"
	restserver "github.com/Nuffle-Labs/nffl/aggregator/rest_server"
	rpcserver "github.com/Nuffle-Labs/nffl/aggregator/rpc_server"
	aggtypes "github.com/Nuffle-Labs/nffl/aggregator/types"
	registryrollup "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLRegistryRollup"
	transparentproxy "github.com/Nuffle-Labs/nffl/contracts/bindings/TransparentUpgradeableProxy"
	"github.com/Nuffle-Labs/nffl/core/chainio"
	"github.com/Nuffle-Labs/nffl/core/config"
	"github.com/Nuffle-Labs/nffl/operator"
	optypes "github.com/Nuffle-Labs/nffl/operator/types"
	"github.com/Nuffle-Labs/nffl/tests/integration/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	_ "github.com/testcontainers/testcontainers-go/wait"
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
)

const (
	TEST_DATA_DIR  = "../../test_data"
	BLS_KEYS_DIR   = "../keys/bls"
	ECDSA_KEYS_DIR = "../keys/ecdsa"
)

type TestEnv struct {
	MainnetAnvil        *utils.AnvilInstance
	RollupAnvils        []*utils.AnvilInstance
	RabbitMq            *rabbitmq.RabbitMQContainer
	IndexerContainer    testcontainers.Container
	Operator            *operator.Operator
	Aggregator          *aggregator.Aggregator
	AggregatorRestUrl   string
	AvsReader           *chainio.AvsReader
	RegistryRollups     []*registryrollup.ContractSFFLRegistryRollup
	RegistryRollupAuths []*bind.TransactOpts
	Cleanup             func()
}

func StartOperator(t *testing.T, ctx context.Context, nodeConfig optypes.NodeConfig) *operator.Operator {
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		t.Fatalf("Failed to create operator: %s", err.Error())
	}

	go operator.Start(ctx)

	t.Log("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	return operator
}

func StartAggregator(t *testing.T, ctx context.Context, config *config.Config, logger sdklogging.Logger) *aggregator.Aggregator {
	t.Log("starting aggregator for integration tests")

	var optRegistry *prometheus.Registry
	if config.EnableMetrics {
		optRegistry = prometheus.NewRegistry()
	}
	agg, err := aggregator.NewAggregator(ctx, config, nil, logger)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}

	rpcServer := rpcserver.NewRpcServer(config.AggregatorServerIpPortAddr, agg, logger)
	if optRegistry != nil {
		err = rpcServer.EnableMetrics(optRegistry)
		if err != nil {
			t.Fatalf("Failed to create metrics for rpc server: %s", err.Error())
		}
	}
	go rpcServer.Start()

	restServer := restserver.NewRestServer(config.AggregatorRestServerIpPortAddr, agg, logger)
	if optRegistry != nil {
		err = restServer.EnableMetrics(optRegistry)
		if err != nil {
			t.Fatalf("Failed to create metrics for rest server: %s", err.Error())
		}
	}
	go restServer.Start()

	go agg.Start(ctx)

	t.Log("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	return agg
}

func StartRabbitMqContainer(t *testing.T, ctx context.Context, name, networkName string) *rabbitmq.RabbitMQContainer {
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

func ReadSfflDeploymentRaw() config.SFFLDeploymentRaw {
	var sfflDeploymentRaw config.SFFLDeploymentRaw
	sfflDeploymentFilePath := "../../contracts/evm/script/output/31337/sffl_avs_deployment_output.json"
	sdkutils.ReadJsonConfig(sfflDeploymentFilePath, &sfflDeploymentRaw)

	return sfflDeploymentRaw
}

func GenOperatorConfig(t *testing.T, ctx context.Context, keyId string, mainnetAnvil *utils.AnvilInstance, rollupAnvils []*utils.AnvilInstance, rabbitMq *rabbitmq.RabbitMQContainer) (optypes.NodeConfig, *bls.KeyPair, *ecdsa.PrivateKey) {
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
		// TODO: fix, Now impossible due to eigensdk limitations
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

func BuildConfigRaw(mainnetAnvil *utils.AnvilInstance, rollupAnvils []*utils.AnvilInstance) config.ConfigRaw {
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

func BuildConfig(t *testing.T, sfflDeploymentRaw config.SFFLDeploymentRaw, addresses []common.Address, rollupAnvils []*utils.AnvilInstance, aggConfigRaw config.ConfigRaw) *config.Config {
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
		EnableMetrics:                  false,
		MetricsIpPortAddress:           aggConfigRaw.MetricsIpPortAddress,
	}
}

func DeployRegistryRollups(t *testing.T, anvils []*utils.AnvilInstance) ([]common.Address, []*registryrollup.ContractSFFLRegistryRollup, []*bind.TransactOpts, []*bind.TransactOpts) {
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

func SetupNearDa(t *testing.T, ctx context.Context, nearEndpointCtn testcontainers.Container, rollupAnvils []*utils.AnvilInstance) []testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	indexerUrl, err := nearEndpointCtn.Endpoint(ctx, "http")
	if err != nil {
		t.Fatalf("Error getting indexer endpoint: %s", err.Error())
	}

	indexerContainerIp, err := nearEndpointCtn.ContainerIP(ctx)
	if err != nil {
		t.Fatalf("Error getting indxer container IP: %s", err.Error())
	}

	hostNearCfgPath := getNearCliConfigPath(t)
	hostNearKeyPath := filepath.Join(hostNearCfgPath, "validator_key.json")
	containerNearCfgPath := "/root/.near"

	time.Sleep(5 * time.Second)

	CopyFileFromContainer(ctx, nearEndpointCtn, filepath.Join(containerNearCfgPath, "validator_key.json"), hostNearKeyPath, 0770)

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

func GetStateRootUpdateAggregation(addr string, rollupID uint32, blockHeight uint64) (*aggtypes.GetStateRootUpdateAggregationResponse, error) {
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

func GetOperatorSetUpdateAggregation(addr string, id uint64) (*aggtypes.GetOperatorSetUpdateAggregationResponse, error) {
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

func CopyFileFromContainer(ctx context.Context, container testcontainers.Container, sourcePath, destinationPath string, destinationPermissions fs.FileMode) error {
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
