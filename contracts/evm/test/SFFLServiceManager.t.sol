// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {ServiceManagerBase} from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";

import {SFFLTaskManager} from "../src/eth/SFFLTaskManager.sol";
import {SFFLServiceManager} from "../src/eth/SFFLServiceManager.sol";
import {Checkpoint} from "../src/eth/task/Checkpoint.sol";
import {StateRootUpdate} from "../src/base/message/StateRootUpdate.sol";
import {SFFLOperatorSetUpdateRegistry} from "../src/eth/SFFLOperatorSetUpdateRegistry.sol";

import {TestUtils} from "./utils/TestUtils.sol";

contract SFFLServiceManagerHarness is SFFLServiceManager {
    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        SFFLTaskManager _taskManager,
        SFFLOperatorSetUpdateRegistry _operatorSetUpdateRegistry
    )
        SFFLServiceManager(_avsDirectory, _registryCoordinator, _stakeRegistry, _taskManager, _operatorSetUpdateRegistry)
    {}

    function forceInitialize(address initialOwner, IPauserRegistry _pauserRegistry) public {
        _transferOwnership(initialOwner);
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
    }

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public override onlyRegistryCoordinator {
        ServiceManagerBase.registerOperatorToAVS(operator, operatorSignature);
    }

    function registerOperatorToAVSWhitelisted(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public onlyRegistryCoordinator {
        SFFLServiceManager.registerOperatorToAVS(operator, operatorSignature);
    }
}

contract SFFLServiceManagerTest is TestUtils {
    using StateRootUpdate for StateRootUpdate.Message;

    event StateRootUpdated(uint32 indexed rollupId, uint64 indexed blockHeight, bytes32 stateRoot);

    SFFLServiceManagerHarness public sfflServiceManager;
    SFFLTaskManager public taskManager;
    SFFLOperatorSetUpdateRegistry operatorSetUpdateRegistry;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address public aggregator;
    address public generator;
    uint256 public thresholdDenominator;

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        aggregator = addr("aggregator");
        generator = addr("generator");

        address impl = address(new SFFLTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK));

        taskManager = SFFLTaskManager(
            deployProxy(
                impl,
                address(proxyAdmin),
                abi.encodeWithSelector(
                    taskManager.initialize.selector, pauserRegistry, registryCoordinatorOwner, aggregator, generator
                )
            )
        );

        vm.label(impl, "taskManagerImpl");
        vm.label(address(taskManager), "taskManagerProxy");

        operatorSetUpdateRegistry = new SFFLOperatorSetUpdateRegistry(registryCoordinator);

        address sfflServiceManagerImplementation = address(
            new SFFLServiceManagerHarness(
                IAVSDirectory(avsDirectoryMock),
                registryCoordinator,
                stakeRegistry,
                taskManager,
                operatorSetUpdateRegistry
            )
        );

        vm.prank(proxyAdminOwner);
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(serviceManager))),
            sfflServiceManagerImplementation,
            abi.encodeWithSignature("forceInitialize(address,address)", serviceManagerOwner, address(pauserRegistry))
        );

        sfflServiceManager = SFFLServiceManagerHarness(address(serviceManager));

        vm.label(sfflServiceManagerImplementation, "serviceManagerImpl");
        vm.label(address(serviceManager), "serviceManagerProxy");

        thresholdDenominator = taskManager.THRESHOLD_DENOMINATOR();
    }

    function test_initialize() public {
        address newOwner = address(0x123);
        IPauserRegistry newPauserRegistry = IPauserRegistry(address(0x456));

        address impl = address(
            new SFFLServiceManagerHarness(
                IAVSDirectory(avsDirectoryMock),
                registryCoordinator,
                stakeRegistry,
                taskManager,
                SFFLOperatorSetUpdateRegistry(address(0))
            )
        );

        SFFLServiceManagerHarness newServiceManager = SFFLServiceManagerHarness(
            deployProxy(
                impl,
                address(proxyAdmin),
                abi.encodeWithSignature("initialize(address,address)", newOwner, newPauserRegistry)
            )
        );

        assertEq(newServiceManager.owner(), newOwner);
        assertEq(address(newServiceManager.pauserRegistry()), address(newPauserRegistry));
    }

    function test_onlyTaskManager_RevertWhen_NonTaskManagerCall() public {
        address nonTaskManager = address(0x123);
        vm.prank(nonTaskManager);
        vm.expectRevert("Task manager must be the caller");
        sfflServiceManager.freezeOperator(address(0x456));
    }

    function test_constructor() public {
        assertEq(address(sfflServiceManager.taskManager()), address(taskManager));
        assertEq(address(sfflServiceManager.operatorSetUpdateRegistry()), address(operatorSetUpdateRegistry));
    }

    function test_initialize_PauserSetToZero() public {
        address newOwner = address(0x123);

        address impl = address(
            new SFFLServiceManagerHarness(
                IAVSDirectory(avsDirectoryMock),
                registryCoordinator,
                stakeRegistry,
                taskManager,
                SFFLOperatorSetUpdateRegistry(address(0))
            )
        );

        vm.expectRevert("Pausable._initializePauser: _initializePauser() can only be called once");
        deployProxy(impl, address(proxyAdmin), abi.encodeWithSignature("initialize(address)", newOwner));
    }

    function test_initialize_PauserSetToPauserRegistry() public {
        address newOwner = address(0x123);
        IPauserRegistry newPauserRegistry = IPauserRegistry(address(0x456));

        address impl = address(
            new SFFLServiceManagerHarness(
                IAVSDirectory(avsDirectoryMock),
                registryCoordinator,
                stakeRegistry,
                taskManager,
                SFFLOperatorSetUpdateRegistry(address(0))
            )
        );

        SFFLServiceManagerHarness newServiceManager = SFFLServiceManagerHarness(
            deployProxy(
                impl,
                address(proxyAdmin),
                abi.encodeWithSignature("initialize(address,address)", newOwner, newPauserRegistry)
            )
        );

        assertEq(newServiceManager.owner(), newOwner);
        assertEq(address(newServiceManager.pauserRegistry()), address(newPauserRegistry));
    }

    function test_registerOperatorToAVS_RevertWhen_NotWhitelisted() public {
        address operator = address(0x789);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature;

        vm.mockCall(
            address(sfflServiceManager.operatorSetUpdateRegistry()),
            abi.encodeWithSignature("isOperatorWhitelisted(address)", operator),
            abi.encode(false)
        );

        vm.prank(address(registryCoordinator));
        vm.expectRevert("Not whitelisted");
        sfflServiceManager.registerOperatorToAVSWhitelisted(operator, operatorSignature);
    }

    function testFail_registerOperatorToAVS_Whitelisted() public {
        address operator = address(0x789);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature;

        vm.mockCall(
            address(sfflServiceManager.operatorSetUpdateRegistry()),
            abi.encodeWithSignature("isOperatorWhitelisted(address)", operator),
            abi.encode(true)
        );

        vm.prank(address(registryCoordinator));
        vm.expectRevert("Not whitelisted");
        sfflServiceManager.registerOperatorToAVSWhitelisted(operator, operatorSignature);
    }

    function test_updateStateRoot() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 1,
            timestamp: 2,
            nearDaTransactionId: keccak256(hex"03"),
            nearDaCommitment: keccak256(hex"04"),
            stateRoot: keccak256(hex"f00d")
        });

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            setUpOperators(message.hash(), 999, 1000, 100, 1);

        vm.expectEmit(true, true, false, true);
        emit StateRootUpdated(message.rollupId, message.blockHeight, message.stateRoot);

        assertEq(sfflServiceManager.getStateRoot(message.rollupId, message.blockHeight), bytes32(0));

        vm.roll(1001);
        sfflServiceManager.updateStateRoot(message, nonSignerStakesAndSignature);

        assertEq(sfflServiceManager.getStateRoot(message.rollupId, message.blockHeight), message.stateRoot);
    }

    function test_updateStateRoot_RevertWhen_QuorumNotMet() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 1,
            timestamp: 2,
            nearDaTransactionId: keccak256(hex"03"),
            nearDaCommitment: keccak256(hex"04"),
            stateRoot: keccak256(hex"f00d")
        });

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            setUpOperators(message.hash(), 999, 1000, 100, maxOperatorsToRegister / 2);

        vm.expectRevert("Quorum not met");

        vm.roll(1001);
        sfflServiceManager.updateStateRoot(message, nonSignerStakesAndSignature);
    }

    function test_updateStateRoot_RevertWhen_Paused() public {
        uint8 flag = sfflServiceManager.PAUSED_UPDATE_STATE_ROOT();

        vm.prank(pauser);
        sfflServiceManager.pause(2 ** flag);

        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 1,
            timestamp: 2,
            nearDaTransactionId: keccak256(hex"03"),
            nearDaCommitment: keccak256(hex"04"),
            stateRoot: keccak256(hex"f00d")
        });

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            setUpOperators(message.hash(), 999, 1000, 100, 1);

        vm.expectRevert("Pausable: index is paused");

        sfflServiceManager.updateStateRoot(message, nonSignerStakesAndSignature);
    }
}
