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

	"github.com/NethermindEth/near-sffl/aggregator/types"
	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

func TestGetStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	msg := messages.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
	}
	msgDigest, err := msg.Digest()
	assert.Nil(t, err)

	aggregation := types.MessageBlsAggregationServiceResponse{
		MessageBlsAggregation: messages.MessageBlsAggregation{
			MessageDigest: msgDigest,
		},
	}

	mockDb.EXPECT().FetchStateRootUpdate(msg.RollupId, msg.BlockHeight, gomock.Any()).DoAndReturn(
		func(rollupId coretypes.RollupId, blockHeight uint64, msgPtr *messages.StateRootUpdateMessage) error {
			if rollupId != msg.RollupId || blockHeight != msg.BlockHeight {
				return errors.New("Unexpected args")
			}

			*msgPtr = msg

			return nil
		},
	)

	mockDb.EXPECT().FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, gomock.Any()).DoAndReturn(
		func(rollupId coretypes.RollupId, blockHeight uint64, aggPtr *messages.MessageBlsAggregation) error {
			if rollupId != msg.RollupId || blockHeight != msg.BlockHeight {
				return errors.New("Unexpected args")
			}

			*aggPtr = aggregation.MessageBlsAggregation

			return nil
		},
	)

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

	assert.Equal(t, recorder.Code, http.StatusOK)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, body, expectedBody)
}

func TestGetOperatorSetUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
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

	mockDb.EXPECT().FetchOperatorSetUpdate(msg.Id, gomock.Any()).DoAndReturn(
		func(id uint64, msgPtr *messages.OperatorSetUpdateMessage) error {
			if id != msg.Id {
				return errors.New("Unexpected args")
			}

			*msgPtr = msg

			return nil
		},
	)

	mockDb.EXPECT().FetchOperatorSetUpdateAggregation(msg.Id, gomock.Any()).DoAndReturn(
		func(id uint64, aggPtr *messages.MessageBlsAggregation) error {
			if id != msg.Id {
				return errors.New("Unexpected args")
			}

			*aggPtr = aggregation

			return nil
		},
	)

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

	assert.Equal(t, recorder.Code, http.StatusOK)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, body, expectedBody)
}

func TestGetCheckpointMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
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

	mockDb.EXPECT().FetchCheckpointMessages(uint64(0), uint64(3), gomock.Any()).DoAndReturn(
		func(fromTimestamp uint64, toTimestamp uint64, result *messages.CheckpointMessages) error {
			*result = messages.CheckpointMessages{
				StateRootUpdateMessages:              []messages.StateRootUpdateMessage{msg},
				StateRootUpdateMessageAggregations:   []messages.MessageBlsAggregation{aggregation},
				OperatorSetUpdateMessages:            []messages.OperatorSetUpdateMessage{msg2},
				OperatorSetUpdateMessageAggregations: []messages.MessageBlsAggregation{aggregation2},
			}

			return nil
		},
	)

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

	assert.Equal(t, recorder.Code, http.StatusOK)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, body, expectedBody)
}
