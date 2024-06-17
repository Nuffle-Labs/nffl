package rest_server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/aggregator"
	"github.com/NethermindEth/near-sffl/core"
)

var (
	errorToCode = map[error]int{
		aggregator.StateRootUpdateNotFoundError: http.StatusNotFound,
		aggregator.StateRootAggNotFoundError:    http.StatusNotFound,
		aggregator.OperatorSetNotFoundError:     http.StatusNotFound,
		aggregator.OperatorAggNotFoundError:     http.StatusNotFound,
	}
)

type RestServer struct {
	serverIpPortAddr string
	app              aggregator.RestAggregatorer

	logger   logging.Logger
	listener EventListener
}

var _ core.Metricable = (*RestServer)(nil)

func NewRestServer(serverIpPortAddr string, app aggregator.RestAggregatorer, logger logging.Logger) *RestServer {
	return &RestServer{
		serverIpPortAddr: serverIpPortAddr,
		app:              app,
		logger:           logger,
		listener:         &SelectiveListener{},
	}
}

func (s *RestServer) EnableMetrics(registry *prometheus.Registry) error {
	listener, err := MakeServerMetrics(registry)
	if err != nil {
		return err
	}

	s.listener = listener
	return nil
}

func (s *RestServer) Start() error {
	s.logger.Info("Starting aggregator REST API.")

	router := mux.NewRouter()
	router.HandleFunc("/aggregation/state-root-update", wrapRequest(s.listener.APIErrors, s.handleGetStateRootUpdateAggregation)).Methods("GET")
	router.HandleFunc("/aggregation/operator-set-update", wrapRequest(s.listener.APIErrors, s.handleGetOperatorSetUpdateAggregation)).Methods("GET")
	router.HandleFunc("/checkpoint/messages", wrapRequest(s.listener.APIErrors, s.handleGetCheckpointMessages)).Methods("GET")

	err := http.ListenAndServe(s.serverIpPortAddr, router)
	if err != nil {
		s.logger.Fatal("ListenAndServe", "err", err)
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

func (s *RestServer) handleGetStateRootUpdateAggregation(w http.ResponseWriter, r *http.Request) error {
	s.listener.IncStateRootUpdateRequests()

	params := r.URL.Query()
	rollupId, err := strconv.ParseUint(params.Get("rollupId"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid rollupId", http.StatusBadRequest)
		return err
	}

	blockHeight, err := strconv.ParseUint(params.Get("blockHeight"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid blockHeight", http.StatusBadRequest)
		return err
	}

	response, err := s.app.GetStateRootUpdateAggregation(uint32(rollupId), blockHeight)
	if err != nil {
		http.Error(w, err.Error(), mapErrorToCode(err))
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(*response)
}

func (s *RestServer) handleGetOperatorSetUpdateAggregation(w http.ResponseWriter, r *http.Request) error {
	s.listener.IncOperatorSetUpdateRequests()

	params := r.URL.Query()
	id, err := strconv.ParseUint(params.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return err
	}

	response, err := s.app.GetOperatorSetUpdateAggregation(id)
	if err != nil {
		http.Error(w, err.Error(), mapErrorToCode(err))
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(*response)
}

func (s *RestServer) handleGetCheckpointMessages(w http.ResponseWriter, r *http.Request) error {
	s.listener.IncCheckpointMessagesRequests()
	params := r.URL.Query()

	fromTimestamp, err := strconv.ParseUint(params.Get("fromTimestamp"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid fromTimestamp", http.StatusBadRequest)
		return err
	}

	toTimestamp, err := strconv.ParseUint(params.Get("toTimestamp"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid toTimestamp", http.StatusBadRequest)
		return err
	}

	response, err := s.app.GetCheckpointMessages(fromTimestamp, toTimestamp)
	if err != nil {
		http.Error(w, err.Error(), mapErrorToCode(err))
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(*response)
}

func mapErrorToCode(err error) int {
	status, ok := errorToCode[err]
	if !ok {
		return http.StatusInternalServerError
	}

	return status
}
