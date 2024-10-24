import { ethers } from "ethers";
import { SFFLRegistryRollup__factory } from "../../typechain-types/factories/SFFLRegistryRollup__factory";
import { fetchOperatorSetUpdate } from "../utils/api";

export interface UpdateOperatorSetOptions {
    rpcUrl: string;
    contractAddress: string;
    id: string;
    aggregatorUrl: string;
    privateKey: string;
}

export async function updateOperatorSet(options: UpdateOperatorSetOptions) {
    const provider = new ethers.JsonRpcProvider(options.rpcUrl);

    const wallet = new ethers.Wallet(options.privateKey);
    const account = wallet.connect(provider);

    const registryRollup = SFFLRegistryRollup__factory.connect(options.contractAddress, account);

    const operatorSetUpdate = await fetchOperatorSetUpdate(options.aggregatorUrl, BigInt(options.id));

    const tx = await registryRollup.updateOperatorSet(operatorSetUpdate.message, operatorSetUpdate.aggregation);
    await tx.wait();
}