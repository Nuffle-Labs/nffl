import { ethers } from "ethers";
import { SFFLRegistryRollup__factory } from "../../typechain-types/factories/SFFLRegistryRollup__factory";
import { fetchStateRootUpdate } from "../utils/api";

export interface UpdateStateRootOptions {
    rpcUrl: string;
    contractAddress: string;
    rollupId: string;
    blockHeight: string;
    aggregatorUrl: string;
    privateKey: string;
}

export async function updateStateRoot(options: UpdateStateRootOptions) {
    const provider = new ethers.JsonRpcProvider(options.rpcUrl);

    const wallet = new ethers.Wallet(options.privateKey);
    const account = wallet.connect(provider);

    const registryRollup = SFFLRegistryRollup__factory.connect(options.contractAddress, account);

    const stateRootUpdate = await fetchStateRootUpdate(options.aggregatorUrl, BigInt(options.rollupId), BigInt(options.blockHeight));

    const tx = await registryRollup.updateStateRoot(stateRootUpdate.message, stateRootUpdate.aggregation);
    await tx.wait();
}