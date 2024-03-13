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
	registryrollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/core"
	"github.com/NethermindEth/near-sffl/core/types"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

func TestGetStateRootUpdateAggregation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, _, _, _, _, _, mockDb, _, err := createMockAggregator(mockCtrl, MOCK_OPERATOR_PUBKEY_DICT)
	assert.Nil(t, err)

	go aggregator.startRestServer()

	msg := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
		StateRoot:   keccak256(4),
	}
	msgDigest, err := core.GetStateRootUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	aggregation := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	mockDb.EXPECT().FetchStateRootUpdate(msg.RollupId, msg.BlockHeight, gomock.Any()).DoAndReturn(
		func(rollupId coretypes.RollupId, blockHeight uint64, msgPtr *servicemanager.StateRootUpdateMessage) error {
			if rollupId != msg.RollupId || blockHeight != msg.BlockHeight {
				return errors.New("Unexpected args")
			}

			*msgPtr = msg

			return nil
		},
	)

	mockDb.EXPECT().FetchStateRootUpdateAggregation(msg.RollupId, msg.BlockHeight, gomock.Any()).DoAndReturn(
		func(rollupId coretypes.RollupId, blockHeight uint64, aggPtr *aggtypes.MessageBlsAggregationServiceResponse) error {
			if rollupId != msg.RollupId || blockHeight != msg.BlockHeight {
				return errors.New("Unexpected args")
			}

			*aggPtr = aggregation

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

	expectedBody := GetStateRootUpdateAggregationResponse{
		Message:     msg,
		Aggregation: aggregation,
	}
	var body GetStateRootUpdateAggregationResponse

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

	msg := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	msgDigest, err := core.GetOperatorSetUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	aggregation := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	mockDb.EXPECT().FetchOperatorSetUpdate(msg.Id, gomock.Any()).DoAndReturn(
		func(id uint64, msgPtr *registryrollup.OperatorSetUpdateMessage) error {
			if id != msg.Id {
				return errors.New("Unexpected args")
			}

			*msgPtr = msg

			return nil
		},
	)

	mockDb.EXPECT().FetchOperatorSetUpdateAggregation(msg.Id, gomock.Any()).DoAndReturn(
		func(id uint64, aggPtr *aggtypes.MessageBlsAggregationServiceResponse) error {
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

	expectedBody := GetOperatorSetUpdateAggregationResponse{
		Message:     msg,
		Aggregation: aggregation,
	}
	var body GetOperatorSetUpdateAggregationResponse

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

	msg := servicemanager.StateRootUpdateMessage{
		RollupId:    1,
		BlockHeight: 2,
		Timestamp:   3,
	}
	msgDigest, err := core.GetStateRootUpdateMessageDigest(&msg)
	assert.Nil(t, err)

	aggregation := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest,
	}

	msg2 := registryrollup.OperatorSetUpdateMessage{
		Id:        1,
		Timestamp: 2,
	}
	msgDigest2, err := core.GetOperatorSetUpdateMessageDigest(&msg2)
	assert.Nil(t, err)

	aggregation2 := aggtypes.MessageBlsAggregationServiceResponse{
		MessageDigest: msgDigest2,
	}

	mockDb.EXPECT().FetchCheckpointMessages(uint64(0), uint64(3), gomock.Any()).DoAndReturn(
		func(fromTimestamp uint64, toTimestamp uint64, result *types.CheckpointMessages) error {
			*result = types.CheckpointMessages{
				StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{msg},
				StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{aggregation},
				OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{msg2},
				OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{aggregation2},
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

	expectedBody := GetCheckpointMessagesResponse{
		CheckpointMessages: types.CheckpointMessages{
			StateRootUpdateMessages:              []servicemanager.StateRootUpdateMessage{msg},
			StateRootUpdateMessageAggregations:   []aggtypes.MessageBlsAggregationServiceResponse{aggregation},
			OperatorSetUpdateMessages:            []registryrollup.OperatorSetUpdateMessage{msg2},
			OperatorSetUpdateMessageAggregations: []aggtypes.MessageBlsAggregationServiceResponse{aggregation2},
		},
	}
	var body GetCheckpointMessagesResponse

	assert.Equal(t, recorder.Code, http.StatusOK)

	if recorder.Code != http.StatusOK {
		fmt.Printf("HTTP Error: %s", recorder.Body.Bytes())
	}

	err = json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Nil(t, err)

	assert.Equal(t, body, expectedBody)
}
