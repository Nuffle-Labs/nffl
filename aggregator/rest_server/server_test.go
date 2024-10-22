package rest_server

import (
	"encoding/json"
	"fmt"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Nuffle-Labs/nffl/aggregator/mocks"
	"github.com/Nuffle-Labs/nffl/core/types/messages"
	"github.com/Nuffle-Labs/nffl/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	aggtypes "github.com/Nuffle-Labs/nffl/aggregator/types"
)

func TestGetStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := sdklogging.NewNoopLogger()
	aggregator := mocks.NewMockRestAggregatorer(mockCtrl)
	restServer := NewRestServer("", aggregator, logger)

	msg := messages.StateRootUpdateMessage{
		RollupId:            1,
		BlockHeight:         2,
		Timestamp:           3,
		NearDaCommitment:    tests.Keccak256(4),
		NearDaTransactionId: tests.Keccak256(5),
		StateRoot:           tests.Keccak256(6),
	}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	response := aggtypes.GetStateRootUpdateAggregationResponse{
		Message: msg,
		Aggregation: messages.MessageBlsAggregation{
			MessageDigest: msgDigest,
		},
	}
	aggregator.EXPECT().GetStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight).Return(&response, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/state-root-update?rollupId=%d&blockHeight=%d", msg.RollupId, msg.BlockHeight),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	err = restServer.handleGetStateRootUpdateAggregation(recorder, req)
	assert.Nil(t, err)
	assert.Equal(t, recorder.Code, http.StatusOK)

	var body aggtypes.GetStateRootUpdateAggregationResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)
	assert.Equal(t, body, response)
}

func TestGetOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	logger := sdklogging.NewNoopLogger()
	aggregator := mocks.NewMockRestAggregatorer(mockCtrl)
	restServer := NewRestServer("", aggregator, logger)

	msg := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	digest, err := msg.Digest()
	assert.Nil(t, err)

	response := aggtypes.GetOperatorSetUpdateAggregationResponse{
		Message: msg,
		Aggregation: messages.MessageBlsAggregation{
			MessageDigest: digest,
		},
	}

	aggregator.EXPECT().GetOperatorSetUpdateAggregation(msg.Id).Return(&response, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/aggregation/operator-set-update?id=%d", msg.Id),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	err = restServer.handleGetOperatorSetUpdateAggregation(recorder, req)
	assert.Nil(t, err)
	assert.Equal(t, recorder.Code, http.StatusOK)

	var actual aggtypes.GetOperatorSetUpdateAggregationResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &actual)
	assert.Nil(t, err)
	assert.Equal(t, response, actual)
}

func TestGetCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	logger := sdklogging.NewNoopLogger()
	aggregator := mocks.NewMockRestAggregatorer(mockCtrl)
	restServer := NewRestServer("", aggregator, logger)

	stateRootMessage := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
	}
	stateRootDigest, err := stateRootMessage.Digest()
	assert.Nil(t, err)
	stateRootAggregation := messages.MessageBlsAggregation{
		MessageDigest: stateRootDigest,
	}

	operatorSetMesssage := messages.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	operatorSetDigest, err := operatorSetMesssage.Digest()
	assert.Nil(t, err)
	operatorSetAggregation := messages.MessageBlsAggregation{
		MessageDigest: operatorSetDigest,
	}

	var fromTimestamp, toTimestamp uint64 = 0, 3
	response := aggtypes.GetCheckpointMessagesResponse{
		CheckpointMessages: messages.CheckpointMessages{
			StateRootUpdateMessages:              []messages.StateRootUpdateMessage{stateRootMessage},
			StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{stateRootAggregation},
			OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{operatorSetMesssage},
			OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{operatorSetAggregation},
		},
	}
	aggregator.EXPECT().GetCheckpointMessages(fromTimestamp, toTimestamp).Return(&response, nil)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/checkpoint/messages?fromTimestamp=%d&toTimestamp=%d", fromTimestamp, toTimestamp),
		nil,
	)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	err = restServer.handleGetCheckpointMessages(recorder, req)
	assert.Nil(t, err)
	assert.Equal(t, recorder.Code, http.StatusOK)

	var actual aggtypes.GetCheckpointMessagesResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &actual)
	assert.Nil(t, err)
	assert.Equal(t, response, actual)
}
