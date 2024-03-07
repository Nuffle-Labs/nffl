// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import {EmptyContract} from "@eigenlayer/test/mocks/EmptyContract.sol";

import {
    IBLSApkRegistry,
    IIndexRegistry,
    IStakeRegistry,
    IRegistryCoordinator
} from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "eigenlayer-middleware/src/StakeRegistry.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {RegistryCoordinator} from "../src/external/RegistryCoordinator.sol";
import {SFFLServiceManager} from "../src/eth/SFFLServiceManager.sol";
import {SFFLTaskManager} from "../src/eth/SFFLTaskManager.sol";
import {SFFLRegistryCoordinator} from "../src/eth/SFFLRegistryCoordinator.sol";
import {SFFLOperatorSetUpdateRegistry} from "../src/eth/SFFLOperatorSetUpdateRegistry.sol";
import {ERC20Mock, IERC20} from "../test/mock/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// forge script script/SFFLDeployer.s.sol:SFFLDeployer --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
contract SFFLDeployer is Script, Utils {
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;

    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

    // SFFL contracts
    ProxyAdmin public sfflProxyAdmin;
    PauserRegistry public sfflPauserReg;

    SFFLRegistryCoordinator public registryCoordinator;
    IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    SFFLOperatorSetUpdateRegistry public operatorSetUpdateRegistry;
    SFFLOperatorSetUpdateRegistry public operatorSetUpdateRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    SFFLServiceManager public sfflServiceManager;
    SFFLServiceManager public sfflServiceManagerImplementation;

    SFFLTaskManager public sfflTaskManager;
    SFFLTaskManager public sfflTaskManagerImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        IStrategyManager strategyManager =
            IStrategyManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategyManager"));
        IDelegationManager delegationManager =
            IDelegationManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegation"));
        ProxyAdmin eigenLayerProxyAdmin =
            ProxyAdmin(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerProxyAdmin"));
        PauserRegistry eigenLayerPauserReg =
            PauserRegistry(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerPauserReg"));
        StrategyBaseTVLLimits baseStrategyImplementation = StrategyBaseTVLLimits(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.baseStrategyImplementation")
        );

        address sfflCommunityMultisig = msg.sender;
        address sfflPauser = msg.sender;

        vm.startBroadcast();
        _deployErc20AndStrategyAndWhitelistStrategy(
            eigenLayerProxyAdmin, eigenLayerPauserReg, baseStrategyImplementation, strategyManager
        );
        _deploySFFLContracts(delegationManager, erc20MockStrategy, sfflCommunityMultisig, sfflPauser);
        vm.stopBroadcast();
    }

    function _deployErc20AndStrategyAndWhitelistStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImplementation,
        IStrategyManager strategyManager
    ) internal {
        erc20Mock = new ERC20Mock();

        erc20MockStrategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector,
                        1 ether, // maxPerDeposit
                        100 ether, // maxDeposits
                        IERC20(erc20Mock),
                        eigenLayerPauserReg
                    )
                )
            )
        );
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;
        strategyManager.addStrategiesToDepositWhitelist(strats);
    }

    function _deploySFFLContracts(
        IDelegationManager delegationManager,
        IStrategy strat,
        address sfflCommunityMultisig,
        address sfflPauser
    ) internal {
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint256 numStrategies = deployedStrategyArray.length;

        sfflProxyAdmin = new ProxyAdmin();

        {
            address[] memory pausers = new address[](2);
            pausers[0] = sfflPauser;
            pausers[1] = sfflCommunityMultisig;
            sfflPauserReg = new PauserRegistry(pausers, sfflCommunityMultisig);
        }

        EmptyContract emptyContract = new EmptyContract();

        sfflServiceManager = SFFLServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        sfflTaskManager = SFFLTaskManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        registryCoordinator = SFFLRegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        blsApkRegistry = IBLSApkRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );
        operatorSetUpdateRegistry = SFFLOperatorSetUpdateRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(sfflProxyAdmin), ""))
        );

        operatorStateRetriever = new OperatorStateRetriever();

        {
            stakeRegistryImplementation = new StakeRegistry(registryCoordinator, delegationManager);

            sfflProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))), address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinator);

            sfflProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))), address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(registryCoordinator);

            sfflProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))), address(indexRegistryImplementation)
            );

            operatorSetUpdateRegistryImplementation = new SFFLOperatorSetUpdateRegistry(registryCoordinator);

            sfflProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(operatorSetUpdateRegistry))),
                address(operatorSetUpdateRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new SFFLRegistryCoordinator(
            sfflServiceManager,
            IStakeRegistry(address(stakeRegistry)),
            IBLSApkRegistry(address(blsApkRegistry)),
            IIndexRegistry(address(indexRegistry)),
            SFFLOperatorSetUpdateRegistry(address(operatorSetUpdateRegistry))
        );

        {
            uint256 numQuorums = 1;
            IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams =
                new IRegistryCoordinator.OperatorSetParam[](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                quorumsOperatorSetParams[i] = IRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][] memory quorumsStrategyParams =
                new IStakeRegistry.StrategyParams[][](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);
                for (uint256 j = 0; j < numStrategies; j++) {
                    quorumsStrategyParams[i][j] = IStakeRegistry.StrategyParams({
                        strategy: deployedStrategyArray[j],
                        // setting this to 1 ether since the divisor is also 1 ether
                        // therefore this allows an operator to register with even just 1 token
                        // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
                        //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                        multiplier: 1 ether
                    });
                }
            }
            sfflProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    RegistryCoordinator.initialize.selector,
                    sfflCommunityMultisig,
                    sfflCommunityMultisig,
                    sfflCommunityMultisig,
                    sfflPauserReg,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        sfflServiceManagerImplementation =
            new SFFLServiceManager(delegationManager, registryCoordinator, stakeRegistry, sfflTaskManager);

        sfflProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(sfflServiceManager))),
            address(sfflServiceManagerImplementation),
            abi.encodeWithSelector(sfflServiceManager.initialize.selector, sfflCommunityMultisig)
        );

        sfflTaskManagerImplementation = new SFFLTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK);

        sfflProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(sfflTaskManager))),
            address(sfflTaskManagerImplementation),
            abi.encodeWithSelector(
                sfflTaskManager.initialize.selector,
                sfflPauserReg,
                sfflCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR
            )
        );

        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "erc20Mock", address(erc20Mock));
        vm.serializeAddress(deployed_addresses, "erc20MockStrategy", address(erc20MockStrategy));
        vm.serializeAddress(deployed_addresses, "sfflServiceManager", address(sfflServiceManager));
        vm.serializeAddress(
            deployed_addresses, "sfflServiceManagerImplementation", address(sfflServiceManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "sfflTaskManager", address(sfflTaskManager));
        vm.serializeAddress(deployed_addresses, "sfflTaskManagerImplementation", address(sfflTaskManagerImplementation));
        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinator));
        vm.serializeAddress(
            deployed_addresses, "registryCoordinatorImplementation", address(registryCoordinatorImplementation)
        );
        string memory deployed_addresses_output =
            vm.serializeAddress(deployed_addresses, "operatorStateRetriever", address(operatorStateRetriever));

        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        writeOutput(finalJson, "sffl_avs_deployment_output");
    }
}
