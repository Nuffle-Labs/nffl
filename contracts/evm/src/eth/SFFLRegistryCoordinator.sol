// SPDX-License-Identifier: MIT
pragma solidity =0.8.12;

import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {IBLSApkRegistry} from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import {RegistryCoordinator} from "../external/RegistryCoordinator.sol";
import {SFFLOperatorSetUpdateRegistry} from "./SFFLOperatorSetUpdateRegistry.sol";

/**
 * @title SFFL AVS Registry Coordinator
 * @notice Coordinator for various registries in an AVS - in this case,
 * the base {StakeRegistry}, {BLSApkRegistry} and {IndexRegistry}, but also
 * {SFFLOperatorSetUpdateRegistry}.
 * This contract's behavior is basically similar to a base RegistryCoordinator,
 * with mainly the addition of operator set change tracking through
 * {SFFLOperatorSetUpdateRegistry} in order to trigger the SFFL AVS to update
 * the rollups' operator set copies.
 */
contract SFFLRegistryCoordinator is RegistryCoordinator {
    SFFLOperatorSetUpdateRegistry public immutable operatorSetUpdateRegistry;

    constructor(
        IServiceManager _serviceManager,
        IStakeRegistry _stakeRegistry,
        IBLSApkRegistry _blsApkRegistry,
        IIndexRegistry _indexRegistry,
        SFFLOperatorSetUpdateRegistry _operatorSetUpdateRegistry
    ) RegistryCoordinator(_serviceManager, _stakeRegistry, _blsApkRegistry, _indexRegistry) {
        operatorSetUpdateRegistry = _operatorSetUpdateRegistry;
    }

    // TODO: Include operator set update registry in registries. Currently
    // RegistryCoordinator code size is already tight enough for any
    // extensions.

    // /**
    //  * @inheritdoc RegistryCoordinator
    //  */
    // function initialize(
    //     address _initialOwner,
    //     address _churnApprover,
    //     address _ejector,
    //     IPauserRegistry _pauserRegistry,
    //     uint256 _initialPausedStatus,
    //     OperatorSetParam[] memory _operatorSetParams,
    //     uint96[] memory _minimumStakes,
    //     IStakeRegistry.StrategyParams[][] memory _strategyParams
    // ) public override initializer {
    //     RegistryCoordinator.initialize(
    //         _initialOwner,
    //         _churnApprover,
    //         _ejector,
    //         _pauserRegistry,
    //         _initialPausedStatus,
    //         _operatorSetParams,
    //         _minimumStakes,
    //         _strategyParams
    //     );

    //     registries.push(address(operatorSetUpdateRegistry));
    // }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function registerOperator(
        bytes calldata quorumNumbers,
        string calldata socket,
        IBLSApkRegistry.PubkeyRegistrationParams calldata params,
        SignatureWithSaltAndExpiry memory operatorSignature
    ) public override {
        _recordOperatorSetUpdate();
        RegistryCoordinator.registerOperator(quorumNumbers, socket, params, operatorSignature);
    }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function registerOperatorWithChurn(
        bytes calldata quorumNumbers,
        string calldata socket,
        IBLSApkRegistry.PubkeyRegistrationParams calldata params,
        OperatorKickParam[] calldata operatorKickParams,
        SignatureWithSaltAndExpiry memory churnApproverSignature,
        SignatureWithSaltAndExpiry memory operatorSignature
    ) public override {
        _recordOperatorSetUpdate();
        RegistryCoordinator.registerOperatorWithChurn(
            quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature
        );
    }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function deregisterOperator(bytes calldata quorumNumbers) public override {
        _recordOperatorSetUpdate();
        RegistryCoordinator.deregisterOperator(quorumNumbers);
    }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function updateOperators(address[] calldata operators) public override {
        _recordOperatorSetUpdate();
        RegistryCoordinator.updateOperators(operators);
    }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function updateOperatorsForQuorum(address[][] calldata operatorsPerQuorum, bytes calldata quorumNumbers)
        public
        override
    {
        _recordOperatorSetUpdate();
        RegistryCoordinator.updateOperatorsForQuorum(operatorsPerQuorum, quorumNumbers);
    }

    /**
     * @inheritdoc RegistryCoordinator
     */
    function ejectOperator(address operator, bytes calldata quorumNumbers) public override {
        _recordOperatorSetUpdate();
        RegistryCoordinator.ejectOperator(operator, quorumNumbers);
    }

    /**
     * @dev Notifies {SFFLOperatorSetUpdateRegistry} of an operator set update
     * in this block.
     */
    function _recordOperatorSetUpdate() internal {
        operatorSetUpdateRegistry.recordOperatorSetUpdate();
    }
}
