package config

import (
	"crypto/ecdsa"
	"errors"
	"os"

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
	EcdsaPrivateKey           *ecdsa.PrivateKey
	BlsPrivateKey             *bls.PrivateKey
	EigenMetricsIpPortAddress string
	// we need the url for the eigensdk currently... eventually standardize api so as to
	// only take an ethclient or an rpcUrl (and build the ethclient at each constructor site)
	EthHttpRpcUrl                  string
	EthWsRpcUrl                    string
	OperatorStateRetrieverAddr     common.Address
	SFFLRegistryCoordinatorAddr    common.Address
	AggregatorServerIpPortAddr     string
	AggregatorRestServerIpPortAddr string
	AggregatorDatabasePath         string
	RegisterOperatorOnStartup      bool
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement marshalJson
	AggregatorAddress common.Address
}

// These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                    sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                      string              `yaml:"eth_rpc_url"`
	EthWsUrl                       string              `yaml:"eth_ws_url"`
	AggregatorServerIpPortAddr     string              `yaml:"aggregator_server_ip_port_address"`
	AggregatorRestServerIpPortAddr string              `yaml:"aggregator_rest_server_ip_port_address"`
	AggregatorDatabasePath         string              `yaml:"aggregator_database_path"`
	RegisterOperatorOnStartup      bool                `yaml:"register_operator_on_startup"`
}

// These are read from SFFLDeploymentFileFlag
type SFFLDeploymentRaw struct {
	Addresses SFFLContractsRaw `json:"addresses"`
}
type SFFLContractsRaw struct {
	RegistryCoordinatorAddr    string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr string `json:"operatorStateRetriever"`
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

// NewConfig parses config file to read from from flags or environment variables
// Note: This config is shared by challenger and aggregator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context, configRaw ConfigRaw, logger sdklogging.Logger) (*Config, error) {
	var sfflDeploymentRaw SFFLDeploymentRaw
	sfflDeploymentFilePath := ctx.GlobalString(SFFLDeploymentFileFlag.Name)
	if _, err := os.Stat(sfflDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + sfflDeploymentFilePath + " does not exist")
	}
	sdkutils.ReadJsonConfig(sfflDeploymentFilePath, &sfflDeploymentRaw)

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
		AggregatorAddress:              aggregatorAddr,
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
}

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	SFFLDeploymentFileFlag = cli.StringFlag{
		Name:     "sffl-deployment",
		Required: true,
		Usage:    "Load SFFL contract addresses from `FILE`",
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "Ethereum private key",
		Required: true,
		EnvVar:   "ECDSA_PRIVATE_KEY",
	}
	/* Optional Flags */
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
