// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import { ILayerZeroEndpointV2 } from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import { ISendLib } from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ISendLib.sol";
import { IDVN } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IDVN.sol";
import { ILayerZeroDVN } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/ILayerZeroDVN.sol";
import { PacketV1Codec } from "@layerzerolabs/lz-evm-protocol-v2/contracts/messagelib/libs/PacketV1Codec.sol";
import { IReceiveUlnE2 } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IReceiveUlnE2.sol";
import { IDVNFeeLib } from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/interfaces/IDVNFeeLib.sol";

import { INuffClient } from "./interfaces/INuffClient.sol";
import { INuffDVNConfig } from "./interfaces/INuffDVNConfig.sol";

import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";

abstract contract NuffDVN is ILayerZeroDVN, AccessControl, IDVN, ReentrancyGuard {
    using PacketV1Codec for bytes;
    using ECDSA for bytes32;

    struct Job {
        address origin;
        uint32 srcEid;
        uint32 dstEid;
        bytes packetHeader;
        bytes32 payloadHash;
        uint64 confirmations;
        address sender;
        address receiver;
        bytes options;
    }

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant MESSAGE_LIB_ROLE = keccak256("MESSAGE_LIB_ROLE");

    ILayerZeroEndpointV2 public layerZeroEndpointV2;
    uint32 public immutable localEid;

    uint256 public lastJobId;

    uint256 public nuffAppId;
    INuffClient.PublicKey public nuffPublicKey;
    INuffClient public nuff;
    INuffDVNConfig public dvnConfig;

    uint16 public defaultMultiplierBps;
    uint64 public quorum;
    address public priceFeed;
    address public feeLib;

    // FIXME: everything is getting stored in cold storage; use a buffer instead
    mapping(uint256 jobId => Job job) public jobs;
    mapping(uint32 eid => bool isSupported) public supportedDstChain;
    mapping(uint32 dstEid => DstConfig config) public dstConfig;
    mapping(uint32 srcEid => mapping(uint256 jobId => bool isVerified)) public verifiedJobs;

    event JobAssigned(uint256 jobId);
    event Verified(uint32 srcEid, uint256 jobId);

    constructor(
        uint256 _nuffAppId,
        INuffClient.PublicKey memory _nuffPublicKey,
        address _nuff,
        address _layerZeroEndpointV2,
        address _dvnConfig,
        uint16 _defaultMultiplierBps,
        uint64 _quorum,
        address _priceFeed,
        address _feeLib
    ) {
        nuffAppId = _nuffAppId;
        nuffPublicKey = _nuffPublicKey;
        nuff = INuffClient(_nuff);
        layerZeroEndpointV2 = ILayerZeroEndpointV2(_layerZeroEndpointV2);
        dvnConfig = INuffDVNConfig(_dvnConfig);
        localEid = layerZeroEndpointV2.eid();
        defaultMultiplierBps = _defaultMultiplierBps;
        quorum = _quorum;
        priceFeed = _priceFeed;
        feeLib = _feeLib;
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function assignJob(
        AssignJobParam calldata _param,
        bytes calldata _options
    )
        external
        nonReentrant
        payable
        override
        onlyRole(MESSAGE_LIB_ROLE)
        returns (uint256 fee)
    {
        require(supportedDstChain[_param.dstEid], "Unsupported chain");

        uint256 jobId = ++lastJobId;
        Job storage newJob = jobs[jobId];

        require(_param.sender != address(0), "Invalid sender address");

        newJob.origin = msg.sender;
        newJob.srcEid = localEid;
        newJob.dstEid = _param.dstEid;
        newJob.packetHeader = _param.packetHeader;
        newJob.payloadHash = _param.payloadHash;
        newJob.confirmations = _param.confirmations;
        newJob.sender = _param.sender;
        newJob.receiver = address(
            uint160(uint256(_param.packetHeader.receiver()))
        );
        newJob.options = _options;

        IDVNFeeLib.FeeParams memory feeParams = IDVNFeeLib.FeeParams(
            priceFeed,
            _param.dstEid,
            _param.confirmations,
            _param.sender,
            quorum,
            defaultMultiplierBps
        );

        fee = IDVNFeeLib(feeLib).getFeeOnSend(
            feeParams,
            dstConfig[_param.dstEid],
            _options
        );

        emit JobAssigned(jobId);
    }

    function verify(
        uint32 _srcEid,
        uint32 _dstEid,
        uint256 _jobId,
        bytes memory _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        address _receiver,
        bytes calldata _reqId,
        INuffClient.BLSSign calldata _signature
    ) external nonReentrant {
        require(_isLocal(_dstEid), "Invalid dstEid");
        require(
            !verifiedJobs[_srcEid][_jobId],
            "src jobId is already verified"
        );

        verifiedJobs[_srcEid][_jobId] = true;

        bytes32 hash = keccak256(
            abi.encodePacked(
                nuffAppId,
                _reqId,
                _srcEid,
                _dstEid,
                _jobId,
                _packetHeader,
                _payloadHash,
                _confirmations,
                _receiver
            )
        );

        _verifyNuffSig(
            _reqId,
            hash,
            _signature
        );

        _lzVerify(
            _srcEid,
            _packetHeader,
            _payloadHash,
            _confirmations,
            _receiver
        );

        emit Verified(_srcEid, _jobId);
    }

    function setNuffAppId(uint256 _nuffAppId) external onlyRole(ADMIN_ROLE) {
        nuffAppId = _nuffAppId;
    }

    function setNuffContract(address addr) external onlyRole(ADMIN_ROLE) {
        nuff = INuffClient(addr);
    }

    function setNuffPubKey(
        INuffClient.PublicKey memory _nuffPublicKey
    ) external onlyRole(ADMIN_ROLE) {
        nuffPublicKey = _nuffPublicKey;
    }

    function setLzEndpointV2(
        address _layerZeroEndpointV2
    ) external onlyRole(ADMIN_ROLE) {
        layerZeroEndpointV2 = ILayerZeroEndpointV2(_layerZeroEndpointV2);
    }

    function updateSupportedDstChain(
        uint32 eid,
        bool status
    ) external onlyRole(ADMIN_ROLE) {
        supportedDstChain[eid] = status;
    }

    function setPriceFeed(address _priceFeed) external onlyRole(ADMIN_ROLE) {
        priceFeed = _priceFeed;
    }

    function setDefaultMultiplierBps(
        uint16 _multiplierBps
    ) external onlyRole(ADMIN_ROLE) {
        defaultMultiplierBps = _multiplierBps;
    }

    function setDstConfig(
        DstConfigParam[] calldata _params
    ) external onlyRole(ADMIN_ROLE) {
        for (uint256 i = 0; i < _params.length; ++i) {
            DstConfigParam calldata param = _params[i];
            dstConfig[param.dstEid] = DstConfig(
                param.gas,
                param.multiplierBps,
                param.floorMarginUSD
            );
        }
        emit SetDstConfig(_params);
    }

    function setFeeLib(address _feeLib) external onlyRole(ADMIN_ROLE) {
        feeLib = _feeLib;
    }

    function withdrawFee(
        address _lib,
        address _to,
        uint256 _amount
    ) external onlyRole(ADMIN_ROLE) {
        require(hasRole(MESSAGE_LIB_ROLE, _lib), "Invalid lib");
        ISendLib(_lib).withdrawFee(_to, _amount);
        emit Withdraw(_lib, _to, _amount);
    }

    function getFee(
        uint32 _dstEid,
        uint64 _confirmations,
        address _sender,
        bytes calldata _options
    ) external view override returns (uint256 _fee) {
        IDVNFeeLib.FeeParams memory params = IDVNFeeLib.FeeParams(
            priceFeed,
            _dstEid,
            _confirmations,
            _sender,
            quorum,
            defaultMultiplierBps
        );
        return IDVNFeeLib(feeLib).getFee(params, dstConfig[_dstEid], _options);
    }

    function _verifyNuffSig(
        bytes calldata reqId,
        bytes32 hash,
        INuffClient.BLSSign calldata sign
    ) internal nonReentrant {
        bool verified = nuff.nuffVerify(
            reqId,
            uint256(hash),
            sign,
            nuffPublicKey
        );
        require(verified, "Invalid signature!");
    }

    function _lzVerify(
        uint32 _srcEid,
        bytes memory _packetHeader,
        bytes32 _payloadHash,
        uint64 _confirmations,
        address _receiver
    ) internal nonReentrant {
        address receiverLib;
        if (_isV2(_srcEid)) {
            (receiverLib, ) = layerZeroEndpointV2.getReceiveLibrary(
                _receiver,
                _srcEid
            );
        }

        IReceiveUlnE2(receiverLib).verify(
            _packetHeader,
            _payloadHash,
            _confirmations
        );
    }

    function _isLocal(uint32 _dstEid) internal view returns (bool) {
        if (localEid == _dstEid || localEid == _dstEid + 30000) {
            return true;
        }
        return false;
    }

    function _isV2(uint32 _eid) internal pure returns (bool) {
        if (_eid > 30000) {
            return true;
        }
        return false;
    }
}
