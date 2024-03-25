package aggregator

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
	"github.com/NethermindEth/near-sffl/core/types/messages"
)

func (agg *Aggregator) startRestServer() error {
	router := mux.NewRouter()
	router.HandleFunc("/aggregation/state-root-update", agg.handleGetStateRootUpdateAggregation).Methods("GET")
	router.HandleFunc("/aggregation/operator-set-update", agg.handleGetOperatorSetUpdateAggregation).Methods("GET")
	router.HandleFunc("/checkpoint/messages", agg.handleGetCheckpointMessages).Methods("GET")

	err := http.ListenAndServe(agg.restServerIpPortAddr, router)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

func (agg *Aggregator) handleGetStateRootUpdateAggregation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	rollupId, err := strconv.ParseUint(params.Get("rollupId"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid rollupId", http.StatusBadRequest)
		return
	}

	blockHeight, err := strconv.ParseUint(params.Get("blockHeight"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid blockHeight", http.StatusBadRequest)
		return
	}

	var message messages.StateRootUpdateMessage
	var aggregation messages.MessageBlsAggregation

	err = agg.msgDb.FetchStateRootUpdate(uint32(rollupId), blockHeight, &message)
	if err != nil {
		http.Error(w, "StateRootUpdate not found", http.StatusNotFound)
		return
	}

	err = agg.msgDb.FetchStateRootUpdateAggregation(uint32(rollupId), blockHeight, &aggregation)
	if err != nil {
		http.Error(w, "StateRootUpdate aggregation not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggtypes.GetStateRootUpdateAggregationResponse{
		Message:     message,
		Aggregation: aggregation,
	})
}

func (agg *Aggregator) handleGetOperatorSetUpdateAggregation(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.ParseUint(params.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var message messages.OperatorSetUpdateMessage
	var aggregation messages.MessageBlsAggregation

	err = agg.msgDb.FetchOperatorSetUpdate(id, &message)
	if err != nil {
		http.Error(w, "OperatorSetUpdate not found", http.StatusNotFound)
		return
	}

	err = agg.msgDb.FetchOperatorSetUpdateAggregation(id, &aggregation)
	if err != nil {
		http.Error(w, "OperatorSetUpdate aggregation not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggtypes.GetOperatorSetUpdateAggregationResponse{
		Message:     message,
		Aggregation: aggregation,
	})
}

func (agg *Aggregator) handleGetCheckpointMessages(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	fromTimestamp, err := strconv.ParseUint(params.Get("fromTimestamp"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid fromTimestamp", http.StatusBadRequest)
		return
	}

	toTimestamp, err := strconv.ParseUint(params.Get("toTimestamp"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid toTimestamp", http.StatusBadRequest)
		return
	}

	var checkpointMessages messages.CheckpointMessages

	agg.msgDb.FetchCheckpointMessages(fromTimestamp, toTimestamp, &checkpointMessages)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggtypes.GetCheckpointMessagesResponse{
		CheckpointMessages: checkpointMessages,
	})
}
