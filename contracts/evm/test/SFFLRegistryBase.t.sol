// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";

import {SFFLRegistryBase} from "../src/base/SFFLRegistryBase.sol";
import {StateRootUpdate} from "../src/base/message/StateRootUpdate.sol";

contract SFFLRegistryBaseHarness is SFFLRegistryBase {
    using StateRootUpdate for StateRootUpdate.Message;

    bool public shouldDenyAgreement;

    function setShouldDenyAgreement(bool _shouldDenyAgreement) external {
        shouldDenyAgreement = _shouldDenyAgreement;
    }

    function pushStateRoot(uint32 rollupId, uint64 blockHeight, bytes32 stateRoot) external {
        _pushStateRoot(rollupId, blockHeight, stateRoot);
    }

    function _updateStateRoot(StateRootUpdate.Message calldata message, bytes calldata) internal override {
        require(!shouldDenyAgreement, "Agreement denied");

        _stateRootBuffers[message.rollupId][message.blockHeight] = message.stateRoot;
    }
}

contract SFFLRegistryBaseTest is Test {
    SFFLRegistryBaseHarness public registry;

    event StateRootUpdated(uint32 indexed rollupId, uint64 indexed blockHeight, bytes32 stateRoot);

    function setUp() public {
        registry = new SFFLRegistryBaseHarness();
    }

    function test_getStateRoot_ReturnZeroOnEmpty() public {
        assertEq(registry.getStateRoot(0, 0), bytes32(0));
    }

    function test_getStateRoot_ReturnStoredStateRoot() public {
        assertEq(registry.getStateRoot(0, 0), bytes32(0));

        registry.pushStateRoot(0, 0, keccak256(hex"def1"));

        assertEq(registry.getStateRoot(0, 0), keccak256(hex"def1"));
    }

    function test_pushStateRoot_EmitStateRootUpdatedEvent() public {
        vm.expectEmit(true, true, false, true);
        emit StateRootUpdated(1, 2, keccak256(hex"beef"));

        registry.pushStateRoot(1, 2, keccak256(hex"beef"));
    }

    function test_pushStateRoot_StoreStateRoot() public {
        registry.pushStateRoot(0, 0, keccak256(hex"def1"));

        assertEq(registry.getStateRoot(0, 0), keccak256(hex"def1"));
    }

    function test_pushStateRoot_OverwriteStateRoot() public {
        registry.pushStateRoot(0, 0, keccak256(hex"def1"));

        assertEq(registry.getStateRoot(0, 0), keccak256(hex"def1"));

        registry.pushStateRoot(0, 0, keccak256(hex"beef"));

        assertEq(registry.getStateRoot(0, 0), keccak256(hex"beef"));
    }

    function test_getStorageValue() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0x0123456789012345678901234567890123456789,
            storageKey: 0xced0071642172bcb312d265ce65b397425c906a7575a4912408cf6c3a3265eb1,
            stateTrieWitness: hex"f86eb86cf86aa1200e61e9d7b7f8a76da4339d2962273cc6bee0df97274cb94e5b05588afe2b3a50b846f8448080a089cf4cf2ddd661535a4e28b774fec12c8fcb6ba78bc55946b00a9ab5e99c36e9a056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f83cb83af838a12008fcd933278b7ead42429ac19785950f9813a683a64c01a93405753905661aca95940000000000000000000000000000000000000001"
        });

        bytes32 stateRoot = 0x52fd73f9175ec160ff5fbf32a985447b4b95b87ac6b1860bf9a1fb81f954d774;
        bytes32 expectedValue = 0x0000000000000000000000000000000000000000000000000000000000000001;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);
        registry.pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);

        assertEq(registry.getStorageValue(message, proofParams), expectedValue);
    }

    function test_getStorageValue_RevertWhen_TargetDoesNotExist() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa,
            storageKey: 0xced0071642172bcb312d265ce65b397425c906a7575a4912408cf6c3a3265eb1,
            stateTrieWitness: hex"f86eb86cf86aa1200e61e9d7b7f8a76da4339d2962273cc6bee0df97274cb94e5b05588afe2b3a50b846f8448080a089cf4cf2ddd661535a4e28b774fec12c8fcb6ba78bc55946b00a9ab5e99c36e9a056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f83cb83af838a12008fcd933278b7ead42429ac19785950f9813a683a64c01a93405753905661aca95940000000000000000000000000000000000000001"
        });

        bytes32 stateRoot = 0x52fd73f9175ec160ff5fbf32a985447b4b95b87ac6b1860bf9a1fb81f954d774;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);
        registry.pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);

        vm.expectRevert("Account does not exist");
        registry.getStorageValue(message, proofParams);
    }

    function test_getStorageValue_RevertWhen_StorageValueDoesNotExist() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0x0123456789012345678901234567890123456789,
            storageKey: 0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb,
            stateTrieWitness: hex"f86eb86cf86aa1200e61e9d7b7f8a76da4339d2962273cc6bee0df97274cb94e5b05588afe2b3a50b846f8448080a089cf4cf2ddd661535a4e28b774fec12c8fcb6ba78bc55946b00a9ab5e99c36e9a056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f83cb83af838a12008fcd933278b7ead42429ac19785950f9813a683a64c01a93405753905661aca95940000000000000000000000000000000000000001"
        });

        bytes32 stateRoot = 0x52fd73f9175ec160ff5fbf32a985447b4b95b87ac6b1860bf9a1fb81f954d774;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);
        registry.pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);

        vm.expectRevert("Storage value does not exist");
        registry.getStorageValue(message, proofParams);
    }

    function test_getStorageValue_RevertWhen_MismatchingStateRoots() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0x0123456789012345678901234567890123456789,
            storageKey: 0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb,
            stateTrieWitness: hex"f86eb86cf86aa1200e61e9d7b7f8a76da4339d2962273cc6bee0df97274cb94e5b05588afe2b3a50b846f8448080a089cf4cf2ddd661535a4e28b774fec12c8fcb6ba78bc55946b00a9ab5e99c36e9a056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f83cb83af838a12008fcd933278b7ead42429ac19785950f9813a683a64c01a93405753905661aca95940000000000000000000000000000000000000001"
        });

        bytes32 stateRoot = 0x52fd73f9175ec160ff5fbf32a985447b4b95b87ac6b1860bf9a1fb81f954d774;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);

        vm.expectRevert("Mismatching state roots");
        registry.getStorageValue(message, proofParams);
    }

    function test_updateAndGetStorageValue_NoUpdate() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE,
            storageKey: 0x470ebe1a3f7c174ece10a895b1c5597999ee280a8a7afaae776da0692fc28e7b,
            stateTrieWitness: hex"f86eb86cf86aa1209f74bd52020a869dbd6c5918e246e54fe47bed2b9e96439c406e5c0732d089bfb846f8448080a01efe59d3e576d132e64cb197fc20e8cb2ed260308c189f2a0bb14843eb126b1ca056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f848b846f844a1206cda01cb275318e1eb18c3a672f0185922be53b3d63d6a92fbf81420cf1a2783a1a0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
        });

        bytes32 stateRoot = 0xd223489d5fdd65f1fb9beb4dd16a35540d34993f68409f8033a27e8115020f15;
        bytes32 expectedValue = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);
        registry.pushStateRoot(message.rollupId, message.blockHeight, message.stateRoot);

        registry.setShouldDenyAgreement(true);

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), stateRoot);

        assertEq(registry.updateAndGetStorageValue(message, proofParams, hex""), expectedValue);

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), stateRoot);
    }

    function test_updateAndGetStorageValue_Update() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE,
            storageKey: 0x470ebe1a3f7c174ece10a895b1c5597999ee280a8a7afaae776da0692fc28e7b,
            stateTrieWitness: hex"f86eb86cf86aa1209f74bd52020a869dbd6c5918e246e54fe47bed2b9e96439c406e5c0732d089bfb846f8448080a01efe59d3e576d132e64cb197fc20e8cb2ed260308c189f2a0bb14843eb126b1ca056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f848b846f844a1206cda01cb275318e1eb18c3a672f0185922be53b3d63d6a92fbf81420cf1a2783a1a0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
        });

        bytes32 stateRoot = 0xd223489d5fdd65f1fb9beb4dd16a35540d34993f68409f8033a27e8115020f15;
        bytes32 expectedValue = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);

        registry.setShouldDenyAgreement(false);

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), bytes32(0));

        assertEq(registry.updateAndGetStorageValue(message, proofParams, hex"ff"), expectedValue);

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), stateRoot);
    }

    function test_updateAndGetStorageValue_RevertWhen_UpdateAndEmptyAgreement() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE,
            storageKey: 0x470ebe1a3f7c174ece10a895b1c5597999ee280a8a7afaae776da0692fc28e7b,
            stateTrieWitness: hex"f86eb86cf86aa1209f74bd52020a869dbd6c5918e246e54fe47bed2b9e96439c406e5c0732d089bfb846f8448080a01efe59d3e576d132e64cb197fc20e8cb2ed260308c189f2a0bb14843eb126b1ca056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f848b846f844a1206cda01cb275318e1eb18c3a672f0185922be53b3d63d6a92fbf81420cf1a2783a1a0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
        });

        bytes32 stateRoot = 0xd223489d5fdd65f1fb9beb4dd16a35540d34993f68409f8033a27e8115020f15;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);

        registry.setShouldDenyAgreement(false);

        vm.expectRevert("Empty agreement");
        registry.updateAndGetStorageValue(message, proofParams, hex"");
    }

    function test_updateAndGetStorageValue_RevertWhen_AgreementDenied() public {
        SFFLRegistryBase.ProofParams memory proofParams = SFFLRegistryBase.ProofParams({
            target: 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE,
            storageKey: 0x470ebe1a3f7c174ece10a895b1c5597999ee280a8a7afaae776da0692fc28e7b,
            stateTrieWitness: hex"f86eb86cf86aa1209f74bd52020a869dbd6c5918e246e54fe47bed2b9e96439c406e5c0732d089bfb846f8448080a01efe59d3e576d132e64cb197fc20e8cb2ed260308c189f2a0bb14843eb126b1ca056570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
            storageTrieWitness: hex"f848b846f844a1206cda01cb275318e1eb18c3a672f0185922be53b3d63d6a92fbf81420cf1a2783a1a0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
        });

        bytes32 stateRoot = 0xd223489d5fdd65f1fb9beb4dd16a35540d34993f68409f8033a27e8115020f15;

        StateRootUpdate.Message memory message = StateRootUpdate.Message(1, 2, 3, stateRoot);

        registry.setShouldDenyAgreement(true);

        vm.expectRevert("Agreement denied");
        registry.updateAndGetStorageValue(message, proofParams, hex"ff");
    }
}
