package chainio

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	erc20mock "github.com/NethermindEth/near-sffl/contracts/bindings/ERC20Mock"
	opsetupdatereg "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/config"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

type AvsReaderer interface {
	sdkavsregistry.AvsRegistryReader

	CheckSignatures(
		ctx context.Context, quorumNumbers []byte, aggregation messages.MessageBlsAggregation,
	) (taskmanager.IBLSSignatureCheckerQuorumStakeTotals, error)
	GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error)
	GetOperatorSetUpdateDelta(ctx context.Context, id uint64) ([]opsetupdatereg.RollupOperatorsOperator, error)
	GetOperatorSetById(ctx context.Context, id uint64) ([]opsetupdatereg.RollupOperatorsOperator, error)
	GetOperatorSetUpdateBlock(ctx context.Context, id uint64) (uint32, error)
	GetNextOperatorSetUpdateId(ctx context.Context) (uint64, error)
	GetLastCheckpointToTimestamp(ctx context.Context) (uint64, error)
	GetMessagingPrefix(ctx context.Context) ([32]byte, error)
}

type AvsReader struct {
	sdkavsregistry.AvsRegistryReader
	AvsServiceBindings *AvsManagersBindings
	logger             logging.Logger
}

var _ AvsReaderer = (*AvsReader)(nil)

func BuildAvsReaderFromConfig(config *config.Config, client eth.Client, logger logging.Logger) (*AvsReader, error) {
	return BuildAvsReader(config.SFFLRegistryCoordinatorAddr, config.OperatorStateRetrieverAddr, client, logger)
}

func BuildAvsReader(registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsReader, error) {
	avsManagersBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	avsRegistryReader, err := sdkavsregistry.BuildAvsRegistryChainReader(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	return NewAvsReader(avsRegistryReader, avsManagersBindings, logger)
}

func NewAvsReader(avsRegistryReader sdkavsregistry.AvsRegistryReader, avsServiceBindings *AvsManagersBindings, logger logging.Logger) (*AvsReader, error) {
	return &AvsReader{
		AvsRegistryReader:  avsRegistryReader,
		AvsServiceBindings: avsServiceBindings,
		logger:             logger,
	}, nil
}

func (r *AvsReader) CheckSignatures(
	ctx context.Context, quorumNumbers []byte, aggregation messages.MessageBlsAggregation,
) (taskmanager.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.AvsServiceBindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, aggregation.MessageDigest, quorumNumbers, uint32(aggregation.EthBlockNumber), aggregation.ExtractBindingMainnet(),
	)
	if err != nil {
		return taskmanager.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}

func (r *AvsReader) GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error) {
	erc20Mock, err := r.AvsServiceBindings.GetErc20Mock(tokenAddr)
	if err != nil {
		r.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return erc20Mock, nil
}

func (r *AvsReader) GetOperatorSetUpdateDelta(ctx context.Context, id uint64) ([]opsetupdatereg.RollupOperatorsOperator, error) {
	result, err := r.AvsServiceBindings.OperatorSetUpdateRegistry.GetOperatorSetUpdate(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}

	type operatorUpdate struct {
		pubkey         opsetupdatereg.BN254G1Point
		previousWeight *big.Int
		newWeight      *big.Int
	}

	operators := make(map[string]operatorUpdate)

	for _, operator := range result.PreviousOperatorSet {
		operatorKey := fmt.Sprintf("%s_%s", operator.Pubkey.X.String(), operator.Pubkey.Y.String())
		operators[operatorKey] = operatorUpdate{operator.Pubkey, operator.Weight, big.NewInt(0)}
	}

	for _, operator := range result.NewOperatorSet {
		operatorKey := fmt.Sprintf("%s_%s", operator.Pubkey.X.String(), operator.Pubkey.Y.String())
		weights, ok := operators[operatorKey]

		if ok {
			weights.newWeight = operator.Weight
		} else {
			weights = operatorUpdate{operator.Pubkey, big.NewInt(0), operator.Weight}
		}

		operators[operatorKey] = weights
	}

	var delta []opsetupdatereg.RollupOperatorsOperator

	for _, operatorUpdate := range operators {
		if operatorUpdate.previousWeight.Cmp(operatorUpdate.newWeight) != 0 {
			delta = append(delta, opsetupdatereg.RollupOperatorsOperator{Pubkey: operatorUpdate.pubkey, Weight: operatorUpdate.newWeight})
		}
	}

	nonSignersPubkeyHashes := make([][32]byte, 0, len(delta))
	for _, operatorDelta := range delta {
		hash, err := core.HashBNG1Point(taskmanager.BN254G1Point(operatorDelta.Pubkey))
		if err != nil {
			return nil, err
		}

		nonSignersPubkeyHashes = append(nonSignersPubkeyHashes, hash)
	}

	sort.Slice(delta, func(i, j int) bool {
		return bytes.Compare(nonSignersPubkeyHashes[i][:], nonSignersPubkeyHashes[j][:]) == -1
	})

	return delta, nil
}

func (r *AvsReader) GetOperatorSetById(ctx context.Context, id uint64) ([]opsetupdatereg.RollupOperatorsOperator, error) {
	result, err := r.AvsServiceBindings.OperatorSetUpdateRegistry.GetOperatorSetUpdate(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}

	return result.NewOperatorSet, nil
}

func (r *AvsReader) GetOperatorSetUpdateBlock(ctx context.Context, id uint64) (uint32, error) {
	return r.AvsServiceBindings.OperatorSetUpdateRegistry.OperatorSetUpdateIdToBlockNumber(&bind.CallOpts{}, big.NewInt(0).SetUint64(id))
}

func (r *AvsReader) GetNextOperatorSetUpdateId(ctx context.Context) (uint64, error) {
	count, err := r.AvsServiceBindings.OperatorSetUpdateRegistry.GetOperatorSetUpdateCount(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *AvsReader) GetLastCheckpointToTimestamp(ctx context.Context) (uint64, error) {
	lastCheckpointToTimestamp, err := r.AvsServiceBindings.TaskManager.LastCheckpointToTimestamp(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return lastCheckpointToTimestamp, nil
}

func (r *AvsReader) GetMessagingPrefix(ctx context.Context) ([32]byte, error) {
	messagingPrefix, err := r.AvsServiceBindings.TaskManager.MessagingPrefix(&bind.CallOpts{})
	if err != nil {
		return [32]byte{}, err
	}
	return messagingPrefix, nil
}
