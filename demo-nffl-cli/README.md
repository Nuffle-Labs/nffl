## How to build
```sh
npm install
npm link
```

## How to run
### getStorageValue
Get the storage value of a target contract at a specified block height from the source chain and verifies it on the destination chain
```sh
demo-nffl-cli getStorageValue\
 --srcRpcUrl https://sepolia.optimism.io\
 --contractAddress 0xB90101779CC5EB84162f72A80e44307752b778b6\
 --storageKey 0x0000000000000000000000000000000000000000000000000000000000000000\
 --blockHeight 13905480\
 --dstRpcUrl https://sepolia-rollup.arbitrum.io/rpc\
 --nfflRegistryRollup 0x23e252b4Ec7cDd3ED84f039EF53DEa494CE878E0\
 --aggregator http://127.0.0.1:4002
 ```

 ### updateStateRoot
 Update the state root for a given rollup in the nfflRegistryRollup contract
```sh
demo-nffl-cli updateStateRoot\
 --rpcUrl https://sepolia-rollup.arbitrum.io/rpc\
 --rollupId 11155420\
 --blockHeight 14095733\
 --nfflRegistryRollup 0x23e252b4Ec7cDd3ED84f039EF53DEa494CE878E0\
 --aggregator http://127.0.0.1:4002\
 --envKey PRIVATE_KEY
```

### updateOperatorSet
Update the operator set for the nfflRegistryRollup contract
```sh
demo-nffl-cli updateOperatorSet\
 --rpcUrl https://sepolia-rollup.arbitrum.io/rpc\
 --nfflRegistryRollup 0x23e252b4Ec7cDd3ED84f039EF53DEa494CE878E0\
 --aggregator http://127.0.0.1:4002\
 --envKey PRIVATE_KEY
```