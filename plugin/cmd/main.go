package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
		EnvVar:   "CONFIG",
	}
	EcdsaKeyPasswordFlag = cli.StringFlag{
		Name:     "ecdsa-key-password",
		Required: false,
		Usage:    "Password to decrypt the ecdsa key",
		EnvVar:   "ECDSA_KEY_PASSWORD",
	}
	BlsKeyPasswordFlag = cli.StringFlag{
		Name:     "bls-key-password",
		Required: false,
		Usage:    "Password to decrypt the bls key",
		EnvVar:   "BLS_KEY_PASSWORD",
	}
	OperationFlag = cli.StringFlag{
		Name:     "operation-type",
		Required: true,
		Usage:    "Supported operations: opt-in, deposit",
		EnvVar:   "OPERATION_TYPE",
	}
	StrategyAddrFlag = cli.StringFlag{
		Name:     "strategy-addr",
		Required: false,
		Usage:    "Strategy address for deposit mock tokens, only used for deposit action",
		EnvVar:   "STRATEGY_ADDR",
	}
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		ConfigFileFlag,
		EcdsaKeyPasswordFlag,
		BlsKeyPasswordFlag,
		OperationFlag,
		StrategyAddrFlag,
	}
	app.Name = "sffl-plugin"
	app.Usage = "SFFL Plugin"
	app.Description = "This is used to run one time operations like avs opt-in/opt-out"
	app.Action = func(ctx *cli.Context) error {
		operatorPlugin, err := NewOperatorPluginFromCLIContext(ctx)
		if err != nil {
			return err
		}
		operationType := ctx.GlobalString(OperationFlag.Name)
		switch operationType {
		case "opt-in":
			return operatorPlugin.OptIn()
		case "opt-out":
			return operatorPlugin.OptOut()
		case "deposit":
			return operatorPlugin.Deposit()
		default:
			return cli.NewExitError(fmt.Sprintf("Invalid operation type: %v", operationType), 1)
		}
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}
