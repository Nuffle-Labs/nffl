const { ethers } = require('ethers');
const {NFFLRegistryRollupABI} = require('./abi/NFFLRegistryRollup');
const {arbContracts} = require('./contracts');
const config = require('./config.json');
const { secretSeedPhrase } = require('../secret/secret');
const { hashG1Point } = require('./src/hashG1Point');
/*
 * Updates the state root on Arbitrum to the Optimism block state.
*/
async function updateStateRoot() {
    // Init provider
    const arbProvider = new ethers.JsonRpcProvider(config.arbRpcUrl, config.arbNetworkId);
    // Init wallet
    const wallet = ethers.Wallet.fromPhrase(secretSeedPhrase);
    const account = wallet.connect(arbProvider);
    console.log('Wallet address:', await account.getAddress());
    // Get RegistryRollup contract
    const registryRollup = new ethers.Contract(arbContracts.addresses.sfflRegistryRollup, NFFLRegistryRollupABI, account);
    // Fetch data
    const blockHeight = 13905480;
    console.log(`${config.aggregator}/aggregation/state-root-update?rollupId=11155420&blockHeight=${blockHeight}`);
    const response = await fetch(`${config.aggregator}/aggregation/state-root-update?rollupId=11155420&blockHeight=${blockHeight}`);
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
    console.log(message);

    // Sort non-signers keys
    const nonSignerPubkeys = data.Aggregation.NonSignersPubkeysG1;
    nonSignerPubkeys.sort((a, b) => {
        const hashA = hashG1Point(a);
        const hashB = hashG1Point(b);
        return hashA.comparedTo(hashB); // Compare BigNumber values
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
    console.log(signatureInfo)

    // Update state root
    const tx = await registryRollup.updateStateRoot(message,signatureInfo);
    console.log('transaction:', tx);
    await tx.wait();
    console.log('State root updated');
}

updateStateRoot()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
});