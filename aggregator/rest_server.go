package aggregator

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
)

func (agg *Aggregator) startRestServer(ctx context.Context) error {
	router := mux.NewRouter()
	router.HandleFunc("aggregation/state-root-update", agg.handleGetStateRootUpdateAggregation).Methods("GET")

	err := http.ListenAndServe(agg.restServerIpPortAddr, router)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type GetStateRootUpdateAggregationResponse struct {
	message     servicemanager.StateRootUpdateMessage
	aggregation types.MessageBlsAggregationServiceResponse
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

	var message servicemanager.StateRootUpdateMessage
	var aggregation types.MessageBlsAggregationServiceResponse

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
	json.NewEncoder(w).Encode(GetStateRootUpdateAggregationResponse{
		message:     message,
		aggregation: aggregation,
	})
}
