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
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 66;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 100;

    uint32 public constant MAX_OPERATOR_COUNT = 10;
    uint16 public constant KICK_BIPS_OF_OPERATOR_STAKE = 15000;
    uint16 public constant KICK_BIPS_OF_TOTAL_STAKE = 100;

    uint96 public constant STRATEGY_MULTIPLIER = 1 ether;
    uint256 public constant WEIGHTING_DIVISOR = 1 ether;

    uint256 public constant NUM_QUORUMS = 1;
    uint256 public constant INITIAL_PAUSED_STATUS = 0;

    // TODO: right now hardcoding these
    address public constant AGGREGATOR_ADDR = 0xEEa35C4E6F933fC04c853a371e170eDA1124c10d;
    address public constant TASK_GENERATOR_ADDR = 0xEEa35C4E6F933fC04c853a371e170eDA1124c10d;

    string public constant EIGENLAYER_DEPLOYMENT_FILE = "eigenlayer_deployment_output";
    string public constant SFFL_DEPLOYMENT_FILE = "sffl_avs_deployment_output";

    string public constant PROTOCOL_VERSION = "v0.0.1-holesky";

    struct EigenlayerDeployedContracts {
        IStrategyManager strategyManager;
        IDelegationManager delegationManager;
        IAVSDirectory avsDirectory;
        ProxyAdmin eigenLayerProxyAdmin;
        PauserRegistry eigenLayerPauserReg;
        StrategyBaseTVLLimits baseStrategyImpl;
    }

    EmptyContract public emptyContract;
    IStrategy[] strategies;

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

        // from the table in https://github.com/Layr-Labs/eigenlayer-contracts/blob/7229f2b426b6f2a24c7795b1a4687a010eac8ef2/README.md
        strategies.push(IStrategy(0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3));
        strategies.push(IStrategy(0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0));
        strategies.push(IStrategy(0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9));
        strategies.push(IStrategy(0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943));
        strategies.push(IStrategy(0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c));
        strategies.push(IStrategy(0x31B6F59e1627cEfC9fA174aD03859fC337666af7));
        strategies.push(IStrategy(0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef));
        strategies.push(IStrategy(0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6));
        strategies.push(IStrategy(0xaccc5A86732BE85b5012e8614AF237801636F8e5));
        strategies.push(IStrategy(0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac));
        strategies.push(IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0));

        vm.startBroadcast();

        emptyContract = new EmptyContract();

        _deploySFFLContracts(
            eigenlayerContracts.delegationManager, eigenlayerContracts.avsDirectory, sfflCommunityMultisig, sfflPauser
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

    function _whitelistOperators() internal {
        // from keys in tests/keys/ecdsa
        operatorSetUpdateRegistry.setOperatorWhitelisting(0xBF1BC6C45051E4330c3Ad921C78DCb242a1C2be6, true);
        operatorSetUpdateRegistry.setOperatorWhitelisting(0xE7444212A0d2394Ab25888F2371548837c410a5d, true);
    }

    /**
     * @dev Deploys the SFFL contracts.
     * @param delegationManager The delegation manager.
     * @param avsDirectory The AVS directory.
     * @param sfflCommunityMultisig The community multisig.
     * @param sfflPauser The pauser.
     */
    function _deploySFFLContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        address sfflCommunityMultisig,
        address sfflPauser
    ) internal {
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
            quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](strategies.length);

            for (uint256 j = 0; j < quorumsStrategyParams[i].length; j++) {
                quorumsStrategyParams[i][j] =
                    IStakeRegistry.StrategyParams({strategy: strategies[j], multiplier: STRATEGY_MULTIPLIER});
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
