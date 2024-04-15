package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/NethermindEth/near-sffl/core/chainio"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	app.Action = plugin
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func plugin(ctx *cli.Context) {
	goCtx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)

	operationType := ctx.GlobalString(OperationFlag.Name)
	configPath := ctx.GlobalString(ConfigFileFlag.Name)

	avsConfig := optypes.NodeConfig{}
	err := utils.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		logger.Error("Failed to read config", "err", err)
		return
	}

	logger.Info("Starting with config", "avsConfig", avsConfig)

	ecdsaKeyPassword := ctx.GlobalString(EcdsaKeyPasswordFlag.Name)

	buildClientConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 avsConfig.EthRpcUrl,
		EthWsUrl:                   avsConfig.EthWsUrl,
		RegistryCoordinatorAddr:    avsConfig.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: avsConfig.OperatorStateRetrieverAddress,
		AvsName:                    "super-fast-finality-layer",
		PromMetricsIpPortAddress:   avsConfig.EigenMetricsIpPortAddress,
	}
	ethHttpClient, err := eth.NewClient(avsConfig.EthRpcUrl)
	if err != nil {
		logger.Error("Failed to connect to eth client", "err", err)
		return
	}
	chainID, err := ethHttpClient.ChainID(goCtx)
	if err != nil {
		logger.Error("Failed to get chain ID", "err", err)
		return
	}
	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: avsConfig.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainID)
	if err != nil {
		logger.Error("Failed to create signer", "err", err)
		return
	}
	ecdsaPrivateKey, err := sdkecdsa.ReadKey(avsConfig.EcdsaPrivateKeyStorePath, ecdsaKeyPassword)
	if err != nil {
		logger.Error("Failed to read ecdsa private key", "err", err)
		return
	}
	clients, err := sdkclients.BuildAll(buildClientConfig, ecdsaPrivateKey, logger)
	if err != nil {
		logger.Error("Failed to create sdk clients", "err", err)
		return
	}
	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create avs reader", "err", err)
		return
	}
	txSender, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, common.HexToAddress(avsConfig.OperatorAddress), logger)
	if err != nil {
		logger.Error("Failed to create tx sender", "err", err)
		return
	}
	txMgr := txmgr.NewSimpleTxManager(txSender, ethHttpClient, logger, signerV2, common.HexToAddress(avsConfig.OperatorAddress))
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr,
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create avs writer", "err", err)
		return
	}

	if operationType == "opt-in" {
		blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

		blsKeypair, err := bls.ReadPrivateKeyFromFile(avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
		if err != nil {
			logger.Error("Failed to read bls private key", "err", err)
			return
		}

		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			avsConfig.EcdsaPrivateKeyStorePath,
			ecdsaKeyPassword,
		)
		if err != nil {
			logger.Error("Failed to read operator ecdsa private key", "err", err)
			return
		}

		// Register with registry coordination
		quorumNumbers := []byte{0}
		socket := "Not Needed"
		sigValidForSeconds := int64(1_000_000)
		operatorToAvsRegistrationSigSalt := [32]byte{123}
		operatorToAvsRegistrationSigExpiry := big.NewInt(int64(time.Now().Unix()) + sigValidForSeconds)
		logger.Infof("Registering with registry coordination with quorum numbers %v and socket %s", quorumNumbers, socket)
		r, err := clients.AvsRegistryChainWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
			goCtx,
			operatorEcdsaPrivateKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
			blsKeypair, quorumNumbers, socket,
		)
		if err != nil {
			logger.Error("Failed to assemble RegisterOperatorWithAVSRegistryCoordinator tx", "err", err)
			return
		}
		logger.Infof("Registered with registry coordination successfully with tx hash %s", r.TxHash.Hex())
	} else if operationType == "opt-out" {
		fmt.Println("Opting out of slashing - unimplemented")
	} else if operationType == "deposit" {
		starategyAddrString := ctx.GlobalString(StrategyAddrFlag.Name)
		if len(starategyAddrString) == 0 {
			logger.Error("Strategy address is required for deposit operation")
			return
		}
		strategyAddr := common.HexToAddress(ctx.GlobalString(StrategyAddrFlag.Name))
		_, tokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
		if err != nil {
			logger.Error("Failed to fetch strategy contract", "err", err)
			return
		}
		contractErc20Mock, err := avsReader.GetErc20Mock(context.Background(), tokenAddr)
		if err != nil {
			logger.Error("Failed to fetch ERC20Mock contract", "err", err)
			return
		}
		txOpts, err := avsWriter.TxMgr.GetNoSendTxOpts()
		if err != nil {
			logger.Error("Failed to get tx opts", "err", err)
			return
		}
		amount := big.NewInt(1000)
		tx, err := contractErc20Mock.Mint(txOpts, common.HexToAddress(avsConfig.OperatorAddress), amount)
		if err != nil {
			logger.Error("Failed to assemble Mint tx", "err", err)
			return
		}
		_, err = avsWriter.TxMgr.Send(context.Background(), tx)
		if err != nil {
			logger.Error("Failed to submit Mint tx", "err", err)
			return
		}

		_, err = clients.ElChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
		if err != nil {
			logger.Error("Failed to deposit into strategy", "err", err)
			return
		}
		return
	} else {
		logger.Error("Invalid operation type")
	}
}

func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
	return regcoord.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}
