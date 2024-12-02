package integration

import (
	"context"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	registryrollup "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLRegistryRollup"
	"github.com/Nuffle-Labs/nffl/core/chainio"
	"github.com/Nuffle-Labs/nffl/core/types"
	"github.com/Nuffle-Labs/nffl/tests/integration/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func TestIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Second)
	setup := setupTestEnv(t, ctx)
	t.Cleanup(func() {
		cancel()

		setup.Cleanup()
	})

	time.Sleep(55 * time.Second)

	taskHash, err := setup.AvsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskHashes(&bind.CallOpts{}, 0)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	taskResponseHash, err := setup.AvsReader.AvsServiceBindings.TaskManager.AllCheckpointTaskResponses(&bind.CallOpts{}, 0)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
	}

	stateRootHeight, err := setup.RollupAnvils[1].HttpClient.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Cannot get current block height: %s", err.Error())
	}

	stateRootUpdate, err := GetStateRootUpdateAggregation(setup.AggregatorRestUrl, uint32(setup.RollupAnvils[0].ChainID.Uint64()), stateRootHeight-1)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}
	_, err = setup.RegistryRollups[1].UpdateStateRoot(setup.RegistryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	newOperatorConfig, _, _ := GenOperatorConfig(t, ctx, "4", setup.MainnetAnvil, setup.RollupAnvils, setup.RabbitMq)
	newOperator := StartOperator(t, ctx, newOperatorConfig)

	time.Sleep(60 * time.Second)

	// Check if operator set was updated on rollups
	for _, registryRollup := range setup.RegistryRollups {
		nextOperatorSetUpdateId, err := registryRollup.NextOperatorUpdateId(&bind.CallOpts{})
		if err != nil {
			t.Fatalf("Error getting next operator set update ID: %s", err.Error())
		}

		if nextOperatorSetUpdateId != 2 {
			t.Fatalf("Wrong next operator set update ID: expected %d, got %d", 2, nextOperatorSetUpdateId)
		}
	}

	stateRootHeight = uint64(16)
	stateRootUpdate, err = GetStateRootUpdateAggregation(setup.AggregatorRestUrl, uint32(setup.RollupAnvils[0].ChainID.Uint64()), stateRootHeight)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}
	_, err = setup.RegistryRollups[1].UpdateStateRoot(setup.RegistryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	operatorSetUpdateCount, err := setup.AvsReader.AvsServiceBindings.OperatorSetUpdateRegistry.GetOperatorSetUpdateCount(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("Error getting operator set update count: %s", err.Error())
	}
	if operatorSetUpdateCount != 2 {
		t.Fatalf("Wrong operator set update count")
	}

	stateRootHeight, err = setup.RollupAnvils[1].HttpClient.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Cannot get current block height: %s", err.Error())
	}

	stateRootUpdate, err = GetStateRootUpdateAggregation(setup.AggregatorRestUrl, uint32(setup.RollupAnvils[0].ChainID.Uint64()), stateRootHeight-1)
	if err != nil {
		t.Fatalf("Cannot get state root update: %s", err.Error())
	}

	// Check if operator sets are same on rollups
	_, err = setup.RegistryRollups[1].UpdateStateRoot(setup.RegistryRollupAuths[1], registryrollup.StateRootUpdateMessage(stateRootUpdate.Message.ToBinding()), stateRootUpdate.Aggregation.ExtractBindingRollup())
	if err != nil {
		t.Fatalf("Error updating state root: %s", err.Error())
	}

	operatorSetUpdate, err := GetOperatorSetUpdateAggregation(setup.AggregatorRestUrl, operatorSetUpdateCount-1)
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

func setupTestEnv(t *testing.T, ctx context.Context) *TestEnv {
	containersCtx, cancelContainersCtx := context.WithCancel(context.Background())

	networkName := "nffl"
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

	mainnetAnvil := utils.StartAnvilTestContainer(t, containersCtx, mainnetAnvilContainerName, "8545", "1", true, networkName)
	rollupAnvils := []*utils.AnvilInstance{
		utils.StartAnvilTestContainer(t, containersCtx, rollup0AnvilContainerName, "8546", "2", false, networkName),
		utils.StartAnvilTestContainer(t, containersCtx, rollup1AnvilContainerName, "8547", "3", false, networkName),
	}
	rabbitMq := StartRabbitMqContainer(t, containersCtx, rmqContainerName, networkName)
	indexerContainer, relayers := startIndexer(t, containersCtx, indexerContainerName, rollupAnvils, rabbitMq, networkName)

	sfflDeploymentRaw := ReadSfflDeploymentRaw()

	configRaw := BuildConfigRaw(mainnetAnvil, rollupAnvils)
	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

	nodeConfig, _, _ := GenOperatorConfig(t, ctx, "3", mainnetAnvil, rollupAnvils, rabbitMq)

	addresses, registryRollups, registryRollupAuths, _ := DeployRegistryRollups(t, rollupAnvils)
	operator := StartOperator(t, ctx, nodeConfig)

	config := BuildConfig(t, sfflDeploymentRaw, addresses, rollupAnvils, configRaw)
	aggregator := StartAggregator(t, ctx, config, logger)

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

	return &TestEnv{
		MainnetAnvil:        mainnetAnvil,
		RollupAnvils:        rollupAnvils,
		RabbitMq:            rabbitMq,
		IndexerContainer:    indexerContainer,
		Operator:            operator,
		Aggregator:          aggregator,
		AggregatorRestUrl:   "http://" + config.AggregatorRestServerIpPortAddr,
		AvsReader:           avsReader,
		RegistryRollups:     registryRollups,
		RegistryRollupAuths: registryRollupAuths,
		Cleanup:             cleanup,
	}
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
		Image:        "nffl-indexer",
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

	relayers := SetupNearDa(t, ctx, indexerContainer, rollupAnvils)
	return indexerContainer, relayers
}
