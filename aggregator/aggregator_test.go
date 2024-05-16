package aggregator

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	blsaggservmock "github.com/Layr-Labs/eigensdk-go/services/mocks/blsagg"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	dbmocks "github.com/NethermindEth/near-sffl/aggregator/database/mocks"
	aggmocks "github.com/NethermindEth/near-sffl/aggregator/mocks"
	"github.com/NethermindEth/near-sffl/aggregator/types"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	chainiomocks "github.com/NethermindEth/near-sffl/core/chainio/mocks"
	safeclientmocks "github.com/NethermindEth/near-sffl/core/safeclient/mocks"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"
var MOCK_OPERATOR_BLS_PRIVATE_KEY, _ = bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
var MOCK_OPERATOR_KEYPAIR = bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
var MOCK_OPERATOR_G1PUBKEY = MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
var MOCK_OPERATOR_G2PUBKEY = MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()
var MOCK_OPERATOR_PUBKEY_DICT = map[eigentypes.OperatorId]types.OperatorInfo{
	MOCK_OPERATOR_ID: {
		OperatorPubkeys: eigentypes.OperatorPubkeys{
			G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
			G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
		},
		OperatorAddr: common.Address{},
	},
}

type MockTask struct {
	TaskNum       uint32
	BlockNumber   uint32
	FromTimestamp uint64
	ToTimestamp   uint64
}

func TestSendNewTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, mockAvsReaderer, mockAvsWriterer, mockTaskBlsAggService, _, _, _, _, mockClient, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var FROM_TIMESTAMP = uint64(3)
	var TO_TIMESTAMP = uint64(4)

	mockClient.EXPECT().BlockNumber(context.Background()).Return(uint64(BLOCK_NUMBER), nil)
	mockClient.EXPECT().BlockByNumber(context.Background(), big.NewInt(int64(BLOCK_NUMBER))).Return(
		gethtypes.NewBlockWithHeader(&gethtypes.Header{Time: TO_TIMESTAMP}),
		nil,
	)

	mockAvsWriterer.EXPECT().SendNewCheckpointTask(
		context.Background(), FROM_TIMESTAMP, TO_TIMESTAMP, types.TASK_QUORUM_THRESHOLD, coretypes.QUORUM_NUMBERS,
	).Return(aggmocks.MockSendNewCheckpointTask(BLOCK_NUMBER, TASK_INDEX, FROM_TIMESTAMP, TO_TIMESTAMP))
	mockAvsReaderer.EXPECT().GetLastCheckpointToTimestamp(context.Background()).Return(FROM_TIMESTAMP-1, nil)

	// 100 blocks, each takes 12 seconds. We hardcode for now since aggregator also hardcodes this value
	taskTimeToExpiry := 100 * 12 * time.Second
	// make sure that initializeNewTask was called on the blsAggService
	// maybe there's a better way to do this? There's a saying "don't mock 3rd party code"
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockTaskBlsAggService.EXPECT().InitializeNewTask(TASK_INDEX, BLOCK_NUMBER, coretypes.QUORUM_NUMBERS, []eigentypes.QuorumThresholdPercentage{types.TASK_AGGREGATION_QUORUM_THRESHOLD}, taskTimeToExpiry)

	aggregator.sendNewCheckpointTask()
}

func TestHandleStateRootUpdateAggregationReachedQuorum(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockMsgDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	blsAggServiceResp := types.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest: msgDigest,
		},
		Finished: true,
	}

	aggregator.stateRootUpdates[msgDigest] = msg

	mockMsgDb.EXPECT().StoreStateRootUpdate(msg)
	mockMsgDb.EXPECT().StoreStateRootUpdateAggregation(msg, blsAggServiceResp.MessageBlsAggregation)

	assert.Contains(t, aggregator.stateRootUpdates, msgDigest)

	aggregator.handleStateRootUpdateReachedQuorum(blsAggServiceResp)

	assert.NotContains(t, aggregator.stateRootUpdates, msgDigest)
}

func TestHandleOperatorSetUpdateAggregationReachedQuorum(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockMsgDb, mockRollupBroadcaster, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	msg := messages.OperatorSetUpdateMessage{}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	blsAggServiceResp := types.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest:       msgDigest,
			NonSignersPubkeysG1: make([]*bls.G1Point, 0),
			SignersApkG2:        bls.NewZeroG2Point(),
			SignersAggSigG1:     bls.NewZeroSignature(),
		},
		Finished: true,
	}

	aggregator.operatorSetUpdates[msgDigest] = msg

	mockMsgDb.EXPECT().StoreOperatorSetUpdate(msg)
	mockMsgDb.EXPECT().StoreOperatorSetUpdateAggregation(msg, blsAggServiceResp.MessageBlsAggregation)

	signatureInfo := blsAggServiceResp.ExtractBindingRollup()
	mockRollupBroadcaster.EXPECT().BroadcastOperatorSetUpdate(context.Background(), msg, signatureInfo)

	assert.Contains(t, aggregator.operatorSetUpdates, msgDigest)

	aggregator.handleOperatorSetUpdateReachedQuorum(context.Background(), blsAggServiceResp)

	assert.NotContains(t, aggregator.operatorSetUpdates, msgDigest)
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[eigentypes.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsReaderer, *chainiomocks.MockAvsWriterer, *blsaggservmock.MockBlsAggregationService, *aggmocks.MockMessageBlsAggregationService, *aggmocks.MockMessageBlsAggregationService, *dbmocks.MockDatabaser, *aggmocks.MockRollupBroadcasterer, *safeclientmocks.MockSafeClient, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockAvsReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
	mockTaskBlsAggregationService := blsaggservmock.NewMockBlsAggregationService(mockCtrl)
	mockStateRootUpdateBlsAggregationService := aggmocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockOperatorSetUpdateBlsAggregationService := aggmocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockMsgDb := dbmocks.NewMockDatabaser(mockCtrl)
	mockRollupBroadcaster := aggmocks.NewMockRollupBroadcasterer(mockCtrl)
	mockClient := safeclientmocks.NewMockSafeClient(mockCtrl)

	aggregator := &Aggregator{
		logger:                                 logger,
		avsWriter:                              mockAvsWriter,
		avsReader:                              mockAvsReader,
		taskBlsAggregationService:              mockTaskBlsAggregationService,
		stateRootUpdateBlsAggregationService:   mockStateRootUpdateBlsAggregationService,
		operatorSetUpdateBlsAggregationService: mockOperatorSetUpdateBlsAggregationService,
		msgDb:                                  mockMsgDb,
		tasks:                                  make(map[coretypes.TaskIndex]taskmanager.CheckpointTask),
		taskResponses:                          make(map[coretypes.TaskIndex]map[eigentypes.TaskResponseDigest]messages.CheckpointTaskResponse),
		stateRootUpdates:                       make(map[coretypes.MessageDigest]messages.StateRootUpdateMessage),
		operatorSetUpdates:                     make(map[coretypes.MessageDigest]messages.OperatorSetUpdateMessage),
		rollupBroadcaster:                      mockRollupBroadcaster,
		httpClient:                             mockClient,
		wsClient:                               mockClient,
		rpcListener:                            &SelectiveRpcListener{},
		restListener:                           &SelectiveRestListener{},
		aggregatorListener:                     &SelectiveAggregatorListener{},
	}
	return aggregator, mockAvsReader, mockAvsWriter, mockTaskBlsAggregationService, mockStateRootUpdateBlsAggregationService, mockOperatorSetUpdateBlsAggregationService, mockMsgDb, mockRollupBroadcaster, mockClient, nil
}
