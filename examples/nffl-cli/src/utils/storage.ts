import { ethers } from "ethers";
import rlp from "rlp";

export interface GetProofResponse {
    address: string;
    accountProof: string[];
    balance: string;
    codeHash: string;
    nonce: string;
    storageHash: string;
    storageProof: {
        key: string;
        value: string;
        proof: string[];
    }[];
}

export async function getProof(provider: ethers.JsonRpcProvider, contractAddress: string, storageKey: string, blockHeight: bigint): Promise<GetProofResponse> {
    return await provider.send("eth_getProof", [
        contractAddress,
        [storageKey],
        blockHeight,
    ]);
}

export interface VerifiedStorage {
    target: string;
    storageKey: string;
    storageValue: string;
    stateTrieWitness: string;
    storageTrieWitness: string;
}

export async function verifyStorage(provider: ethers.JsonRpcProvider, contractAddress: string, storageKey: string, blockHeight: bigint): Promise<VerifiedStorage> {
    const proof = await getProof(provider, contractAddress, storageKey, blockHeight);

    return {
        target: contractAddress,
        storageKey,
        storageValue: proof.storageProof[0].value,
        stateTrieWitness: ethers.hexlify(rlp.encode(proof.accountProof)),
        storageTrieWitness: ethers.hexlify(rlp.encode(proof.storageProof[0].proof)),
    };
}
