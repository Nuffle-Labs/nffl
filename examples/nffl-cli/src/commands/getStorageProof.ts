import { ethers } from "ethers";
import { verifyStorage } from "../utils/storage";

export interface GetStorageProofOptions {
    rpcUrl: string;
    contractAddress: string;
    storageKey: string;
}

export async function getStorageProof(options: GetStorageProofOptions) {
    const provider = new ethers.JsonRpcProvider(options.rpcUrl);

    const verifiedStorage = await verifyStorage(provider, options.contractAddress, options.storageKey);

    console.log(verifiedStorage);
}