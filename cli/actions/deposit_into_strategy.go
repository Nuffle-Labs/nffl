package actions

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/operator"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

func DepositIntoStrategy(ctx *cli.Context) error {
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig, err := readNodeConfig(configPath)
	if err != nil {
		return err
	}

	operator, err := operator.NewOperatorFromConfig(*nodeConfig)
	if err != nil {
		return err
	}

	strategyAddrStr := ctx.String("strategy-addr")
	strategyAddr := common.HexToAddress(strategyAddrStr)
	amountStr := ctx.String("amount")
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		fmt.Println("Error converting amount to big.Int")
		return err
	}

	err = operator.DepositIntoStrategy(strategyAddr, amount)
	if err != nil {
		return err
	}

	return nil
}
