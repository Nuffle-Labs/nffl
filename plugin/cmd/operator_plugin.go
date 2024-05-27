package main

import (
	"context"
	"errors"
	"math/big"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/NethermindEth/near-sffl/core/chainio"
	"github.com/NethermindEth/near-sffl/operator"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

type CliOperatorPlugin struct {
	ecdsaKeyPassword string
	ethHttpClient    eth.Client
	clients          *sdkclients.Clients
	avsConfig        optypes.NodeConfig
	avsManager       *operator.AvsManager
	avsReader        *chainio.AvsReader
	avsWriter        *chainio.AvsWriter

	logger logging.Logger
}

func NewOperatorPluginFromCLIContext(ctx *cli.Context) (*CliOperatorPlugin, error) {
	goCtx := context.Background()
	logger, _ := logging.NewZapLogger(logging.Development)

	configPath := ctx.GlobalString(ConfigFileFlag.Name)

	avsConfig := optypes.NodeConfig{}
	err := utils.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		logger.Error("Failed to read config", "err", err)
		return nil, err
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
		return nil, err
	}

	chainID, err := ethHttpClient.ChainID(goCtx)
	if err != nil {
		logger.Error("Failed to get chain ID", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: avsConfig.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainID)
	if err != nil {
		logger.Error("Failed to create signer", "err", err)
		return nil, err
	}

	ecdsaPrivateKey, err := sdkecdsa.ReadKey(avsConfig.EcdsaPrivateKeyStorePath, ecdsaKeyPassword)
	if err != nil {
		logger.Error("Failed to read ecdsa private key", "err", err)
		return nil, err
	}

	clients, err := sdkclients.BuildAll(buildClientConfig, ecdsaPrivateKey, logger)
	if err != nil {
		logger.Error("Failed to create sdk clients", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create avs reader", "err", err)
		return nil, err
	}

	txSender, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, common.HexToAddress(avsConfig.OperatorAddress), logger)
	if err != nil {
		logger.Error("Failed to create tx sender", "err", err)
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(txSender, ethHttpClient, logger, common.HexToAddress(avsConfig.OperatorAddress)).WithGasLimitMultiplier(1.5)
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr,
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create avs writer", "err", err)
		return nil, err
	}

	avsManager, err := operator.NewAvsManager(&avsConfig, clients.EthHttpClient, clients.EthWsClient, clients, txMgr, logger)
	if err != nil {
		logger.Error("Failed to create avs manager", "err", err)
		return nil, err
	}

	return &CliOperatorPlugin{
		ecdsaKeyPassword: ecdsaKeyPassword,
		ethHttpClient:    ethHttpClient,
		clients:          clients,
		avsConfig:        avsConfig,
		avsManager:       avsManager,
		avsReader:        avsReader,
		avsWriter:        avsWriter,
	}, nil
}

func (o *CliOperatorPlugin) OptIn(ctx *cli.Context) error {
	blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

	blsKeypair, err := bls.ReadPrivateKeyFromFile(o.avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		o.logger.Error("Failed to read bls private key", "err", err)
		return err
	}

	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		o.avsConfig.EcdsaPrivateKeyStorePath,
		o.ecdsaKeyPassword,
	)
	if err != nil {
		o.logger.Error("Failed to read operator ecdsa private key", "err", err)
		return err
	}

	err = o.avsManager.RegisterOperatorWithAvs(o.ethHttpClient, operatorEcdsaPrivateKey, blsKeypair)
	if err != nil {
		o.logger.Error("Failed to register operator with avs", "err", err)
		return err
	}

	return nil
}

func (o *CliOperatorPlugin) OptOut(ctx *cli.Context) error {
	blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

	blsKeypair, err := bls.ReadPrivateKeyFromFile(o.avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		o.logger.Error("Failed to read bls private key", "err", err)
		return err
	}

	err = o.avsManager.DeregisterOperator(blsKeypair)
	if err != nil {
		o.logger.Error("Failed to deregister operator", "err", err)
		return err
	}

	return nil
}

func (o *CliOperatorPlugin) Deposit(ctx *cli.Context) error {
	strategy := ctx.GlobalString(StrategyAddrFlag.Name)
	if len(strategy) == 0 {
		o.logger.Error("Strategy address is required for deposit operation")
		return errors.New("strategy address is required for deposit operation")
	}

	strategyAddr := common.HexToAddress(ctx.GlobalString(StrategyAddrFlag.Name))
	_, tokenAddr, err := o.clients.ElChainReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
	if err != nil {
		o.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}

	contractErc20Mock, err := o.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		o.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}

	txOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
	if err != nil {
		o.logger.Error("Failed to get tx opts", "err", err)
		return err
	}

	amount := big.NewInt(1000)
	tx, err := contractErc20Mock.Mint(txOpts, common.HexToAddress(o.avsConfig.OperatorAddress), amount)
	if err != nil {
		o.logger.Error("Failed to assemble Mint tx", "err", err)
		return err
	}

	_, err = o.avsWriter.TxMgr.Send(context.Background(), tx)
	if err != nil {
		o.logger.Error("Failed to submit Mint tx", "err", err)
		return err
	}

	_, err = o.clients.ElChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
	if err != nil {
		o.logger.Error("Failed to deposit into strategy", "err", err)
		return err
	}

	return nil
}
