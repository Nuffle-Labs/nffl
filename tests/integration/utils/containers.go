package utils

import (
	"context"
	"github.com/NethermindEth/near-sffl/relayer"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	NetworkName = "near-sffl"
	// TODO(edwin): near-sffl-test-relayer -> near-sffl-relayer?
	RelayerImageName     = "near-sffl-test-relayer"
	RelayerContainerName = "relayer"
)

func StartRelayer(ctx context.Context, config relayer.RelayerConfig) (testcontainers.Container, error) {
	cmd, err := config.CompileContainerCmd()
	if err != nil {
		return nil, err
	}

	target, err := config.GetContainerRootKeyPath()
	if err != nil {
		return nil, err
	}

	req := testcontainers.ContainerRequest{
		Image:        RelayerImageName,
		Name:         RelayerContainerName,
		Cmd:          cmd,
		ExposedPorts: []string{"3030/tcp"},
		WaitingFor:   wait.ForLog("starting relayer"),
		Networks:     []string{NetworkName},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: config.GetHostKeyPath(),
					//Type:   mount.TypeVolume,
					//Source: "near_cli_data",
					//Target: "/near-cli",
				},
				Target:   testcontainers.ContainerMountTarget(target),
				ReadOnly: true,
			},
		},
	}

	genericReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	indexerContainer, err := testcontainers.GenericContainer(ctx, genericReq)
	if err != nil {
		return nil, err
	}

	return indexerContainer, nil
}
