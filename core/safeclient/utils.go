package safeclient

import (
	"crypto/sha256"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/ethereum/go-ethereum/core/types"
)

func createDefaultClient(rpcUrl string, logger logging.Logger) (eth.Client, error) {
	client, err := eth.NewClient(rpcUrl)
	if err != nil {
		return nil, err
	}
	logger.Debug("Created new eth client without collector")
	return client, nil
}

func createInstrumentedClient(rpcUrl string, collector *rpccalls.Collector, logger logging.Logger) (eth.Client, error) {
	client, err := eth.NewInstrumentedClient(rpcUrl, collector)
	if err != nil {
		return nil, err
	}
	logger.Debug("Created new instrumented eth client with collector")
	return client, nil
}

func hashLog(log *types.Log) [32]byte {
	h := sha256.New()

	log.EncodeRLP(h)

	// EncodeRLP only serializes the address, topics and data, so adding some additional block and tx info
	h.Write(log.BlockHash.Bytes())
	h.Write(log.TxHash.Bytes())
	h.Write(new(big.Int).SetUint64(uint64(log.Index)).Bytes())

	return [32]byte(h.Sum(nil))
}
