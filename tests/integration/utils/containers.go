package utils

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/NethermindEth/near-sffl/relayer/config"
)

const (
	NearNetworkId = "localnet"
	NetworkName   = "near-sffl"
	// TODO(edwin): near-sffl-test-relayer -> near-sffl-relayer?
	RelayerImageName     = "near-sffl-test-relayer"
	RelayerContainerName = "relayer"
	IndexerPort          = "3030"
	MetricsPort          = "9091"
)

type AnvilInstance struct {
	Container  testcontainers.Container
	HttpClient eth.Client
	HttpUrl    string
	WsClient   eth.Client
	WsUrl      string
	RpcClient  *rpc.Client
	ChainID    *big.Int
}

func (ai *AnvilInstance) SetBalance(address common.Address, balance *big.Int) error {
	return ai.RpcClient.Call(nil, "anvil_setBalance", address.Hex(), "0x"+balance.Text(16))
}

func (ai *AnvilInstance) Mine(blockCount, timestampInterval *big.Int) error {
	return ai.RpcClient.Call(nil, "anvil_mine", "0x"+blockCount.Text(16), "0x"+timestampInterval.Text(16))
}

func GetDaContractAccountId(anvil *AnvilInstance) string {
	return fmt.Sprintf("da%s.test.near", anvil.ChainID.String())
}

func GetRelayerContainerName(anvil *AnvilInstance) string {
	return fmt.Sprintf("%s%s", RelayerContainerName, anvil.ChainID.String())
}

func getContainerRootKeyPath(keyPath string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(keyPath, usr.HomeDir) {
		return "", fmt.Errorf("key path shall be within: %s", usr.HomeDir)
	}

	return filepath.Join("/root", strings.TrimPrefix(keyPath, usr.HomeDir)), nil
}

func compileContainerConfig(ctx context.Context, daAccountId, keyPath, indexerIp string, anvil *AnvilInstance) (*config.RelayerConfig, error) {
	ports, err := anvil.Container.Ports(ctx)
	if err != nil {
		return nil, err
	}

	var port nat.Port
	// There's only one
	for containerPort, _ := range ports {
		port = containerPort
		break
	}

	containerKeyPath, err := getContainerRootKeyPath(keyPath)
	if err != nil {
		return nil, err
	}

	containerIp, err := anvil.Container.ContainerIP(ctx)
	if err != nil {
		return nil, err
	}

	return &config.RelayerConfig{
		Production:        false,
		DaAccountId:       daAccountId,
		KeyPath:           containerKeyPath,
		RpcUrl:            fmt.Sprintf("ws://%s:%s", containerIp, port.Port()),
		Network:           fmt.Sprintf("http://%s:%s", indexerIp, IndexerPort),
		MetricsIpPortAddr: fmt.Sprintf("%s:%s", GetRelayerContainerName(anvil), MetricsPort),
	}, nil
}

func StartRelayer(t *testing.T, ctx context.Context, daAccountId, indexerContainerIp string, anvil *AnvilInstance) (testcontainers.Container, error) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Couldn't get current user: #%s", err.Error())
	}

	keyFileName := daAccountId + ".json"
	keyPath := filepath.Join(usr.HomeDir, ".near-credentials", NearNetworkId, keyFileName)

	config, err := compileContainerConfig(ctx, daAccountId, keyPath, indexerContainerIp, anvil)
	if err != nil {
		t.Fatalf("Error compiling relayer config: %s", err.Error())
	}

	cmd := config.CompileCMD()
	req := testcontainers.ContainerRequest{
		Image:      RelayerImageName,
		Name:       GetRelayerContainerName(anvil),
		Cmd:        cmd,
		WaitingFor: wait.ForLog("starting relayer"),
		Networks:   []string{NetworkName},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: keyPath,
				},
				Target:   testcontainers.ContainerMountTarget(config.KeyPath),
				ReadOnly: true,
			},
		},
		ExposedPorts: []string{MetricsPort + "/tcp"},
	}

	genericReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	relayerContainer, err := testcontainers.GenericContainer(ctx, genericReq)
	if err != nil {
		return nil, err
	}

	addr, err := relayerContainer.Endpoint(ctx, "")
	if err != nil {
		return nil, err
	}

	t.Log("Relayer metrics endpoint:", addr)
	return relayerContainer, nil
}

func StartAnvilTestContainer(t *testing.T, ctx context.Context, name, exposedPort, chainId string, isMainnet bool, networkName string) *AnvilInstance {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest@sha256:8b843eb65cc7b155303b316f65d27173c862b37719dc095ef3a2ef27ce8d3c00",
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

	anvil := &AnvilInstance{
		Container:  anvilC,
		HttpClient: httpClient,
		HttpUrl:    httpUrl,
		WsClient:   wsClient,
		WsUrl:      wsUrl,
		RpcClient:  rpcClient,
		ChainID:    fetchedChainId,
	}

	if isMainnet {
		err := anvil.Mine(big.NewInt(100), big.NewInt(1))
		if err != nil {
			t.Fatalf("Anvil failed to Mine: %s", err.Error())
		}
	}

	return anvil
}
