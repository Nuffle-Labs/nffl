package actions

import (
	"github.com/urfave/cli"

	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
)

func RegisterOperatorWithEigenlayer(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig, err := readNodeConfig(configPath)
	if err != nil {
		return err
	}

	operator, err := operator.NewOperatorFromConfig(*nodeConfig)
	if err != nil {
		return err
	}

	err = operator.RegisterOperatorWithEigenlayer()
	if err != nil {
		return err
	}

	return nil
}
