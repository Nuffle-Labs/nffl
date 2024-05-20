package safeclient

//go:generate mockgen -destination=./mocks/safe_client.go -package=mocks github.com/NethermindEth/near-sffl/core/safeclient SafeClient
//go:generate mockgen -destination=./mocks/eth_client.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/eth Client
//go:generate mockgen -destination=./mocks/subscription.go -package=mocks github.com/ethereum/go-ethereum Subscription
