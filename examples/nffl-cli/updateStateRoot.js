const { ethers } = require('ethers');
const {NFFLRegistryRollupABI} = require('./abi/NFFLRegistryRollup');
const { hashG1Point } = require('./src/hashG1Point');
const { createWallet } = require('./src/createWallet');
/*
 * Updates the state root.
*/
async function updateStateRoot(options) {
    // Init provider
    const provider = new ethers.JsonRpcProvider(options.rpcUrl);
    // Init wallet
    const wallet = createWallet(options.envKey);
    const account = wallet.connect(provider);
    console.log('Wallet address:', await account.getAddress());
    // Get RegistryRollup contract
    const registryRollup = new ethers.Contract(options.nfflRegistryRollup, NFFLRegistryRollupABI, account);
    // Fetch data
    console.log(`${options.aggregator}/aggregation/state-root-update?rollupId=${options.rollupId}&blockHeight=${options.blockHeight}`);
    const response = await fetch(`${options.aggregator}/aggregation/state-root-update?rollupId=${options.rollupId}&blockHeight=${options.blockHeight}`);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    // Build message
    const nearDaTransactionId = '0x'+Buffer.from(data.Message.NearDaTransactionId).toString('hex');
    const nearDaCommitment = '0x'+Buffer.from(data.Message.NearDaCommitment).toString('hex');
    const stateRoot = '0x'+Buffer.from(data.Message.StateRoot).toString('hex');
    const message = {
        rollupId: data.Message.RollupId,
        blockHeight: data.Message.BlockHeight,
        timestamp:data.Message.Timestamp,
        nearDaTransactionId,
        nearDaCommitment,
        stateRoot
    };
    // Sort non-signers keys
    const nonSignerPubkeys = data.Aggregation.NonSignersPubkeysG1;
    nonSignerPubkeys.sort((a, b) => {
        const hashA = hashG1Point(a);
        const hashB = hashG1Point(b);
        return hashA.comparedTo(hashB);
    });
    // Create signature info
    const signatureInfo =  {
        nonSignerPubkeys,
        apkG2: {
            X: [data.Aggregation.SignersApkG2.X.A1,data.Aggregation.SignersApkG2.X.A0],
            Y: [data.Aggregation.SignersApkG2.Y.A1,data.Aggregation.SignersApkG2.Y.A0]
        },
        sigma: data.Aggregation.SignersAggSigG1.g1_point
    }
    // Update state root
    const tx = await registryRollup.updateStateRoot(message,signatureInfo);
    console.log('transaction:', tx);
    await tx.wait();
    console.log('State root updated');
}

module.exports = {updateStateRoot}
