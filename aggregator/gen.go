package aggregator

//go:generate mockgen -destination=./mocks/message_blsagg.go -package=mocks github.com/NethermindEth/near-sffl/aggregator MessageBlsAggregationService
//go:generate mockgen -destination=./mocks/message_database.go -package=mocks github.com/NethermindEth/near-sffl/aggregator MessageDatabaser
//go:generate mockgen -destination=./mocks/rollup_broadcaster.go -package=mocks github.com/NethermindEth/near-sffl/aggregator RollupBroadcasterer
