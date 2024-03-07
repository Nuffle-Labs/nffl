// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {BLSMockAVSDeployer} from "eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";
import {ServiceManagerBase} from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IBLSSignatureChecker} from "eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import {SFFLTaskManager} from "../src/eth/SFFLTaskManager.sol";
import {SFFLServiceManager} from "../src/eth/SFFLServiceManager.sol";
import {Checkpoint} from "../src/eth/task/Checkpoint.sol";
import {StateRootUpdate} from "../src/base/message/StateRootUpdate.sol";

import {TestUtils} from "./utils/TestUtils.sol";

contract SFFLServiceManagerTest is TestUtils {
    using StateRootUpdate for StateRootUpdate.Message;

    event StateRootUpdated(uint32 indexed rollupId, uint64 indexed blockHeight, bytes32 stateRoot);

    SFFLServiceManager public sfflServiceManager;
    SFFLTaskManager public taskManager;

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

        address sfflServiceManagerImplementation = address(
            new SFFLServiceManager(IDelegationManager(delegationMock), registryCoordinator, stakeRegistry, taskManager)
        );

        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(serviceManager))), sfflServiceManagerImplementation
        );

        sfflServiceManager = SFFLServiceManager(address(serviceManager));

        vm.label(sfflServiceManagerImplementation, "serviceManagerImpl");
        vm.label(address(serviceManager), "serviceManagerProxy");

        thresholdDenominator = taskManager.THRESHOLD_DENOMINATOR();
    }

    function test_updateStateRoot() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 1,
            timestamp: 2,
            stateRoot: bytes32(keccak256(hex"f00d"))
        });

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            setUpOperators(message.hash(), 1000, 100, 1);

        vm.expectEmit(true, true, false, true);
        emit StateRootUpdated(message.rollupId, message.blockHeight, message.stateRoot);

        assertEq(sfflServiceManager.getStateRoot(message.rollupId, message.blockHeight), bytes32(0));

        sfflServiceManager.updateStateRoot(message, nonSignerStakesAndSignature);

        assertEq(sfflServiceManager.getStateRoot(message.rollupId, message.blockHeight), message.stateRoot);
    }

    function test_updateStateRoot_RevertWhen_QuorumNotMet() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message({
            rollupId: 0,
            blockHeight: 1,
            timestamp: 2,
            stateRoot: bytes32(keccak256(hex"f00d"))
        });

        (, IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature) =
            setUpOperators(message.hash(), 1000, 100, maxOperatorsToRegister / 2);

        vm.expectRevert("Quorum not met");

        sfflServiceManager.updateStateRoot(message, nonSignerStakesAndSignature);
    }
}
