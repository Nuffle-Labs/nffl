package operator

//go:generate mockgen -destination=./mocks/rpc_client.go -package=mocks github.com/NethermindEth/near-sffl/operator AggregatorRpcClienter
//go:generate mockgen -destination=./mocks/consumer.go -package=mocks github.com/NethermindEth/near-sffl/operator Consumerer
