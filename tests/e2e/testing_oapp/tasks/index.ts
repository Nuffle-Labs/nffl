import {task} from 'hardhat/config';

import {
  createGetHreByEid,
  createProviderFactory,
  getEidForNetworkName,
} from '@layerzerolabs/devtools-evm-hardhat';
import {Options} from '@layerzerolabs/lz-v2-utilities';

// send messages from a contract on one network to another
task('send', 'test send')
  // contract to send a message from
  .addParam('contractA', 'contract address on network A')
  // network that sender contract resides on
  .addParam('networkA', 'name of the network A')
  // network that receiver contract resides on
  .addParam('networkB', 'name of the network B')
  // message to send from network a to network b
  //.addParam('message', 'message to send from network A to network B')
  .setAction(async (taskArgs, {ethers}) => {
    const eidA = getEidForNetworkName(taskArgs.networkA);
    const eidB = getEidForNetworkName(taskArgs.networkB);
    const contractA = taskArgs.contractA;
    const environmentFactory = createGetHreByEid();
    const providerFactory = createProviderFactory(environmentFactory);
    const signer = (await providerFactory(eidA)).getSigner();

    const oappContractFactory = await ethers.getContractFactory('TestingOApp', signer);
    const oapp = oappContractFactory.attach(contractA);

    const options = Options.newOptions().addExecutorLzReceiveOption(200000, 0).toHex().toString();
    //const msg = taskArgs.message == "" ? "Hola Mundo" : taskArgs.message;
    const msg = "Hola Nuff";
    const [nativeFee] = await oapp.quote(eidB, msg, options, false);
    console.log('native fee:', nativeFee);

    const r = await oapp.send(eidB, msg, options, {
      value: nativeFee,
    });

    console.log(`Tx initiated. See: https://layerzeroscan.com/tx/${r.hash}`);
  });


task('read', 'read message stored in OApp')
  .addParam('contractA', 'contract address on network A')
  .addParam('contractB', 'contract address on network B')
  .addParam('networkA', 'name of the network A')
  .addParam('networkB', 'name of the network B')
  .setAction(async (taskArgs, {ethers}) => {
    const eidA = getEidForNetworkName(taskArgs.networkA);
    const eidB = getEidForNetworkName(taskArgs.networkB);
    const contractA = taskArgs.contractA;
    const contractB = taskArgs.contractB;
    const environmentFactory = createGetHreByEid();
    const providerFactory = createProviderFactory(environmentFactory);
    const signerA = (await providerFactory(eidA)).getSigner();
    const signerB = (await providerFactory(eidB)).getSigner();

    const oappContractAFactory = await ethers.getContractFactory('TestingOApp', signerA);
    const oappContractBFactory = await ethers.getContractFactory('TestingOApp', signerB);

    const oappA = oappContractAFactory.attach(contractA);
    const oappB = oappContractBFactory.attach(contractB);

    const dataOnOAppA = await oappA.data();
    const dataOnOAppB = await oappB.data();

    console.log({
      [taskArgs.networkA]: dataOnOAppA,
      [taskArgs.networkB]: dataOnOAppB,
    });
  });


import holesky from "../deployments/holesky/TestingOApp.json";
import amoy from "../deployments/amoy-testnet/TestingOApp.json";
import arbitrum from "../deployments/arbitrum-sepolia/TestingOApp.json";
import { assert } from 'console';

const chainSelector = ((chain: string) => {
      switch (chain) {
        case "holesky": return holesky['address'];
        case "amoy-testnet": return amoy['address'];
        case "arbitrum-sepolia": return arbitrum['address'];
    }});

task("e2e", "Send a message and check that it has arrived.")
  .addParam("source", "source chain of the message")
  .addParam("target", "target chain of the message")
  .setAction(async (taskArgs, {ethers}) => {
    console.log("=== Regular DVN used ===");

    // Create the necessary data for querying the contracts and calling them
    const sourceEid = getEidForNetworkName(taskArgs.source);
    const targetEid = getEidForNetworkName(taskArgs.target);

    const sourceContractAddr = chainSelector(taskArgs.source);
    const targetContractAddr = chainSelector(taskArgs.target);

    const sourceContract = sourceContractAddr ? sourceContractAddr : "0x0";
    const targetContract = targetContractAddr ? targetContractAddr : "0x0";

    const environmentFactory = createGetHreByEid();
    const providerFactory = createProviderFactory(environmentFactory);

    const sourceSigner = (await providerFactory(sourceEid)).getSigner();
    const targetSigner = (await providerFactory(targetEid)).getSigner();

    const oappSourceContractFactory = await ethers.getContractFactory('TestingOApp', sourceSigner);
    const oappTargetContractFactory = await ethers.getContractFactory('TestingOApp', targetSigner);

    const sourceOapp = oappSourceContractFactory.attach(sourceContract);
    const targetOapp = oappTargetContractFactory.attach(targetContract);

    // Query the data on the target chain using the OApp contract.
    const initialdataOnOAppB = await targetOapp.data();
    const initialData = Number(initialdataOnOAppB);
    console.log("Initial data on target OApp:", Number(initialData));

    // Increase the queried value by one, and send it.
    const msg = initialData + 1;

    const options = Options.newOptions().addExecutorLzReceiveOption(200000, 0).toHex().toString();
    const [nativeFee] = await sourceOapp.quote(targetEid, msg, options, false);
    console.log('native fee:', nativeFee);

    const r = await sourceOapp.send(targetEid, initialData, options, {
      value: nativeFee,
    });

    console.log(`Tx initiated. See: https://layerzeroscan.com/tx/${r.hash}`);
    console.log("Waiting 30 seconds for the message to arrive...");

    // Wait for the transaction to arrive on the target chain.
    await new Promise((resolve) => setTimeout(resolve, 30000));

    // Check the value has changed.
    const dataOnOAppB = await sourceOapp.data();
    assert(msg == dataOnOAppB, "Newly queried data differs from the expected on. Message did not arrive on target OApp");
  })

import simple_holesky from "../deployments/holesky/SimpleTestingOApp.json";
import simple_amoy from "../deployments/amoy-testnet/SimpleTestingOApp.json";
import simple_arbitrum from "../deployments/arbitrum-sepolia/SimpleTestingOApp.json";

const chainSelectorSimple = ((chain: string) => {
      switch (chain) {
        case "holesky": return simple_holesky['address'];
        case "amoy-testnet": return simple_amoy['address'];
        case "arbitrum-sepolia": return simple_arbitrum['address'];
    }});

task("e2e:simple", "Send a message and check that it has arrived.")
  .addParam("source", "source chain of the message")
  .addParam("target", "target chain of the message")
  .setAction(async (taskArgs, {ethers}) => {
    console.log("=== Simple DVN used ===");

    // Create the necessary data for querying the contracts and calling them
    const sourceEid = getEidForNetworkName(taskArgs.source);
    const targetEid = getEidForNetworkName(taskArgs.target);

    const sourceContractAddr = chainSelectorSimple(taskArgs.source);
    const targetContractAddr = chainSelectorSimple(taskArgs.target);

    const sourceContract = sourceContractAddr ? sourceContractAddr : "0x0";
    const targetContract = targetContractAddr ? targetContractAddr : "0x0";

    const environmentFactory = createGetHreByEid();
    const providerFactory = createProviderFactory(environmentFactory);

    const sourceSigner = (await providerFactory(sourceEid)).getSigner();
    const targetSigner = (await providerFactory(targetEid)).getSigner();

    const oappSourceContractFactory = await ethers.getContractFactory('TestingOApp', sourceSigner);
    const oappTargetContractFactory = await ethers.getContractFactory('TestingOApp', targetSigner);

    const sourceOapp = oappSourceContractFactory.attach(sourceContract);
    const targetOapp = oappTargetContractFactory.attach(targetContract);

    // Query the data on the target chain using the OApp contract.
    const initialdataOnOAppB = await targetOapp.data();
    const initialData = Number(initialdataOnOAppB);
    console.log("Initial data on target OApp:", Number(initialData));

    // Increase the queried value by one, and send it.
    const msg = initialData + 1;

    const options = Options.newOptions().addExecutorLzReceiveOption(200000, 0).toHex().toString();
    const [nativeFee] = await sourceOapp.quote(targetEid, msg, options, false);
    console.log('native fee:', nativeFee);

    const r = await sourceOapp.send(targetEid, initialData, options, {
      value: nativeFee,
    });

    console.log(`Tx initiated. See: https://layerzeroscan.com/tx/${r.hash}`);
    console.log("Waiting 30 seconds for the message to arrive...");

    // Wait for the transaction to arrive on the target chain.
    await new Promise((resolve) => setTimeout(resolve, 30000));

    // Check the value has changed.
    const dataOnOAppB = await sourceOapp.data();
    assert(msg == dataOnOAppB, "Newly queried data differs from the expected on. Message did not arrive on target OApp");
  })

