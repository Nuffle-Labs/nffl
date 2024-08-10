#!/usr/bin/env ts-node

import { program, Option } from "@commander-js/extra-typings";

import { getStorageProof } from "./commands/getStorageProof";
import { updateStateRoot } from "./commands/updateStateRoot";
import { updateOperatorSet } from "./commands/updateOperatorSet";

program
    .version("0.0.1")
    .description("NFFL demo CLI tool");

program
    .command("storage-proof")
    .description("Get the storage proof of a target contract from the source chain") 
    .requiredOption("--rpc-url <url>", "RPC URL")
    .requiredOption("--contract-address <address>", "Address of the target contract")
    .requiredOption("--storage-key <hex>", "Storage key")
    .action(getStorageProof);

program
    .command("update-state-root")
    .description("Update the state root for a given rollup in the NFFLRegistryRollup contract")
    .requiredOption("--rpc-url <url>", "RPC URL")
    .requiredOption("--contract-address <address>", "Address of the target contract")
    .requiredOption("--rollup-id <number>", "Rollup Id")
    .requiredOption("--block-height <number>", "Block height")
    .requiredOption("--aggregator-url <url>", "Aggregator REST API URL")
    .addOption(new Option("--private-key <hex>", "Private key").env("PRIVATE_KEY").makeOptionMandatory())
    .action(updateStateRoot);

program
    .command("update-operator-set")
    .description("Update the operator set for the NFFLRegistryRollup contract")
    .requiredOption("--rpc-url <url>", "RPC URL")
    .requiredOption("--contract-address <address>", "Address of the target contract")
    .requiredOption("--id <number>", "Update Id")
    .requiredOption("--aggregator-url <url>", "Aggregator REST API URL")
    .addOption(new Option("--private-key <hex>", "Private key").env("PRIVATE_KEY").makeOptionMandatory())
    .action(updateOperatorSet);

program.parse();
