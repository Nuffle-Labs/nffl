const { ethers } = require('ethers');
const {NFFLRegistryRollupABI} = require('./abi/NFFLRegistryRollup');
const RLP = require('rlp');

/*
 * Get the storage value of a target contract at a specified block height  
 * from the source chain and verifies it on the destination chain
*/
async function getStorageValue(options) {
    //
    // Get proof
    //
    const srcProvider = new ethers.JsonRpcProvider(options.srcRpcUrl);
    //chainId
    const { chainId } = await srcProvider.getNetwork();
    // Prepare params
    const params = [
        options.contractAddress,
        [options.storageKey],
        `0x${Number(options.blockHeight).toString(16)}`
    ];
    // Send the RPC request
    const proof = await srcProvider.send("eth_getProof", params);
    // Encode proof to RLP
    const rlpStorageProof = RLP.encode(proof.storageProof[0].proof);
    const rlpAccountProof = RLP.encode(proof.accountProof);
    
    //
    // Prove
    //
    // Get RegistryRollup contract
    const dstProvider = new ethers.JsonRpcProvider(options.dstRpcUrl);
    const registryRollup = new ethers.Contract(options.nfflRegistryRollup, NFFLRegistryRollupABI, dstProvider);
    // Fetch data
    const response = await fetch(`${options.aggregator}/aggregation/state-root-update?rollupId=${chainId}&blockHeight=${options.blockHeight}`);
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
        target: options.contractAddress,
        storageKey: options.storageKey,
        stateTrieWitness: rlpAccountProof,
        storageTrieWitness: rlpStorageProof
    }

    // Get storage value
    const storageValue = await registryRollup.getStorageValue(message,proofParams);
    console.log(`Account ${options.contractAddress} storage slot ${options.storageKey} equals to ${storageValue}`);
}

module.exports = {getStorageValue}