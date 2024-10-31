// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import { ILayerZeroDVN } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/ILayerZeroDVN.sol";

contract SimpleNuffDVN is ILayerZeroDVN {
    event VerifierFeePaid(uint256 fee);

    constructor() { }

    function assignJob(
        AssignJobParam calldata _param,
        bytes calldata _options
    )
        external
        payable
        override
        returns (uint256 fee)
    {
        fee = 10000;
        emit VerifierFeePaid(fee);
    }

    function getFee(
        uint32 _dstEid,
        uint64 _confirmations,
        address _sender,
        bytes calldata _options
    ) external view override returns (uint256 fee) {
        fee = 10000;
    }
}
