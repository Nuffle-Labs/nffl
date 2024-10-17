// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

struct Verification {
    bool submitted;
    uint64 confirmations;
}

struct UlnConfig {
    uint64 confirmations;
    uint8 requiredDVNCount;
    uint8 optionalDVNCount;
    uint8 optionalDVNThreshold;
    address[] requiredDVNs;
    address[] optionalDVNs;
}

/// @dev should be implemented by the ReceiveUln302 contract and future ReceiveUln contracts on EndpointV2
interface IReceiveUlnE2 {
    /// @notice for each dvn to verify the payload
    function verify(
        bytes calldata _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations
    ) external;

    /// @notice verify the payload at endpoint, will check if all DVNs verified
    function commitVerification(
        bytes calldata _packetHeader,
        bytes32 _payloadHash
    ) external;

    function hashLookup(
        bytes32 _headerHash,
        bytes32 _payloadHash,
        address _dnv
    ) external view returns (Verification memory);

    function getUlnConfig(
        address _oapp,
        uint32 _remoteEid
    ) external view returns (UlnConfig memory rtnConfig);
}
