const { ethers } = require('ethers');
const {NFFLRegistryRollupABI} = require('./abi/NFFLRegistryRollup');
const {arbContracts} = require('./contracts');
const config = require('./config.json');
const RLP = require('rlp');

/*
 * Gets sttorage slot from Optimism.
*/
const contractAddress = '0xB90101779CC5EB84162f72A80e44307752b778b6';
const storageKey = '0x0000000000000000000000000000000000000000000000000000000000000000';
const blockNumber = '0xd42e48';//'13905480';

async function getStorageValue() {
    //
    // Get proof on Optimisp
    //
    const opProvider = new ethers.JsonRpcProvider(config.opRpcUrl, config.opNetworkId);
    // Prepear params
    const params = [
        contractAddress,
        [storageKey],
        blockNumber
    ];
    // Send the RPC request
    const proof = await opProvider.send("eth_getProof", params);
    // Encode proof to RLP
    const rlpStorageProof = RLP.encode(proof.storageProof[0].proof);
    const rlpAccountProof = RLP.encode(proof.accountProof);
    
    //
    // Prove on Arbitrum
    //
    // Get RegistryRollup contract on Arbitrum
    const arbProvider = new ethers.JsonRpcProvider(config.arbRpcUrl, config.arbNetworkId);
    const registryRollup = new ethers.Contract(arbContracts.addresses.sfflRegistryRollup, NFFLRegistryRollupABI, arbProvider);
    // Convert block number from hex
    const blockHeight = parseInt(blockNumber, 16);
    // Fetch data
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
    // Build proof parameters
    const proofParams = {
        target: contractAddress,
        storageKey: storageKey,
        stateTrieWitness: rlpAccountProof,
        storageTrieWitness: rlpStorageProof
    }

    // Get storage value
    const tx = await registryRollup.getStorageValue(message,proofParams);
    console.log(`Account ${contractAddress} storage slot ${storageKey} equals to ${tx}`);
}

getStorageValue()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
});