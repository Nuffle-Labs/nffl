import { ethers } from "ethers";
import { OperatorSetUpdate, RollupOperators, StateRootUpdate } from "../../typechain-types/SFFLRegistryRollup";

export interface G1Point {
    X: string;
    Y: string;
}

export interface G2Point {
    X: { A0: string; A1: string };
    Y: { A0: string; A1: string };
}

export interface OperatorInfo {
    Pubkey: G1Point;
    Weight: string;
}

export interface MessageAggregation {
    NonSignersPubkeysG1: G1Point[];
    SignersApkG2: G2Point;
    SignersAggSigG1: {
        g1_point: G1Point;
    };
}

export interface GetOperatorSetUpdateResponse {
    Message: {
        Id: string;
        Timestamp: string;
        Operators: OperatorInfo[];
    };
    Aggregation: MessageAggregation;
}

export interface GetStateRootUpdateResponse {
    Message: {
        RollupId: string;
        BlockHeight: string;
        Timestamp: string;
        NearDaTransactionId: Array<number>;
        NearDaCommitment: Array<number>;
        StateRoot: Array<number>;
    };
    Aggregation: MessageAggregation;
}

function hashG1Point(pk: G1Point): bigint {
    return BigInt(ethers.keccak256(ethers.concat([
      ethers.zeroPadValue(ethers.toBeHex(BigInt(pk.X)), 32),
      ethers.zeroPadValue(ethers.toBeHex(BigInt(pk.Y)), 32)
    ])));
}

function createSignatureInfo(aggregation: MessageAggregation): RollupOperators.SignatureInfoStruct {
    return {
        nonSignerPubkeys: [...aggregation.NonSignersPubkeysG1].sort(
            (a, b) => (hashG1Point(a) < hashG1Point(b) ? -1 : hashG1Point(a) > hashG1Point(b) ? 1 : 0)
        ),
        apkG2: {
            X: [aggregation.SignersApkG2.X.A1, aggregation.SignersApkG2.X.A0],
            Y: [aggregation.SignersApkG2.Y.A1, aggregation.SignersApkG2.Y.A0]
        },
        sigma: aggregation.SignersAggSigG1.g1_point,
    };
}

function createOperatorSetUpdateMessage(data: GetOperatorSetUpdateResponse): OperatorSetUpdate.MessageStruct {
    return {
        id: data.Message.Id,
        timestamp: data.Message.Timestamp,
        operators: data.Message.Operators.map((operator) => ({
            pubkey: operator.Pubkey,
            weight: operator.Weight
        })),
    };
}

async function fetchOperatorSetUpdateData(aggregatorUrl: string, id: bigint): Promise<GetOperatorSetUpdateResponse> {
    const response = await fetch(`${aggregatorUrl}/aggregation/operator-set-update?id=${id}`);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    const text = await response.text();

    // Make sure the weight is represented as a string, since it's a uint128
    const replacedText = text.replace(/"Weight":\s*(\d+)/g, "\"Weight\": \"$1\"");

    return JSON.parse(replacedText);
}

export async function fetchOperatorSetUpdate(aggregatorUrl: string, id: bigint): Promise<{message: OperatorSetUpdate.MessageStruct, aggregation: RollupOperators.SignatureInfoStruct}> {
    const data = await fetchOperatorSetUpdateData(aggregatorUrl, id);
    
    return {
        message: createOperatorSetUpdateMessage(data),
        aggregation: createSignatureInfo(data.Aggregation),
    };
}

function createStateRootUpdateMessage(data: GetStateRootUpdateResponse): StateRootUpdate.MessageStruct {
    return {
        rollupId: data.Message.RollupId,
        blockHeight: data.Message.BlockHeight,
        timestamp: data.Message.Timestamp,
        nearDaTransactionId: Uint8Array.from(data.Message.NearDaTransactionId),
        nearDaCommitment: Uint8Array.from(data.Message.NearDaCommitment),
        stateRoot: Uint8Array.from(data.Message.StateRoot)
    };
}

async function fetchStateRootUpdateData(aggregatorUrl: string, rollupId: bigint, blockHeight: bigint): Promise<GetStateRootUpdateResponse> {
    const response = await fetch(`${aggregatorUrl}/aggregation/state-root-update?rollupId=${rollupId}&blockHeight=${blockHeight}`);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return await response.json() as GetStateRootUpdateResponse;
}

export async function fetchStateRootUpdate(aggregatorUrl: string, rollupId: bigint, blockHeight: bigint): Promise<{message: StateRootUpdate.MessageStruct, aggregation: RollupOperators.SignatureInfoStruct}> {
    const data = await fetchStateRootUpdateData(aggregatorUrl, rollupId, blockHeight);
    return {
        message: createStateRootUpdateMessage(data),
        aggregation: createSignatureInfo(data.Aggregation),
    };
}
