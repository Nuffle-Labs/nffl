package aggregator

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

func TestGetStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	msg := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaCommitment:    keccak256(4),
		NearDaTransactionId: keccak256(5),
		StateRoot:           keccak256(6),
	}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	aggregation := aggtypes.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest: msgDigest,
		},
	}

	mockDb.EXPECT().FetchStateRootUpdate(msg.RollupId, msg.BlockHeight).Return(&msg, nil)

	mockDb.EXPECT().FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight).Return(&aggregation.MessageBlsAggregation, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%d", msg.RollupId, msg.BlockHeight),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	aggregator.handleGetStateRootUpdateAggregation(recorder, req)

	expectedBody := aggtypes.GetStateRootUpdateAggregationResponse{
		Message:     msg,
		Aggregation: aggregation.MessageBlsAggregation,
	}
	var body aggtypes.GetStateRootUpdateAggregationResponse

	assert.Equal(t, http.StatusOK, recorder.Code)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, expectedBody, body)
}

func TestGetOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	msg := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	aggregation := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	mockDb.EXPECT().FetchOperatorSetUpdate(msg.Id).Return(&msg, nil)

	mockDb.EXPECT().FetchOperatorSetUpdateAggregation(msg.Id).Return(&aggregation, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/operator-set-update?id=%d", msg.Id),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	aggregator.handleGetOperatorSetUpdateAggregation(recorder, req)

	expectedBody := aggtypes.GetOperatorSetUpdateAggregationResponse{
		Message:     msg,
		Aggregation: aggregation,
	}
	var body aggtypes.GetOperatorSetUpdateAggregationResponse

	assert.Equal(t, http.StatusOK, recorder.Code)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, expectedBody, body)
}

func TestGetCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	msg := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
	}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	aggregation := messages.MessageBlsAggregation{
		MessageDigest: msgDigest,
	}

	msg2 := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	msgDigest2, err := msg2.Digest()
	assert.Nil(t, err)

	aggregation2 := messages.MessageBlsAggregation{
		MessageDigest: msgDigest2,
	}

	mockDb.EXPECT().FetchCheckpointMessages(uint64(0), uint64(3)).Return(&messages.CheckpointMessages{
		StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg},
		StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation},
		OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg2},
		OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation2},
	}, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/checkpoint/messages?fromTimestamp=%d&toTimestamp=%d", 0, 3),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	aggregator.handleGetCheckpointMessages(recorder, req)

	expectedBody := aggtypes.GetCheckpointMessagesResponse{
		CheckpointMessages: messages.CheckpointMessages{
			StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg},
			StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation},
			OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg2},
			OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation2},
		},
	}
	var body aggtypes.GetCheckpointMessagesResponse

	assert.Equal(t, http.StatusOK, recorder.Code)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, expectedBody, body)
}

func TestGetStateRootUpdateAggregation_MissingParameters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	t.Run("Missing rollupId", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/aggregation/state-root-update?blockHeight=%d", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Missing blockHeight", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/aggregation/state-root-update?&rollupId=%d", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetStateRootUpdateAggregation_InvalidParameters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	t.Run("Invalid rollupId - incorrect type", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/aggregation/state-root-update?rollupId=%s&blockHeight=%d", "foo", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Invalid rollupId - too large", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%d", uint64(0xFFFFFFFFFFFFFFFF), 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Invalid blockHeight - incorrect type", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%s", 0, "foo"),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetStateRootUpdateAggregation_EmptyParameters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	req, err := http.NewRequest(
		"GET",
		"/aggregation/state-root-update?rollupId=&blockHeight=",
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestGetStateRootUpdateAggregation_StateRootUpdateNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	notFound := errors.New("not found")
	mockDb.EXPECT().FetchStateRootUpdate(gomock.Any(), gomock.Any()).Return(nil, notFound)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%d", 0, 0),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetStateRootUpdateAggregation_StateRootUpdateAggregationNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	notFound := errors.New("not found")
	mockDb.EXPECT().FetchStateRootUpdate(gomock.Any(), gomock.Any()).Return(&messages.StateRootUpdateMessage{}, nil)
	mockDb.EXPECT().FetchStateRootUpdateAggregation(gomock.Any(), gomock.Any()).Return(nil, notFound)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%d", 0, 0),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetStateRootUpdateAggregation(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetOperatorSetUpdateAggregation_MissingParameter(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	t.Run("Missing id", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			"/aggregation/operator-set-update",
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetOperatorSetUpdateAggregation(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetOperatorSetUpdateAggregation_OperatorSetUpdateNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	notFound := errors.New("not found")
	mockDb.EXPECT().FetchOperatorSetUpdate(gomock.Any()).Return(nil, notFound)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/operator-set-update?id=%d", 0),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetOperatorSetUpdateAggregation(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetOperatorSetUpdateAggregation_OperatorSetUpdateAggregationNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	notFound := errors.New("not found")
	mockDb.EXPECT().FetchOperatorSetUpdate(gomock.Any()).Return(&messages.OperatorSetUpdateMessage{}, nil)
	mockDb.EXPECT().FetchOperatorSetUpdateAggregation(gomock.Any()).Return(nil, notFound)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/operator-set-update?id=%d", 0),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetOperatorSetUpdateAggregation(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetCheckpointMessages_MissingParameters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	t.Run("Missing fromTimestamp", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/checkpoint/messages?toTimestamp=%d", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetCheckpointMessages(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Missing toTimestamp", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/checkpoint/messages?fromTimestamp=%d", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetCheckpointMessages(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetCheckpointMessages_InvalidParameters(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, _, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	t.Run("Invalid fromTimestamp - incorrect type", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/checkpoint/messages?fromTimestamp=%s&toTimestamp=%d", "foo", 0),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetCheckpointMessages(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Invalid toTimestamp - incorrect type", func(t *testing.T) {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("/checkpoint/messages?fromTimestamp=%d&toTimestamp=%s", 0, "foo"),
			nil,
		)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		err = aggregator.handleGetCheckpointMessages(recorder, req)
		assert.NotNil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetCheckpointMessages_CheckpointMessageNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	notFound := errors.New("not found")
	mockDb.EXPECT().FetchCheckpointMessages(gomock.Any(), gomock.Any()).Return(nil, notFound)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/checkpoint/messages?fromTimestamp=%d&toTimestamp=%d", 100, 200),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()

	err = aggregator.handleGetCheckpointMessages(recorder, req)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
