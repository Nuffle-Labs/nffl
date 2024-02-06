package actions

import (
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/urfave/cli"
)

func PrintOperatorStatus(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig, err := readNodeConfig(configPath)
	if err != nil {
		return err
	}

	operator, err := operator.NewOperatorFromConfig(*nodeConfig)
	if err != nil {
		return err
	}

	err = operator.PrintOperatorStatus()
	if err != nil {
		return err
	}

	return nil
}
