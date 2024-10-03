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
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/NethermindEth/near-sffl/aggregator/blsagg"
	dbmocks "github.com/NethermindEth/near-sffl/aggregator/database/mocks"
	"github.com/NethermindEth/near-sffl/aggregator/database/models"
	aggmocks "github.com/NethermindEth/near-sffl/aggregator/mocks"
	"github.com/NethermindEth/near-sffl/aggregator/types"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
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
var MOCK_OPERATOR_PUBKEYS = eigentypes.OperatorPubkeys{
	G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
	G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
}
var MOCK_OPERATOR_PUBKEY_DICT = map[eigentypes.OperatorId]types.OperatorInfo{
	MOCK_OPERATOR_ID: {
		OperatorPubkeys: MOCK_OPERATOR_PUBKEYS,
		OperatorAddr:    common.Address{},
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

	aggregator, mockAvsReaderer, mockAvsWriterer, mockTaskBlsAggService, _, _, _, _, _, mockClient, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint64(100)
	var FROM_TIMESTAMP = uint64(30_000)
	var TO_TIMESTAMP = uint64(40_000)

	mockClient.EXPECT().BlockNumber(context.Background()).Return(uint64(BLOCK_NUMBER), nil)
	mockClient.EXPECT().BlockByNumber(context.Background(), big.NewInt(int64(BLOCK_NUMBER))).Return(
		gethtypes.NewBlockWithHeader(&gethtypes.Header{Time: TO_TIMESTAMP}),
		nil,
	)

	mockAvsWriterer.EXPECT().SendNewCheckpointTask(
		context.Background(),
		FROM_TIMESTAMP,
		TO_TIMESTAMP-uint64(types.MESSAGE_SUBMISSION_TIMEOUT.Seconds())-uint64(types.MESSAGE_BLS_AGGREGATION_TIMEOUT.Seconds()),
		types.TASK_QUORUM_THRESHOLD,
		coretypes.QUORUM_NUMBERS,
	).Return(aggmocks.MockSendNewCheckpointTask(uint32(BLOCK_NUMBER), TASK_INDEX, FROM_TIMESTAMP, TO_TIMESTAMP))
	mockAvsReaderer.EXPECT().GetLastCheckpointToTimestamp(context.Background()).Return(FROM_TIMESTAMP-1, nil)

	// 100 blocks, each takes 12 seconds. We hardcode for now since aggregator also hardcodes this value
	taskTimeToExpiry := (100-15)*12*time.Second - 1*time.Minute
	taskAggregationTimeout := 1 * time.Minute
	// make sure that initializeNewTask was called on the blsAggService
	// maybe there's a better way to do this? There's a saying "don't mock 3rd party code"
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockTaskBlsAggService.EXPECT().InitializeMessageIfNotExists(
		messages.CheckpointTaskResponse{ReferenceTaskIndex: TASK_INDEX}.Key(),
		coretypes.QUORUM_NUMBERS,
		[]eigentypes.QuorumThresholdPercentage{types.TASK_AGGREGATION_QUORUM_THRESHOLD},
		taskTimeToExpiry,
		taskAggregationTimeout,
		BLOCK_NUMBER,
	)

	aggregator.sendNewCheckpointTask()
}

func TestHandleStateRootUpdateAggregationReachedQuorum(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, mockMsgDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	msg := messages.StateRootUpdateMessage{}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	blsAggServiceResp := blsagg.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest: msgDigest,
		},
		Message:  msg,
		Finished: true,
	}

	model := models.NewStateRootUpdateMessageModel(msg)

	// get first return from StoreStateRootUpdate and use it as first argument on StoreStateRootUpdateAggregation
	mockMsgDb.EXPECT().StoreStateRootUpdate(msg).Return(&model, nil)
	mockMsgDb.EXPECT().StoreStateRootUpdateAggregation(&model, blsAggServiceResp.MessageBlsAggregation)

	aggregator.handleStateRootUpdateReachedQuorum(blsAggServiceResp)
}

func TestHandleOperatorSetUpdateAggregationReachedQuorum(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, mockMsgDb, mockRollupBroadcaster, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	msg := messages.OperatorSetUpdateMessage{}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	blsAggServiceResp := blsagg.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest:       msgDigest,
			NonSignersPubkeysG1: make([]*bls.G1Point, 0),
			SignersApkG2:        bls.NewZeroG2Point(),
			SignersAggSigG1:     bls.NewZeroSignature(),
		},
		Message:  msg,
		Finished: true,
	}

	msgModel := models.NewOperatorSetUpdateMessageModel(msg)

	mockMsgDb.EXPECT().StoreOperatorSetUpdate(msg).Return(&msgModel, nil)
	mockMsgDb.EXPECT().StoreOperatorSetUpdateAggregation(&msgModel, blsAggServiceResp.MessageBlsAggregation)

	signatureInfo := blsAggServiceResp.ExtractBindingRollup()
	mockRollupBroadcaster.EXPECT().BroadcastOperatorSetUpdate(context.Background(), msg, signatureInfo)

	aggregator.handleOperatorSetUpdateReachedQuorum(context.Background(), blsAggServiceResp)
}

func TestTimeoutStateRootUpdateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.NoError(t, err)

	nowTimestamp := uint64(6000)
	aggregator.clock = core.Clock{Now: func() time.Time { return time.Unix(int64(nowTimestamp), 0) }}
	messageTimestamp := nowTimestamp - 60 - 1 // 60 seconds for message submission timeout and 1 second to be out of range

	err = aggregator.ProcessSignedStateRootUpdateMessage(&messages.SignedStateRootUpdateMessage{
		Message: messages.StateRootUpdateMessage{
			Timestamp: messageTimestamp,
		},
	})

	assert.Equal(t, MessageTimeoutError, err)
}

func TestTimeoutOperatorSetUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.NoError(t, err)

	nowTimestamp := uint64(8000)
	aggregator.clock = core.Clock{Now: func() time.Time { return time.Unix(int64(nowTimestamp), 0) }}
	messageTimestamp := nowTimestamp - 60 - 1 // 60 seconds for message submission timeout and 1 second to be out of range

	err = aggregator.ProcessSignedOperatorSetUpdateMessage(&messages.SignedOperatorSetUpdateMessage{
		Message: messages.OperatorSetUpdateMessage{
			Timestamp: messageTimestamp,
		},
	})

	assert.Equal(t, MessageTimeoutError, err)
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[eigentypes.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsReaderer, *chainiomocks.MockAvsWriterer, *aggmocks.MockMessageBlsAggregationService, *aggmocks.MockMessageBlsAggregationService, *aggmocks.MockMessageBlsAggregationService, *aggmocks.MockOperatorRegistrationsService, *dbmocks.MockDatabaser, *aggmocks.MockRollupBroadcasterer, *safeclientmocks.MockSafeClient, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockAvsReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
	mockTaskBlsAggregationService := aggmocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockStateRootUpdateBlsAggregationService := aggmocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockOperatorSetUpdateBlsAggregationService := aggmocks.NewMockMessageBlsAggregationService(mockCtrl)
	mockMsgDb := dbmocks.NewMockDatabaser(mockCtrl)
	mockRollupBroadcaster := aggmocks.NewMockRollupBroadcasterer(mockCtrl)
	mockClient := safeclientmocks.NewMockSafeClient(mockCtrl)
	mockOperatorRegistrationsService := aggmocks.NewMockOperatorRegistrationsService(mockCtrl)

	aggregator := &Aggregator{
		logger:                                 logger,
		avsWriter:                              mockAvsWriter,
		avsReader:                              mockAvsReader,
		taskBlsAggregationService:              mockTaskBlsAggregationService,
		stateRootUpdateBlsAggregationService:   mockStateRootUpdateBlsAggregationService,
		operatorSetUpdateBlsAggregationService: mockOperatorSetUpdateBlsAggregationService,
		operatorRegistrationsService:           mockOperatorRegistrationsService,
		msgDb:                                  mockMsgDb,
		tasks:                                  make(map[coretypes.TaskIndex]taskmanager.CheckpointTask),
		rollupBroadcaster:                      mockRollupBroadcaster,
		httpClient:                             mockClient,
		wsClient:                               mockClient,
		aggregatorListener:                     &SelectiveAggregatorListener{},
		clock:                                  core.SystemClock,
	}
	return aggregator, mockAvsReader, mockAvsWriter, mockTaskBlsAggregationService, mockStateRootUpdateBlsAggregationService, mockOperatorSetUpdateBlsAggregationService, mockOperatorRegistrationsService, mockMsgDb, mockRollupBroadcaster, mockClient, nil
}
