
const { ethers, keccak256 } = require('ethers');
const {NFFLRegistryRollupABI} = require('./abi/NFFLRegistryRollup');
const {arbContracts} = require('./contracts');
const config = require('./config.json');
const { secretSeedPhrase } = require('../secret/secret');
const BigNumber = require('bignumber.js');

function hashG1Point(pk) {
    const xValue = BigNumber(pk.X);
    const yValue = BigNumber(pk.Y);
    const buf = Buffer.alloc(64);

    // Store in buf
    buf.write(xValue.toString(16).padStart(64, '0'), 0, 'hex'); 
    buf.write(yValue.toString(16).padStart(64, '0'), 32, 'hex'); 
    // Calculate keccak256 hash of buf
    const hashedG1 = keccak256(buf);

    return BigNumber(hashedG1);
}
async function updateOperatorSet() {
    // Init provider
    const arbProvider = new ethers.JsonRpcProvider(config.arbRpcUrl, config.arbNetworkId);
    // Init wallet
    const wallet = ethers.Wallet.fromPhrase(secretSeedPhrase);
    const account = wallet.connect(arbProvider);
    console.log('Wallet address:', await account.getAddress());
    // Get next operator update id
    const registryRollup = new ethers.Contract(arbContracts.addresses.sfflRegistryRollup, NFFLRegistryRollupABI, account);
    const nextOperatorUpdateId = await registryRollup.nextOperatorUpdateId();
    console.log('nextOperatorUpdateId',nextOperatorUpdateId);
    // Fetch data
    console.log(`${config.aggregator}/aggregation/operator-set-update?id=${nextOperatorUpdateId}`);
    const response = await fetch(`${config.aggregator}/aggregation/operator-set-update?id=${nextOperatorUpdateId}`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const respText = await response.text();
    text = respText.replace(/"Weight":\s*(\d+)/g, '"Weight": "$1"');
    data = JSON.parse(text);
    const operators = data.Message.Operators.map(({ Pubkey, Weight }) => ({
        pubkey: Pubkey,
        weight: Weight
    }));
    const message = {
        id: data.Message.Id,
        timestamp: data.Message.Timestamp,
        operators
    }
    console.log('New operators:',message.operators);
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
    // Call contact
    const tx = await registryRollup.updateOperatorSet(message,signatureInfo);
    console.log('transaction:', tx);
    await tx.wait();
    // Get next operator update id
    console.log('new nextOperatorUpdateId',await registryRollup.nextOperatorUpdateId());
}

updateOperatorSet()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
});