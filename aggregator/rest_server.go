package aggregator

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	aggtypes "github.com/NethermindEth/near-sffl/aggregator/types"
)

func (agg *Aggregator) startRestServer() error {
	router := mux.NewRouter()
	router.HandleFunc("/aggregation/state-root-update", wrapRequest(agg.restListener.APIErrors, agg.handleGetStateRootUpdateAggregation)).Methods("GET")
	router.HandleFunc("/aggregation/operator-set-update", wrapRequest(agg.restListener.APIErrors, agg.handleGetOperatorSetUpdateAggregation)).Methods("GET")
	router.HandleFunc("/checkpoint/messages", wrapRequest(agg.restListener.APIErrors, agg.handleGetCheckpointMessages)).Methods("GET")

	err := http.ListenAndServe(agg.restServerIpPortAddr, router)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

func wrapRequest(errorHandler func(), requestCallback func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	if errorHandler == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			_ = requestCallback(w, r)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := requestCallback(w, r)
		if err != nil {
			errorHandler()
		}
	}
}

func (agg *Aggregator) handleGetStateRootUpdateAggregation(w http.ResponseWriter, r *http.Request) error {
	agg.restListener.IncStateRootUpdateRequests()

	params := r.URL.Query()
	rollupId, err := strconv.ParseUint(params.Get("rollupId"), 10, 32)
	if err != nil {
		agg.logger.Error("Invalid rollupId", "params", params)
		http.Error(w, "Invalid rollupId", http.StatusBadRequest)
		return err
	}

	blockHeight, err := strconv.ParseUint(params.Get("blockHeight"), 10, 64)
	if err != nil {
		agg.logger.Error("Invalid blockHeight", "params", params)
		http.Error(w, "Invalid blockHeight", http.StatusBadRequest)
		return err
	}

	message, err := agg.msgDb.FetchStateRootUpdate(uint32(rollupId), blockHeight)
	if err != nil {
		agg.logger.Error("StateRootUpdate not found", "params", params)
		http.Error(w, "StateRootUpdate not found", http.StatusNotFound)
		return err
	}

	aggregation, err := agg.msgDb.FetchStateRootUpdateAggregation(uint32(rollupId), blockHeight)
	if err != nil {
		agg.logger.Error("StateRootUpdate aggregation not found", "params", params)
		http.Error(w, "StateRootUpdate aggregation not found", http.StatusNotFound)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(aggtypes.GetStateRootUpdateAggregationResponse{
		Message:     *message,
		Aggregation: *aggregation,
	})

	return err
}

func (agg *Aggregator) handleGetOperatorSetUpdateAggregation(w http.ResponseWriter, r *http.Request) error {
	agg.restListener.IncOperatorSetUpdateRequests()

	params := r.URL.Query()
	id, err := strconv.ParseUint(params.Get("id"), 10, 64)
	if err != nil {
		agg.logger.Error("Invalid id", "params", params)
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return err
	}

	message, err := agg.msgDb.FetchOperatorSetUpdate(id)
	if err != nil {
		agg.logger.Error("OperatorSetUpdate not found", "params", params)
		http.Error(w, "OperatorSetUpdate not found", http.StatusNotFound)
		return err
	}

	aggregation, err := agg.msgDb.FetchOperatorSetUpdateAggregation(id)
	if err != nil {
		agg.logger.Error("OperatorSetUpdate aggregation not found", "params", params)
		http.Error(w, "OperatorSetUpdate aggregation not found", http.StatusNotFound)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(aggtypes.GetOperatorSetUpdateAggregationResponse{
		Message:     *message,
		Aggregation: *aggregation,
	})

	return err
}

func (agg *Aggregator) handleGetCheckpointMessages(w http.ResponseWriter, r *http.Request) error {
	agg.restListener.IncCheckpointMessagesRequests()
	params := r.URL.Query()

	fromTimestamp, err := strconv.ParseUint(params.Get("fromTimestamp"), 10, 64)
	if err != nil {
		agg.logger.Error("Invalid fromTimestamp", "params", params)
		http.Error(w, "Invalid fromTimestamp", http.StatusBadRequest)
		return err
	}

	toTimestamp, err := strconv.ParseUint(params.Get("toTimestamp"), 10, 64)
	if err != nil {
		agg.logger.Error("Invalid toTimestamp", "params", params)
		http.Error(w, "Invalid toTimestamp", http.StatusBadRequest)
		return err
	}

	checkpointMessages, err := agg.msgDb.FetchCheckpointMessages(fromTimestamp, toTimestamp)
	if err != nil {
		agg.logger.Error("CheckpointMessages not found", "params", params)
		http.Error(w, "CheckpointMessages not found", http.StatusNotFound)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(aggtypes.GetCheckpointMessagesResponse{
		CheckpointMessages: *checkpointMessages,
	})

	return err
}
