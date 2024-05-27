package integration

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/NethermindEth/near-sffl/core/chainio"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
	"github.com/NethermindEth/near-sffl/tests/integration/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRegistration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)

	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

	networkName := "near-sffl-registration"
	net, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
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

	mainnetAnvilContainerName := "mainnet-anvil"

	mainnetAnvil := utils.StartAnvilTestContainer(t, ctx, mainnetAnvilContainerName, "8545", "1", true, networkName)

	nodeConfig := optypes.NodeConfig{}
	nodeConfigFilePath := "../../config-files/plugin.anvil.yaml"
	err = sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	operatorAddr := common.HexToAddress(nodeConfig.OperatorAddress)

	avsReader, err := chainio.BuildAvsReader(common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress), common.HexToAddress(nodeConfig.OperatorStateRetrieverAddress), mainnetAnvil.HttpClient, logger)
	if err != nil {
		t.Fatalf("Error building avs reader: %s", err.Error())
	}

	blsPassword, err := os.ReadFile("../../tests/keys/bls/1/password.txt")
	if err != nil {
		t.Fatalf("Error reading bls password: %s", err.Error())
	}

	ecdsaPassword, err := os.ReadFile("../../tests/keys/ecdsa/1/password.txt")
	if err != nil {
		t.Fatalf("Error reading ecdsa password: %s", err.Error())
	}

	ecdsaPrivateKey, err := sdkecdsa.ReadKey("../../tests/keys/ecdsa/1/key.json", string(ecdsaPassword))
	if err != nil {
		t.Fatalf("Error reading ecdsa private key: %s", err.Error())
	}

	buildClientConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 mainnetAnvil.HttpUrl,
		EthWsUrl:                   mainnetAnvil.WsUrl,
		RegistryCoordinatorAddr:    nodeConfig.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: nodeConfig.OperatorStateRetrieverAddress,
		AvsName:                    "super-fast-finality-layer",
		PromMetricsIpPortAddress:   "127.0.0.1:9090",
	}

	clients, err := sdkclients.BuildAll(buildClientConfig, ecdsaPrivateKey, logger)
	if err != nil {
		t.Fatalf("Error building clients: %s", err.Error())
	}

	_, err = clients.ElChainWriter.RegisterAsOperator(ctx, types.Operator{Address: operatorAddr.String(), EarningsReceiverAddress: operatorAddr.String()})
	if err != nil {
		t.Fatalf("Error registering operator: %s", err.Error())
	}

	operatorPluginContainerOptIn := runOperatorPluginContainer(t, ctx, "plugin-opt-in", networkName, "opt-in", string(ecdsaPassword), string(blsPassword))

	isOperatorRegistered, err := avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(nodeConfig.OperatorAddress))
	if err != nil {
		t.Fatalf("Error checking if operator is registered: %s", err.Error())
	}

	if !isOperatorRegistered {
		t.Fatal("Operator should be registered after opting in")
	}

	operatorPluginContainerOptOut := runOperatorPluginContainer(t, ctx, "plugin-opt-out", networkName, "opt-out", string(ecdsaPassword), string(blsPassword))

	isOperatorRegistered, err = avsReader.IsOperatorRegistered(&bind.CallOpts{}, common.HexToAddress(nodeConfig.OperatorAddress))
	if err != nil {
		t.Fatalf("Error checking if operator is registered: %s", err.Error())
	}

	if isOperatorRegistered {
		t.Fatal("Operator should not be registered after opting out")
	}

	defer func() {
		if err := mainnetAnvil.Container.Terminate(ctx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		if err := operatorPluginContainerOptIn.Terminate(ctx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}
		if err := operatorPluginContainerOptOut.Terminate(ctx); err != nil {
			t.Fatalf("Error terminating container: %s", err.Error())
		}

		if err := net.Remove(ctx); err != nil {
			t.Fatalf("Error removing network: %s", err.Error())
		}

		cancel()
	}()

	t.Log("Done")
	<-ctx.Done()
}

func runOperatorPluginContainer(t *testing.T, ctx context.Context, name, networkName, operation, ecdsaPassword, blsPassword string) testcontainers.Container {
	keysPath, err := filepath.Abs("../../tests/keys/")
	if err != nil {
		t.Fatalf("Error getting absolute path: %s", err.Error())
	}

	req := testcontainers.ContainerRequest{
		Image:    "near-sffl-operator-plugin",
		Name:     name,
		Networks: []string{networkName},
		Env: map[string]string{
			"ECDSA_KEY_PASSWORD": ecdsaPassword,
			"BLS_KEY_PASSWORD":   blsPassword,
		},
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      "../../config-files/plugin.anvil.yaml",
				ContainerFilePath: "/near-sffl/config.yml",
			},
		},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: keysPath,
				},
				Target:   testcontainers.ContainerMountTarget("/near-sffl/keys/"),
				ReadOnly: true,
			},
		},
		Cmd: []string{
			"--config",
			"/near-sffl/config.yml",
			"--operation-type",
			operation,
		},
		WaitingFor: wait.ForExit(),
	}

	genericReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	container, err := testcontainers.GenericContainer(ctx, genericReq)
	if err != nil {
		t.Fatalf("Error starting operator plugin container: %s", err.Error())
	}

	return container
}
