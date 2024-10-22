package aggregator

//go:generate mockgen -destination=./mocks/rest_aggregator.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator RestAggregatorer
//go:generate mockgen -destination=./mocks/rpc_aggregator.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator RpcAggregatorer
//go:generate mockgen -destination=./mocks/message_blsagg.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator/blsagg MessageBlsAggregationService
//go:generate mockgen -destination=./mocks/rollup_broadcaster.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator RollupBroadcasterer
//go:generate mockgen -destination=./mocks/operator_registrations_inmemory.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator OperatorRegistrationsService
//go:generate mockgen -destination=./mocks/eth_client.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth Client
