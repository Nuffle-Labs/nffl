package aggregator

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	blsaggservmock "github.com/Layr-Labs/eigensdk-go/services/mocks/blsagg"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/NethermindEth/near-sffl/aggregator/mocks"
	"github.com/NethermindEth/near-sffl/aggregator/types"
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"

type MockTask struct {
	TaskNum       uint32
	BlockNumber   uint32
	FromTimestamp uint64
	ToTimestamp   uint64
}

func TestSendNewTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[bls.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}

	aggregator, mockAvsWriterer, mockTaskBlsAggService, _, _, _, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var FROM_NEAR_BLOCK = uint64(3)
	var TO_NEAR_BLOCK = uint64(4)

	mockAvsWriterer.EXPECT().SendNewCheckpointTask(
		context.Background(), FROM_NEAR_BLOCK, TO_NEAR_BLOCK, types.QUORUM_THRESHOLD_NUMERATOR, coretypes.QUORUM_NUMBERS,
	).Return(mocks.MockSendNewCheckpointTask(BLOCK_NUMBER, TASK_INDEX, FROM_NEAR_BLOCK, TO_NEAR_BLOCK))

	// 100 blocks, each takes 12 seconds. We hardcode for now since aggregator also hardcodes this value
	taskTimeToExpiry := 100 * 12 * time.Second
	// make sure that initializeNewTask was called on the blsAggService
	// maybe there's a better way to do this? There's a saying "don't mock 3rd party code"
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockTaskBlsAggService.EXPECT().InitializeNewTask(TASK_INDEX, BLOCK_NUMBER, coretypes.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, taskTimeToExpiry)

	err = aggregator.sendNewCheckpointTask(FROM_NEAR_BLOCK, TO_NEAR_BLOCK)
	assert.Nil(t, err)
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[bls.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsWriterer, *blsaggservmock.MockBlsAggregationService, *mocks.MockMessageBlsAggregationService, *mocks.MockMessageBlsAggregationService, *mocks.MockMessageDatabaser, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockTaskBlsAggregationService := blsaggservmock.NewMockBlsAggregationService(mockCtrl)
	mockStateRootUpdateBlsAggregationService := mocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockOperatorSetUpdateBlsAggregationService := mocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockMsgDb := mocks.NewMockMessageDatabaser(mockCtrl)

	aggregator := &Aggregator{
		logger:                                 logger,
		avsWriter:                              mockAvsWriter,
		taskBlsAggregationService:              mockTaskBlsAggregationService,
		stateRootUpdateBlsAggregationService:   mockStateRootUpdateBlsAggregationService,
		operatorSetUpdateBlsAggregationService: mockOperatorSetUpdateBlsAggregationService,
		msgDb:                                  mockMsgDb,
		tasks:                                  make(map[coretypes.TaskIndex]taskmanager.CheckpointTask),
		taskResponses:                          make(map[coretypes.TaskIndex]map[sdktypes.TaskResponseDigest]taskmanager.CheckpointTaskResponse),
		stateRootUpdates:                       make(map[coretypes.MessageDigest]servicemanager.StateRootUpdateMessage),
		operatorSetUpdates:                     make(map[coretypes.MessageDigest]registryrollup.OperatorSetUpdateMessage),
	}
	return aggregator, mockAvsWriter, mockTaskBlsAggregationService, mockStateRootUpdateBlsAggregationService, mockOperatorSetUpdateBlsAggregationService, mockMsgDb, nil
}

// just a mock ethclient to pass to bindings
// so that we can access abi methods
func createMockEthClient() *backends.SimulatedBackend {
	genesisAlloc := map[common.Address]gethcore.GenesisAccount{}
	blockGasLimit := uint64(1000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client
}
