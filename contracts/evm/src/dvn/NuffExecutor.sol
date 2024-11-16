// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import { ILayerZeroExecutor } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/interfaces/ILayerZeroExecutor.sol";

contract NuffExecutor is ILayerZeroExecutor {
   function assignJob(
       uint32 _dstEid,
       address _sender,
       uint256 _calldataSize,
       bytes calldata _options
   ) external returns (uint256 price) {
       price = 10000;
   }

   function getFee(
       uint32 _dstEid,
       address _sender,
       uint256 _calldataSize,
       bytes calldata _options
   ) external view returns (uint256 price) {
       price = 10000;
   }

}

