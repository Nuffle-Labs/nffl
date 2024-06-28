package aggregator

//go:generate mockgen -destination=./mocks/rest_aggregator.go -package=mocks github.com/NethermindEth/near-sffl/aggregator RestAggregatorer
//go:generate mockgen -destination=./mocks/message_blsagg.go -package=mocks github.com/NethermindEth/near-sffl/aggregator MessageBlsAggregationService
//go:generate mockgen -destination=./mocks/task_blsagg.go -package=mocks github.com/NethermindEth/near-sffl/aggregator TaskBlsAggregationService
//go:generate mockgen -destination=./mocks/rollup_broadcaster.go -package=mocks github.com/NethermindEth/near-sffl/aggregator RollupBroadcasterer
//go:generate mockgen -destination=./mocks/eth_client.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth Client
