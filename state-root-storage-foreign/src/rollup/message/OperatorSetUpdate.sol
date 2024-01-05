// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Operators} from "../utils/Operators.sol";

library OperatorSetUpdate {
    struct Message {
        Operators.Operator[] operators;
    }

    function hashCalldata(Message calldata message) internal pure returns (bytes32) {
        return keccak256(abi.encode(message));
    }
}
