// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import {IBLSApkRegistry} from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {IStakeRegistry} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {SFFLRegistryCoordinator} from "./SFFLRegistryCoordinator.sol";
import {RegistryCoordinator} from "../external/RegistryCoordinator.sol";
import {Operators as RollupOperators} from "../rollup/utils/Operators.sol";

/**
 * @title SFFL AVS Operator Set Update Registry
 * @notice This registry keeps track of operator set changes in order to
 * trigger the SFFL AVS to update the rollups' operator set copies.
 * @dev Operator set updates are block-based changes in the operator set which
 * are used by the AVS operators in order to update rollups' operator sets
 * (see {SFFLRegistryRollup}) through an {OperatorSetUpdate.Message}
 * attestation.
 * An operator set update is comprised of all the updates in operator weights
 * in one block, and as such happens at most once a block. It also has an
 * incrementing ID - which is then used on {OperatorSetUpdate.Message} and
 * could be used to fetch the update content for verifying evidences on bad
 * messages.
 */
contract SFFLOperatorSetUpdateRegistry is Initializable {
    /**
     * @notice Address of the RegistryCoordinator contract
     */
    SFFLRegistryCoordinator public immutable registryCoordinator;

    /**
     * @notice Reference block numbers for each operator set update
     */
    uint32[] public operatorSetUpdateIdToBlockNumber;

    /**
     * @dev Storage gap for upgradeability
     */
    uint256[50] private __GAP;

    /**
     * @notice Emitted when an operator set update is registered
     * @param id Operator set update ID
     */
    event OperatorSetUpdatedAtBlock(uint64 indexed id, uint64 indexed timestamp);

    /**
     * @dev Reverts if the caller is not the RegistryCoordinator contract
     */
    modifier onlyRegistryCoordinator() {
        require(
            msg.sender == address(registryCoordinator),
            "BLSApkRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"
        );
        _;
    }

    constructor(SFFLRegistryCoordinator _registryCoordinator) {
        registryCoordinator = _registryCoordinator;
        _disableInitializers();
    }

    /**
     * @notice Gets the count of how many operator set updates have happened
     * to date, which is also the ID of the next update
     * @return Operator set update count
     */
    function getOperatorSetUpdateCount() external view returns (uint64) {
        return uint64(operatorSetUpdateIdToBlockNumber.length);
    }

    /**
     * @dev Records an operator set update if necessary, i.e., if no other
     * update happened in the same block.
     * Emits {OperatorSetUpdatedAtBlock}.
     */
    function recordOperatorSetUpdate() external onlyRegistryCoordinator {
        uint64 id = uint64(operatorSetUpdateIdToBlockNumber.length);

        if (id > 0 && operatorSetUpdateIdToBlockNumber[id - 1] == block.number) {
            return;
        }

        emit OperatorSetUpdatedAtBlock(id, uint64(block.timestamp));

        operatorSetUpdateIdToBlockNumber.push(uint32(block.number));
    }

    /**
     * @notice Gets the previous and next operator sets for an operator set
     * update. This should be used by AVS operators to agree on operator set
     * updates to be pushed to rollups.
     * Important: this assumes the AVS has only a #0 quorum.
     * @dev This method's gas usage is high, and is meant for external calls,
     * not transactions.
     * @param operatorSetUpdateId Operator set update ID. Refer to
     * {SFFLRegistryCoordinator}
     * @return previousOperatorSet Operator set in the previous update, or an
     * empty set if operatorSetUpdateId is 0
     * @return newOperatorSet Operator set in the update indicated by
     * `operatorSetUpdateId`
     */
    function getOperatorSetUpdate(uint64 operatorSetUpdateId)
        external
        view
        returns (
            RollupOperators.Operator[] memory previousOperatorSet,
            RollupOperators.Operator[] memory newOperatorSet
        )
    {
        IStakeRegistry _stakeRegistry = registryCoordinator.stakeRegistry();
        IIndexRegistry _indexRegistry = registryCoordinator.indexRegistry();
        IBLSApkRegistry _blsApkRegistry = registryCoordinator.blsApkRegistry();

        if (operatorSetUpdateId > 0) {
            previousOperatorSet = _getOperatorSetAtBlock(
                operatorSetUpdateIdToBlockNumber[operatorSetUpdateId - 1],
                _stakeRegistry,
                _indexRegistry,
                _blsApkRegistry
            );
        }

        newOperatorSet = _getOperatorSetAtBlock(
            operatorSetUpdateIdToBlockNumber[operatorSetUpdateId], _stakeRegistry, _indexRegistry, _blsApkRegistry
        );
    }

    /**
     * @dev Gets the AVS operator set at a block number/height.
     * Important: This assumes the AVS has only a #0 quorum. This method's gas
     * usage is high, and is meant for usage in external calls, not
     * transactions.
     * @param blockNumber Block number for which to fetch the operator set
     * @param _stakeRegistry Address of the AVS's {IStakeRegistry}
     * @param _indexRegistry Address of the AVS's {IIndexRegistry}
     * @param _blsApkRegistry Address of the AVS's {IBlsApkRegistry}
     * @return Operator set at the specified block number
     */
    function _getOperatorSetAtBlock(
        uint32 blockNumber,
        IStakeRegistry _stakeRegistry,
        IIndexRegistry _indexRegistry,
        IBLSApkRegistry _blsApkRegistry
    ) internal view returns (RollupOperators.Operator[] memory) {
        bytes32[] memory operatorIds = _indexRegistry.getOperatorListAtBlockNumber(0, blockNumber);
        RollupOperators.Operator[] memory operators = new RollupOperators.Operator[](operatorIds.length);

        for (uint256 i = 0; i < operatorIds.length; i++) {
            bytes32 operatorId = operatorIds[i];

            address operator = _blsApkRegistry.getOperatorFromPubkeyHash(operatorId);
            (BN254.G1Point memory pubkey,) = _blsApkRegistry.getRegisteredPubkey(operator);
            uint96 stake = _stakeRegistry.getStakeAtBlockNumber(operatorId, 0, blockNumber);

            operators[i] = RollupOperators.Operator({pubkey: pubkey, weight: stake});
        }

        return operators;
    }
}
