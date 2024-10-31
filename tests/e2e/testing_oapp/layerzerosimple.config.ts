import { EndpointId } from '@layerzerolabs/lz-definitions'
import type { OAppOmniGraphHardhat, OmniPointHardhat } from '@layerzerolabs/toolbox-hardhat'

import holesky from "../../../contracts/evm/broadcast/DeploySimple.s.sol/17000/run-latest.json";
import amoy from "../../../contracts/evm/broadcast/DeploySimple.s.sol/80002/run-latest.json";
import arbitrum from "../../../contracts/evm/broadcast/DeploySimple.s.sol/421614/run-latest.json";

console.log("=== LayerZero Simple Testing OApp ===");
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
                    ulnConfig: {
                        confirmations: BigInt(1),
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
                    ulnConfig: {
                        confirmations: BigInt(1),
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
                    ulnConfig: {
                        confirmations: BigInt(1),
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
                    ulnConfig: {
                        confirmations: BigInt(1),
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
                    ulnConfig: {
                        confirmations: BigInt(1),
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
