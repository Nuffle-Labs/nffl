// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
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

import {RegistryCoordinator} from "../../../src/external/RegistryCoordinator.sol";
import {SFFLServiceManager} from "../../../src/eth/SFFLServiceManager.sol";
import {SFFLTaskManager} from "../../../src/eth/SFFLTaskManager.sol";
import {SFFLRegistryCoordinator} from "../../../src/eth/SFFLRegistryCoordinator.sol";
import {SFFLOperatorSetUpdateRegistry} from "../../../src/eth/SFFLOperatorSetUpdateRegistry.sol";
import {ERC20Mock, IERC20} from "../../../test/mock/ERC20Mock.sol";

import {Utils} from "../../utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract SFFLDeployer is Script, Utils {
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;

    uint32 public constant MAX_OPERATOR_COUNT = 10000;
    uint16 public constant KICK_BIPS_OF_OPERATOR_STAKE = 15000;
    uint16 public constant KICK_BIPS_OF_TOTAL_STAKE = 100;

    uint96 public constant STRATEGY_MULTIPLIER = 1 ether;
    uint256 public constant WEIGHTING_DIVISOR = 1 ether;

    uint256 public constant NUM_QUORUMS = 1;
    uint256 public constant INITIAL_PAUSED_STATUS = 0;

    uint256 public constant MOCK_STRATEGY_MAX_PER_DEPOSIT = 1 ether;
    uint256 public constant MOCK_STRATEGY_MAX_DEPOSITS = 100 ether;

    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    string public constant EIGENLAYER_DEPLOYMENT_FILE = "eigenlayer_deployment_output";
    string public constant SFFL_DEPLOYMENT_FILE = "sffl_avs_deployment_output";

    string public constant PROTOCOL_VERSION = "v0.0.1-devnet";

    struct EigenlayerDeployedContracts {
        IStrategyManager strategyManager;
        IDelegationManager delegationManager;
        IAVSDirectory avsDirectory;
        ProxyAdmin eigenLayerProxyAdmin;
        PauserRegistry eigenLayerPauserReg;
        StrategyBaseTVLLimits baseStrategyImpl;
    }

    EmptyContract public emptyContract;
    ERC20Mock public erc20Mock;

    StrategyBaseTVLLimits public erc20MockStrategy;
    TransparentUpgradeableProxy public erc20MockStrategyProxy;

    // SFFL contracts
    ProxyAdmin public sfflProxyAdmin;
    PauserRegistry public sfflPauserReg;

    SFFLRegistryCoordinator public registryCoordinator;
    TransparentUpgradeableProxy public registryCoordinatorProxy;
    address public registryCoordinatorImpl;

    IBLSApkRegistry public blsApkRegistry;
    TransparentUpgradeableProxy public blsApkRegistryProxy;
    address public blsApkRegistryImpl;

    IIndexRegistry public indexRegistry;
    TransparentUpgradeableProxy public indexRegistryProxy;
    address public indexRegistryImpl;

    IStakeRegistry public stakeRegistry;
    TransparentUpgradeableProxy public stakeRegistryProxy;
    address public stakeRegistryImpl;

    SFFLOperatorSetUpdateRegistry public operatorSetUpdateRegistry;
    TransparentUpgradeableProxy public operatorSetUpdateRegistryProxy;
    address public operatorSetUpdateRegistryImpl;

    SFFLServiceManager public sfflServiceManager;
    TransparentUpgradeableProxy public sfflServiceManagerProxy;
    address public sfflServiceManagerImpl;

    SFFLTaskManager public sfflTaskManager;
    TransparentUpgradeableProxy public sfflTaskManagerProxy;
    address public sfflTaskManagerImpl;

    OperatorStateRetriever public operatorStateRetriever;

    function run() external {
        EigenlayerDeployedContracts memory eigenlayerContracts = _readEigenlayerDeployedContracts();

        address sfflCommunityMultisig = msg.sender;
        address sfflPauser = msg.sender;

        vm.startBroadcast();

        emptyContract = new EmptyContract();

        _deployErc20AndStrategyAndWhitelistStrategy(
            eigenlayerContracts.eigenLayerProxyAdmin,
            eigenlayerContracts.eigenLayerPauserReg,
            eigenlayerContracts.baseStrategyImpl,
            eigenlayerContracts.strategyManager
        );
        _deploySFFLContracts(
            eigenlayerContracts.delegationManager,
            eigenlayerContracts.avsDirectory,
            erc20MockStrategy,
            sfflCommunityMultisig,
            sfflPauser
        );

        _whitelistOperators();
        vm.stopBroadcast();
    }

    /**
     * @dev Reads the output of the Eigenlayer deployment script from the forge output and returns the EL deployed contracts.
     */
    function _readEigenlayerDeployedContracts() internal view returns (EigenlayerDeployedContracts memory) {
        string memory json = readOutput(EIGENLAYER_DEPLOYMENT_FILE);
        return EigenlayerDeployedContracts({
            strategyManager: IStrategyManager(stdJson.readAddress(json, ".addresses.strategyManager")),
            delegationManager: IDelegationManager(stdJson.readAddress(json, ".addresses.delegationManager")),
            avsDirectory: IAVSDirectory(stdJson.readAddress(json, ".addresses.avsDirectory")),
            eigenLayerProxyAdmin: ProxyAdmin(stdJson.readAddress(json, ".addresses.eigenLayerProxyAdmin")),
            eigenLayerPauserReg: PauserRegistry(stdJson.readAddress(json, ".addresses.eigenLayerPauserReg")),
            baseStrategyImpl: StrategyBaseTVLLimits(stdJson.readAddress(json, ".addresses.baseStrategyImplementation"))
        });
    }

    /**
     * @dev Deploys the ERC20 mock and the strategy and whitelists the strategy.
     */
    function _deployErc20AndStrategyAndWhitelistStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImpl,
        IStrategyManager strategyManager
    ) internal {
        erc20Mock = new ERC20Mock();

        erc20MockStrategyProxy = _deployProxy(
            eigenLayerProxyAdmin,
            address(baseStrategyImpl),
            abi.encodeWithSelector(
                StrategyBaseTVLLimits.initialize.selector,
                MOCK_STRATEGY_MAX_PER_DEPOSIT,
                MOCK_STRATEGY_MAX_DEPOSITS,
                IERC20(erc20Mock),
                eigenLayerPauserReg
            )
        );
        erc20MockStrategy = StrategyBaseTVLLimits(address(erc20MockStrategyProxy));

        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;

        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);

        strategyManager.addStrategiesToDepositWhitelist(strats, thirdPartyTransfersForbiddenValues);
    }

    function _whitelistOperators() internal {
        // from keys in tests/keys/ecdsa
        operatorSetUpdateRegistry.setOperatorWhitelisting(0xD5A0359da7B310917d7760385516B2426E86ab7f, true);
        operatorSetUpdateRegistry.setOperatorWhitelisting(0x9441540E8183d416f2Dc1901AB2034600f17B65a, true);
        operatorSetUpdateRegistry.setOperatorWhitelisting(0x49d0D93C30f799343745d482695a0Fdb952B1d02, true);
        operatorSetUpdateRegistry.setOperatorWhitelisting(0x4b35F09961ed53545f7508f5ac1e8414D7c31D7A, true);
    }

    /**
     * @dev Deploys the SFFL contracts.
     * @param delegationManager The delegation manager.
     * @param avsDirectory The AVS directory.
     * @param strat The deployed strategy.
     * @param sfflCommunityMultisig The community multisig.
     * @param sfflPauser The pauser.
     */
    function _deploySFFLContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IStrategy strat,
        address sfflCommunityMultisig,
        address sfflPauser
    ) internal {
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint256 numStrategies = deployedStrategyArray.length;

        sfflProxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](2);
        pausers[0] = sfflPauser;
        pausers[1] = sfflCommunityMultisig;

        sfflPauserReg = new PauserRegistry(pausers, sfflCommunityMultisig);

        sfflServiceManagerProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        sfflServiceManager = SFFLServiceManager(address(sfflServiceManagerProxy));

        sfflTaskManagerProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        sfflTaskManager = SFFLTaskManager(address(sfflTaskManagerProxy));

        registryCoordinatorProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        registryCoordinator = SFFLRegistryCoordinator(address(registryCoordinatorProxy));

        blsApkRegistryProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        blsApkRegistry = BLSApkRegistry(address(blsApkRegistryProxy));

        indexRegistryProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        indexRegistry = IndexRegistry(address(indexRegistryProxy));

        operatorSetUpdateRegistryProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        operatorSetUpdateRegistry = SFFLOperatorSetUpdateRegistry(address(operatorSetUpdateRegistryProxy));

        stakeRegistryProxy = _deployEmptyProxy(sfflProxyAdmin, address(emptyContract));
        stakeRegistry = StakeRegistry(address(stakeRegistryProxy));

        operatorStateRetriever = new OperatorStateRetriever();

        stakeRegistryImpl = address(new StakeRegistry(registryCoordinator, delegationManager));
        _upgradeProxy(sfflProxyAdmin, stakeRegistryProxy, stakeRegistryImpl);

        blsApkRegistryImpl = address(new BLSApkRegistry(registryCoordinator));
        _upgradeProxy(sfflProxyAdmin, blsApkRegistryProxy, blsApkRegistryImpl);

        indexRegistryImpl = address(new IndexRegistry(registryCoordinator));
        _upgradeProxy(sfflProxyAdmin, indexRegistryProxy, indexRegistryImpl);

        operatorSetUpdateRegistryImpl = address(new SFFLOperatorSetUpdateRegistry(registryCoordinator));
        _upgradeProxy(sfflProxyAdmin, operatorSetUpdateRegistryProxy, operatorSetUpdateRegistryImpl);

        registryCoordinatorImpl = address(
            new SFFLRegistryCoordinator(
                sfflServiceManager, stakeRegistry, blsApkRegistry, indexRegistry, operatorSetUpdateRegistry
            )
        );

        IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams =
            new IRegistryCoordinator.OperatorSetParam[](NUM_QUORUMS);

        for (uint256 i = 0; i < quorumsOperatorSetParams.length; i++) {
            quorumsOperatorSetParams[i] = IRegistryCoordinator.OperatorSetParam({
                maxOperatorCount: MAX_OPERATOR_COUNT,
                kickBIPsOfOperatorStake: KICK_BIPS_OF_OPERATOR_STAKE,
                kickBIPsOfTotalStake: KICK_BIPS_OF_TOTAL_STAKE
            });
        }

        uint96[] memory quorumsMinimumStake = new uint96[](NUM_QUORUMS);
        IStakeRegistry.StrategyParams[][] memory quorumsStrategyParams =
            new IStakeRegistry.StrategyParams[][](NUM_QUORUMS);

        for (uint256 i = 0; i < quorumsStrategyParams.length; i++) {
            quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);

            for (uint256 j = 0; j < quorumsStrategyParams[i].length; j++) {
                quorumsStrategyParams[i][j] =
                    IStakeRegistry.StrategyParams({strategy: deployedStrategyArray[j], multiplier: STRATEGY_MULTIPLIER});
            }
        }

        _upgradeProxyAndCall(
            sfflProxyAdmin,
            registryCoordinatorProxy,
            registryCoordinatorImpl,
            abi.encodeWithSelector(
                registryCoordinator.initialize.selector,
                sfflCommunityMultisig,
                sfflCommunityMultisig,
                sfflCommunityMultisig,
                sfflPauserReg,
                INITIAL_PAUSED_STATUS,
                quorumsOperatorSetParams,
                quorumsMinimumStake,
                quorumsStrategyParams
            )
        );

        sfflServiceManagerImpl = address(
            new SFFLServiceManager(
                avsDirectory, registryCoordinator, stakeRegistry, sfflTaskManager, operatorSetUpdateRegistry
            )
        );

        _upgradeProxyAndCall(
            sfflProxyAdmin,
            sfflServiceManagerProxy,
            sfflServiceManagerImpl,
            abi.encodeWithSignature("initialize(address,address)", sfflCommunityMultisig, sfflPauserReg)
        );

        sfflTaskManagerImpl =
            address(new SFFLTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK, address(sfflTaskManager), PROTOCOL_VERSION));

        _upgradeProxyAndCall(
            sfflProxyAdmin,
            sfflTaskManagerProxy,
            sfflTaskManagerImpl,
            abi.encodeWithSelector(
                SFFLTaskManager.initialize.selector,
                sfflPauserReg,
                sfflCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR
            )
        );

        _serializeSFFLDeployedContracts();
    }

    /**
     * @dev Serializes the SFFL deployed contracts to the forge output.
     */
    function _serializeSFFLDeployedContracts() internal {
        string memory parent_object = "parent object";
        string memory addresses = "addresses";

        string memory addressesOutput;

        addressesOutput = vm.serializeAddress(addresses, "deployer", address(msg.sender));
        addressesOutput = vm.serializeAddress(addresses, "erc20Mock", address(erc20Mock));
        addressesOutput = vm.serializeAddress(addresses, "erc20MockStrategy", address(erc20MockStrategy));
        addressesOutput = vm.serializeAddress(addresses, "sfflProxyAdmin", address(sfflProxyAdmin));
        addressesOutput = vm.serializeAddress(addresses, "sfflPauserReg", address(sfflPauserReg));
        addressesOutput = vm.serializeAddress(addresses, "registryCoordinator", address(registryCoordinator));
        addressesOutput = vm.serializeAddress(addresses, "registryCoordinatorImpl", address(registryCoordinatorImpl));
        addressesOutput = vm.serializeAddress(addresses, "blsApkRegistry", address(blsApkRegistry));
        addressesOutput = vm.serializeAddress(addresses, "blsApkRegistryImpl", address(blsApkRegistryImpl));
        addressesOutput = vm.serializeAddress(addresses, "indexRegistry", address(indexRegistry));
        addressesOutput = vm.serializeAddress(addresses, "indexRegistryImpl", address(indexRegistryImpl));
        addressesOutput = vm.serializeAddress(addresses, "stakeRegistry", address(stakeRegistry));
        addressesOutput = vm.serializeAddress(addresses, "stakeRegistryImpl", address(stakeRegistryImpl));
        addressesOutput =
            vm.serializeAddress(addresses, "operatorSetUpdateRegistry", address(operatorSetUpdateRegistry));
        addressesOutput =
            vm.serializeAddress(addresses, "operatorSetUpdateRegistryImpl", address(operatorSetUpdateRegistryImpl));
        addressesOutput = vm.serializeAddress(addresses, "sfflServiceManager", address(sfflServiceManager));
        addressesOutput = vm.serializeAddress(addresses, "sfflServiceManagerImpl", address(sfflServiceManagerImpl));
        addressesOutput = vm.serializeAddress(addresses, "sfflTaskManager", address(sfflTaskManager));
        addressesOutput = vm.serializeAddress(addresses, "sfflTaskManagerImpl", address(sfflTaskManagerImpl));
        addressesOutput = vm.serializeAddress(addresses, "operatorStateRetriever", address(operatorStateRetriever));

        string memory chainInfo = "chainInfo";
        string memory chainInfoOutput;
        chainInfoOutput = vm.serializeUint(chainInfo, "chainId", block.chainid);
        chainInfoOutput = vm.serializeUint(chainInfo, "deploymentBlock", block.number);

        string memory finalJson;
        finalJson = vm.serializeString(parent_object, addresses, addressesOutput);
        finalJson = vm.serializeString(parent_object, chainInfo, chainInfoOutput);

        writeOutput(finalJson, SFFL_DEPLOYMENT_FILE);
    }
}
