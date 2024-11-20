package chainio

//go:generate mockgen -destination=./mocks/avs_subscriber.go -package=mocks github.com/Nuffle-Labs/nffl/core/chainio AvsSubscriberer
//go:generate mockgen -destination=./mocks/avs_writer.go -package=mocks github.com/Nuffle-Labs/nffl/core/chainio AvsWriterer
//go:generate mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/Nuffle-Labs/nffl/core/chainio AvsReaderer
