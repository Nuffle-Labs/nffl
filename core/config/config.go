package config

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
)

// Config contains all of the configuration information for SFFL aggregators and challengers.
// Operators use a separate config. (see config-files/operator.anvil.yaml)
type Config struct {
	EcdsaPrivateKey *ecdsa.PrivateKey `json:"-"`
	BlsPrivateKey   *bls.PrivateKey   `json:"-"`
	// we need the url for the eigensdk currently... eventually standardize api so as to
	// only take an ethclient or an rpcUrl (and build the ethclient at each constructor site)
	EthHttpRpcUrl                  string                `json:"ethHttpRpcUrl"`
	EthWsRpcUrl                    string                `json:"ethWsRpcUrl"`
	RollupsInfo                    map[uint32]RollupInfo `json:"rollupsInfo"`
	OperatorStateRetrieverAddr     common.Address        `json:"operatorStateRetrieverAddr"`
	SFFLRegistryCoordinatorAddr    common.Address        `json:"sfflRegistryCoordinatorAddr"`
	AggregatorServerIpPortAddr     string                `json:"aggregatorServerIpPortAddr"`
	AggregatorRestServerIpPortAddr string                `json:"aggregatorRestServerIpPortAddr"`
	AggregatorDatabasePath         string                `json:"aggregatorDatabasePath"`
	AggregatorCheckpointInterval   time.Duration         `json:"aggregatorCheckpointInterval"`
	RegisterOperatorOnStartup      bool                  `json:"registerOperatorOnStartup"`
	AggregatorAddress              common.Address        `json:"aggregatorAddress"`

	// metrics related
	EnableMetrics        bool   `json:"enableMetrics"`
	MetricsIpPortAddress string `json:"metricsIpPortAddress"`
}

// These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                    sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                      string              `yaml:"eth_rpc_url"`
	EthWsUrl                       string              `yaml:"eth_ws_url"`
	AggregatorServerIpPortAddr     string              `yaml:"aggregator_server_ip_port_address"`
	AggregatorRestServerIpPortAddr string              `yaml:"aggregator_rest_server_ip_port_address"`
	AggregatorDatabasePath         string              `yaml:"aggregator_database_path"`
	AggregatorCheckpointInterval   uint32              `yaml:"aggregator_checkpoint_interval"`
	RegisterOperatorOnStartup      bool                `yaml:"register_operator_on_startup"`
	RollupIdsToRpcUrls             map[uint32]string   `yaml:"rollup_ids_to_rpc_urls"`
	RollupIdsToRegistryAddresses   map[uint32]string   `yaml:"rollup_ids_to_registry_addresses"`

	EnableMetrics        bool   `yaml:"enable_metrics"`
	MetricsIpPortAddress string `yaml:"metrics_ip_port_address"`
}

// These are read from SFFLDeploymentFileFlag
type SFFLDeploymentRaw struct {
	Addresses SFFLContractsRaw `json:"addresses"`
}
type SFFLContractsRaw struct {
	RegistryCoordinatorAddr    string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr string `json:"operatorStateRetriever"`
}

type RollupInfo struct {
	SFFLRegistryRollupAddr common.Address
	RpcUrl                 string
}

func NewConfigRaw(ctx *cli.Context) (*ConfigRaw, error) {
	var configRaw ConfigRaw
	configFilePath := ctx.GlobalString(ConfigFileFlag.Name)
	if configFilePath == "" {
		return nil, errors.New("config file path is empty")
	}

	err := sdkutils.ReadYamlConfig(configFilePath, &configRaw)
	if err != nil {
		return nil, err
	}

	return &configRaw, nil
}

func CompileRollupsInfo(configRaw *ConfigRaw) map[uint32]RollupInfo {
	if len(configRaw.RollupIdsToRpcUrls) != len(configRaw.RollupIdsToRegistryAddresses) {
		panic("RollupIdsToRpcUrls and RollupIdsToRegistryAddresses must have the same length")
	}

	rollupsInfo := make(map[uint32]RollupInfo)
	for chainId, registryAddress := range configRaw.RollupIdsToRegistryAddresses {
		url, exist := configRaw.RollupIdsToRpcUrls[chainId]
		if !exist {
			panic(fmt.Sprintf("RPC URL doesn't exist for chainId %d", chainId))
		}

		rollupsInfo[chainId] = RollupInfo{
			RpcUrl:                 url,
			SFFLRegistryRollupAddr: common.HexToAddress(registryAddress),
		}
	}

	return rollupsInfo
}

// NewConfig parses config file to read from from flags or environment variables
// Note: This config is shared by challenger and aggregator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context, configRaw ConfigRaw, logger sdklogging.Logger) (*Config, error) {
	rollupsInfo := CompileRollupsInfo(&configRaw)

	var sfflDeploymentRaw SFFLDeploymentRaw
	sfflDeploymentFilePath := ctx.GlobalString(SFFLDeploymentFileFlag.Name)
	if _, err := os.Stat(sfflDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + sfflDeploymentFilePath + " does not exist")
	}
	err := sdkutils.ReadJsonConfig(sfflDeploymentFilePath, &sfflDeploymentRaw)
	if err != nil {
		logger.Error("Cannot read JSON config", "err", err)
		return nil, err
	}

	ecdsaPrivateKeyString := ctx.GlobalString(EcdsaPrivateKeyFlag.Name)
	if ecdsaPrivateKeyString[:2] == "0x" {
		ecdsaPrivateKeyString = ecdsaPrivateKeyString[2:]
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivateKeyString)
	if err != nil {
		logger.Error("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}

	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		logger.Error("Cannot get operator address", "err", err)
		return nil, err
	}

	config := &Config{
		EcdsaPrivateKey:                ecdsaPrivateKey,
		EthWsRpcUrl:                    configRaw.EthWsUrl,
		EthHttpRpcUrl:                  configRaw.EthRpcUrl,
		OperatorStateRetrieverAddr:     common.HexToAddress(sfflDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		SFFLRegistryCoordinatorAddr:    common.HexToAddress(sfflDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr:     configRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:      configRaw.RegisterOperatorOnStartup,
		AggregatorRestServerIpPortAddr: configRaw.AggregatorRestServerIpPortAddr,
		AggregatorDatabasePath:         configRaw.AggregatorDatabasePath,
		AggregatorCheckpointInterval:   time.Duration(configRaw.AggregatorCheckpointInterval) * time.Millisecond,
		AggregatorAddress:              aggregatorAddr,
		RollupsInfo:                    rollupsInfo,
		EnableMetrics:                  configRaw.EnableMetrics,
		MetricsIpPortAddress:           configRaw.MetricsIpPortAddress,
	}
	config.validate()

	return config, nil
}

func (c *Config) validate() {
	if c.OperatorStateRetrieverAddr == common.HexToAddress("") {
		panic("Config: BLSOperatorStateRetrieverAddr is required")
	}

	if c.SFFLRegistryCoordinatorAddr == common.HexToAddress("") {
		panic("Config: SFFLRegistryCoordinatorAddr is required")
	}
	if c.MetricsIpPortAddress == "" {
		panic("Config: MetricsIpPortAddress shall be valid socket addr even if disabled")
	}
}

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	SFFLDeploymentFileFlag = cli.StringFlag{
		Name:     "nffl-deployment",
		Required: true,
		Usage:    "Load SFFL contract addresses from `FILE`",
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "Ethereum private key",
		Required: true,
		EnvVar:   "ECDSA_PRIVATE_KEY",
	}
)

var requiredFlags = []cli.Flag{
	ConfigFileFlag,
	SFFLDeploymentFileFlag,
	EcdsaPrivateKeyFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
