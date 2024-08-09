#!/usr/bin/env node

const { program } = require('commander');

const pkg = require('./package.json');

const {getStorageValue} = require('./getStorageValue');
const {updateStateRoot} = require('./updateStateRoot');
const {updateOperatorSet} = require('./updateOperatorSet');

program
  .version(pkg.version)
  .description('NFFL demo CLI tool');

program
.command('getStorageValue')
.description('Get the storage value of a target contract at a specified block height from the source chain and verifies it on the destination chain') 
.requiredOption('--srcRpcUrl <type>', 'Source RPC URL')
.requiredOption('--contractAddress <type>', 'Address of the target contract')
.requiredOption('--storageKey <type>', 'Storage key')
.requiredOption('--blockHeight <type>', 'Block height')
.requiredOption('--dstRpcUrl <type>', 'Destination RPC URL')
.requiredOption('--nfflRegistryRollup <type>', 'nfflRegistryRollup contract address on destination chain')
.requiredOption('--aggregator <type>', 'Aggregator REST API')
.action(getStorageValue);

program
.command('updateStateRoot')
.description('Update the state root for a given rollup in the nfflRegistryRollup contract')
.requiredOption('--rpcUrl <type>', 'RPC URL')
.requiredOption('--rollupId <type>', 'Rollup Id')
.requiredOption('--blockHeight <type>', 'blockHeight')
.requiredOption('--nfflRegistryRollup <type>', 'nfflRegistryRollup contract address on destination chain')
.requiredOption('--aggregator <type>', 'Aggregator REST API')
.requiredOption('--envKey <type>', 'Name of the environment variable containing the private key')
.action(updateStateRoot);

program
.command('updateOperatorSet')
.description('Update the operator set for the nfflRegistryRollup contract')
.requiredOption('--rpcUrl <type>', 'RPC URL')
.requiredOption('--nfflRegistryRollup <type>', 'nfflRegistryRollup contract address on destination chain')
.requiredOption('--aggregator <type>', 'Aggregator REST API')
.requiredOption('--envKey <type>', 'Name of the environment variable containing the private key')
.action(updateOperatorSet);

program.parse(process.argv);
