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
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";

import {SFFLTaskManager} from "../src/eth/SFFLTaskManager.sol";
import {SFFLOperatorSetUpdateRegistry} from "../src/eth/SFFLOperatorSetUpdateRegistry.sol";
import {SFFLServiceManager} from "../src/eth/SFFLServiceManager.sol";
import {SFFLRegistryCoordinator} from "../src/eth/SFFLRegistryCoordinator.sol";
import {RollupOperators} from "../src/base/utils/RollupOperators.sol";

import {TestUtils} from "./utils/TestUtils.sol";

contract SFFLServiceManagerHarness is SFFLServiceManager, Test {
    bool public autoWhitelist;

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
        if (autoWhitelist) {
            vm.prank(_registryCoordinator.owner());
            operatorSetUpdateRegistry.setOperatorWhitelisting(operator, true);
        }

        SFFLServiceManager.registerOperatorToAVS(operator, operatorSignature);
    }

    function setAutoWhitelist(bool _autoWhitelist) public {
        autoWhitelist = _autoWhitelist;
    }
}

contract SFFLOperatorSetUpdateRegistryTest is TestUtils {
    using BN254 for BN254.G1Point;

    event OperatorSetUpdatedAtBlock(uint64 indexed id, uint64 indexed timestamp);
    event OperatorWhitelistingUpdated(address indexed operator, bool isWhitelisted);

    SFFLOperatorSetUpdateRegistry public operatorSetUpdateRegistry;
    SFFLServiceManagerHarness public sfflServiceManager;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address public aggregator;
    address public generator;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        aggregator = addr("aggregator");
        generator = addr("generator");

        address impl = address(new SFFLTaskManager(registryCoordinator, TASK_RESPONSE_WINDOW_BLOCK));

        SFFLTaskManager taskManager = SFFLTaskManager(
            deployProxy(
                impl,
                address(proxyAdmin),
                abi.encodeWithSelector(
                    SFFLTaskManager.initialize.selector, pauserRegistry, registryCoordinatorOwner, aggregator, generator
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

        address registryCoordinatorImpl = address(
            new SFFLRegistryCoordinator(
                serviceManager, stakeRegistry, blsApkRegistry, indexRegistry, operatorSetUpdateRegistry
            )
        );

        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(registryCoordinator))), registryCoordinatorImpl);
    }

    function test_constructor() public {
        assertEq(address(operatorSetUpdateRegistry.registryCoordinator()), address(registryCoordinator));
    }

    function test_getOperatorSetUpdateCount() public {
        assertEq(operatorSetUpdateRegistry.getOperatorSetUpdateCount(), 0);

        vm.prank(address(registryCoordinator));
        operatorSetUpdateRegistry.recordOperatorSetUpdate();

        assertEq(operatorSetUpdateRegistry.getOperatorSetUpdateCount(), 1);
    }

    function test_recordOperatorSetUpdate() public {
        vm.prank(address(registryCoordinator));
        vm.expectEmit(true, true, false, true);
        emit OperatorSetUpdatedAtBlock(0, uint64(block.timestamp));
        operatorSetUpdateRegistry.recordOperatorSetUpdate();

        assertEq(operatorSetUpdateRegistry.operatorSetUpdateIdToBlockNumber(0), uint32(block.number));
    }

    function test_recordOperatorSetUpdate_multipleInSameBlock() public {
        vm.prank(address(registryCoordinator));
        operatorSetUpdateRegistry.recordOperatorSetUpdate();

        vm.prank(address(registryCoordinator));
        operatorSetUpdateRegistry.recordOperatorSetUpdate();

        assertEq(operatorSetUpdateRegistry.getOperatorSetUpdateCount(), 1);
    }

    function test_recordOperatorSetUpdate_RevertWhen_NotRegistryCoordinator() public {
        vm.expectRevert("SFFLOperatorSetUpdateRegistry.onlyRegistryCoordinator: caller is not the registry coordinator");
        operatorSetUpdateRegistry.recordOperatorSetUpdate();
    }

    function test_getOperatorSetUpdate() public {
        sfflServiceManager.setAutoWhitelist(true);
        setUpOperators(keccak256("test"), 999, 1000, 100, 1);

        assertEq(operatorSetUpdateRegistry.getOperatorSetUpdateCount(), maxOperatorsToRegister);

        vm.roll(block.number + 1);

        vm.prank(address(registryCoordinator));
        operatorSetUpdateRegistry.recordOperatorSetUpdate();

        assertEq(operatorSetUpdateRegistry.getOperatorSetUpdateCount(), maxOperatorsToRegister + 1);

        (RollupOperators.Operator[] memory previousOperatorSet, RollupOperators.Operator[] memory newOperatorSet) =
            operatorSetUpdateRegistry.getOperatorSetUpdate(1);

        assertEq(previousOperatorSet.length, 1);
        assertEq(newOperatorSet.length, 2);
        assertNotEq(keccak256(abi.encode(previousOperatorSet)), keccak256(abi.encode(newOperatorSet)));

        (previousOperatorSet, newOperatorSet) =
            operatorSetUpdateRegistry.getOperatorSetUpdate(uint64(maxOperatorsToRegister));

        assertEq(previousOperatorSet.length, maxOperatorsToRegister);
        assertEq(newOperatorSet.length, maxOperatorsToRegister);
        assertEq(keccak256(abi.encode(previousOperatorSet)), keccak256(abi.encode(newOperatorSet)));

        (previousOperatorSet, newOperatorSet) = operatorSetUpdateRegistry.getOperatorSetUpdate(0);

        assertEq(previousOperatorSet.length, 0);
        assertEq(newOperatorSet.length, 1);
    }

    function test_setOperatorWhitelisting() public {
        address operator = address(0x123);

        vm.prank(registryCoordinatorOwner);
        vm.expectEmit(true, false, false, true);
        emit OperatorWhitelistingUpdated(operator, true);
        operatorSetUpdateRegistry.setOperatorWhitelisting(operator, true);

        assertTrue(operatorSetUpdateRegistry.isOperatorWhitelisted(operator));

        vm.prank(registryCoordinatorOwner);
        vm.expectEmit(true, false, false, true);
        emit OperatorWhitelistingUpdated(operator, false);
        operatorSetUpdateRegistry.setOperatorWhitelisting(operator, false);

        assertFalse(operatorSetUpdateRegistry.isOperatorWhitelisted(operator));
    }

    function test_setOperatorWhitelisting_RevertWhen_NotCoordinatorOwner() public {
        address operator = address(0x123);

        vm.expectRevert(
            "SFFLOperatorSetUpdateRegistry.onlyCoordinatorOwner: caller is not the owner of the registryCoordinator"
        );
        operatorSetUpdateRegistry.setOperatorWhitelisting(operator, true);
    }
}
