const { keccak256 } = require('ethers');
const BigNumber = require('bignumber.js');
/*
 * Hashes a G1 point
*/
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

module.exports = {hashG1Point}
