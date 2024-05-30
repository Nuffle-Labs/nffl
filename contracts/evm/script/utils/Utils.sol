// SPDX-License-Identifier: MIT
pragma solidity =0.8.12;

import {ProxyAdmin, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract Utils is Script {
    function readInput(string memory inputFileName) internal view returns (string memory) {
        string memory file = getInputPath(inputFileName);
        return vm.readFile(file);
    }

    function getInputPath(string memory inputFileName) internal view returns (string memory) {
        string memory inputDir = string.concat(vm.projectRoot(), "/script/input/");
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory file = string.concat(inputFileName, ".json");
        return string.concat(inputDir, chainDir, file);
    }

    function readOutput(string memory outputFileName) internal view returns (string memory) {
        string memory file = getOutputPath(outputFileName);
        return vm.readFile(file);
    }

    function writeOutput(string memory outputJson, string memory outputFileName) internal {
        string memory outputFilePath = getOutputPath(outputFileName);
        vm.writeJson(outputJson, outputFilePath);
    }

    function getOutputPath(string memory outputFileName) internal view returns (string memory) {
        string memory outputDir = string.concat(vm.projectRoot(), "/script/output/");
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory outputFilePath = string.concat(outputDir, chainDir, outputFileName, ".json");
        return outputFilePath;
    }

    /**
     * @dev Deploys a new proxy contract using the given implementation and initialization data.
     * @param _impl Address of the implementation contract.
     * @param _admin Proxy admin.
     * @param _initCode Initialization code.
     */
    function _deployProxy(ProxyAdmin _admin, address _impl, bytes memory _initCode)
        internal
        returns (TransparentUpgradeableProxy)
    {
        return new TransparentUpgradeableProxy(_impl, address(_admin), _initCode);
    }

    /**
     * @dev Deploys an empty proxy - i.e. a zero implementation and with no init code
     * @param _admin Proxy admin.
     */
    function _deployEmptyProxy(ProxyAdmin _admin, address emptyContract)
        internal
        returns (TransparentUpgradeableProxy)
    {
        return new TransparentUpgradeableProxy(emptyContract, address(_admin), "");
    }

    /**
     * @dev Upgrades a proxy to a new implementation.
     * @param _admin Proxy admin.
     * @param _proxy The proxy to upgrade.
     * @param _impl The new implementation to upgrade to.
     */
    function _upgradeProxy(ProxyAdmin _admin, TransparentUpgradeableProxy _proxy, address _impl) internal {
        _admin.upgrade(_proxy, _impl);
    }

    /**
     * @dev Upgrades a proxy to a new impl and calls a function on the implementation.
     * @param _admin Proxy admin.
     * @param _proxy The proxy to upgrade.
     * @param _impl The new impl to upgrade to.
     * @param _data The encoded calldata to use in the call after upgrading.
     */
    function _upgradeProxyAndCall(
        ProxyAdmin _admin,
        TransparentUpgradeableProxy _proxy,
        address _impl,
        bytes memory _data
    ) internal {
        _admin.upgradeAndCall(_proxy, _impl, _data);
    }
}
