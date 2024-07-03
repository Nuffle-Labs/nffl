const { ethers } = require('ethers');

// Function to get private key from evironment variable
function getPrivateKey(envKey) {
    if (envKey && process.env[envKey]) {
        return process.env[envKey];
    }
    console.error('Error: Private key is not provided.')``;
    process.exit(1);
}

// Function to create a wallet from a evironment variable containing the private key
function createWallet(envKey) {
    const privateKey = getPrivateKey(envKey);
    try {
        const wallet = new ethers.Wallet(privateKey);
        return wallet;
    } catch (error) {
        console.error('Invalid private key.');
        process.exit(1);
    }
}

module.exports = {createWallet}
