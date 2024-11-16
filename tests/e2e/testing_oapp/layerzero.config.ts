import { EndpointId } from '@layerzerolabs/lz-definitions'
import type { OAppOmniGraphHardhat, OmniPointHardhat } from '@layerzerolabs/toolbox-hardhat'

import holesky from "../../../contracts/evm/broadcast/Deploy.s.sol/17000/run-latest.json";
import amoy from "../../../contracts/evm/broadcast/Deploy.s.sol/80002/run-latest.json";
import arbitrum from "../../../contracts/evm/broadcast/Deploy.s.sol/421614/run-latest.json";

import holeskyExecutor from "../../../contracts/evm/broadcast/DeployExecutor.s.sol/17000/run-latest.json";
import amoyExecutor from "../../../contracts/evm/broadcast/DeployExecutor.s.sol/80002/run-latest.json";
import arbitrumExecutor from "../../../contracts/evm/broadcast/DeployExecutor.s.sol/421614/run-latest.json";

console.log("=== LayerZero Testing OApp ===");
console.log("Setting path with `Holesky`'s DVN @", holesky.receipts[0].contractAddress);
console.log("Setting path with `Arbitrum-Sepolia`'s DVN @", arbitrum.receipts[0].contractAddress);
console.log("Setting path with `Polygon-Amoy`'s DVN @", amoy.receipts[0].contractAddress);

const arbitrumSepoliaContract: OmniPointHardhat = {
    eid: EndpointId.ARBSEP_V2_TESTNET,
    contractName: 'TestingOApp',
}

const holeskyContract: OmniPointHardhat = {
    eid: EndpointId.HOLESKY_V2_TESTNET,
    contractName: 'TestingOApp',
}

const amoyContract: OmniPointHardhat = {
    eid: EndpointId.AMOY_V2_TESTNET,
    contractName: 'TestingOApp',
}

const config: OAppOmniGraphHardhat = {
    contracts: [
        {
            contract: holeskyContract,
        },
        {
            contract: amoyContract,
        },
        {
            contract: arbitrumSepoliaContract,
        },
    ],
    connections: [
        {
            from: holeskyContract,
            to: amoyContract,
            config: {
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: "0xBc0C24E6f24eC2F1fd7E859B8322A1277F80aaD5" // holeskyExecutor.receipts[0].contractAddress,
                    },
                    ulnConfig: {
                        confirmations: BigInt(0),
                        requiredDVNs: [
                            holesky.receipts[0].contractAddress,
                        ],
                    }
                },
            }
        },
        {
            from: holeskyContract,
            to: arbitrumSepoliaContract,
            config: {
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: "0xBc0C24E6f24eC2F1fd7E859B8322A1277F80aaD5" // holeskyExecutor.receipts[0].contractAddress,
                    },
                    ulnConfig: {
                        confirmations: BigInt(0),
                        requiredDVNs: [
                            holesky.receipts[0].contractAddress,
                        ],
                    }
                },
            }
        },
        {
            from: amoyContract,
            to: holeskyContract,
            config: {
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: "0x4Cf1B3Fa61465c2c907f82fC488B43223BA0CF93" // amoyExecutor.receipts[0].contractAddress,
                    },
                    ulnConfig: {
                        confirmations: BigInt(0),
                        requiredDVNs: [
                            amoy.receipts[0].contractAddress,
                        ],
                    }
                },
            }
        },
        {
            from: amoyContract,
            to: arbitrumSepoliaContract,
            config: {
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: "0x4Cf1B3Fa61465c2c907f82fC488B43223BA0CF93" // amoyExecutor.receipts[0].contractAddress,
                    },
                    ulnConfig: {
                        confirmations: BigInt(0),
                        requiredDVNs: [
                            amoy.receipts[0].contractAddress,
                        ],
                    }
                },
            }
        },
        {
            from: arbitrumSepoliaContract,
            to: holeskyContract,
            config: {
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: "0x5Df3a1cEbBD9c8BA7F8dF51Fd632A9aef8308897" // arbitrumExecutor.receipts[0].contractAddress,
                    },
                    ulnConfig: {
                        confirmations: BigInt(0),
                        requiredDVNs: [
                            arbitrum.receipts[0].contractAddress,
                        ],
                    }
                },
            }
        },
    ],
}

export default config
