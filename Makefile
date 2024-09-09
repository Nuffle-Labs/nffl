############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

OPERATOR_BLS_KEY_PASS=fDUMDLmBROwlzzPXyIcy
OPERATOR_ECDSA_KEY_PASS=EnJuncq01CiVk9UbuBYl
AGGREGATOR_ECDSA_PRIV_KEY=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
CHALLENGER_ECDSA_PRIV_KEY=0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a

INDEXER_NEAR_ENV=localnet
INDEXER_NEAR_HELPER_ACCOUNT=near
INDEXER_NEAR_CLI_LOCALNET_KEY_PATH=${HOME}/.near/localnet/validator_key.json

CHAINID=31337
DEPLOYMENT_FILES_DIR=contracts/evm/script/output/${CHAINID}

-----------------------------: ##

___CONTRACTS___: ##

deploy-eigenlayer-contracts-to-anvil-and-save-state: ## Deploy eigenlayer
	./tests/anvil/deploy-eigenlayer-save-anvil-state.sh

deploy-sffl-contracts-to-anvil-and-save-state: ## Deploy avs
	./tests/anvil/deploy-avs-save-anvil-state.sh

deploy-rollup-sffl-contracts-to-anvil-and-save-state: ## Deploy rollup contracts
	./tests/anvil/deploy-rollup-avs-save-anvil-state.sh

deploy-all-to-anvil-and-save-state: deploy-eigenlayer-contracts-to-anvil-and-save-state deploy-sffl-contracts-to-anvil-and-save-state deploy-rollup-sffl-contracts-to-anvil-and-save-state ## deploy eigenlayer and avs contracts

start-anvil-chain-with-el-and-avs-deployed: ## starts anvil from a saved state file (with el and avs contracts deployed)
	./tests/anvil/start-anvil-chain-with-el-and-avs-deployed.sh

start-rollup-anvil-chain-with-avs-deployed: ## starts an anvil instance with the rollup avs contracts
	./tests/anvil/start-rollup-anvil-chain-with-avs-deployed.sh

setup-near-da: export NEAR_ENV=$(INDEXER_NEAR_ENV)
setup-near-da: export NEAR_HELPER_ACCOUNT=$(INDEXER_NEAR_HELPER_ACCOUNT)
setup-near-da: export NEAR_CLI_LOCALNET_KEY_PATH=$(INDEXER_NEAR_CLI_LOCALNET_KEY_PATH)
setup-near-da:
	near create-account da.test.near --masterAccount test.near
	near deploy da.test.near ./tests/near/near_da_blob_store.wasm --initFunction "new" --initArgs "{}" --masterAccount test.near

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh

___DOCKER___: ##
docker-build-indexer:
	docker build -t near-sffl-indexer -f ./indexer/Dockerfile .
docker-build-relayer:
	docker build -t near-sffl-test-relayer -f ./relayer/cmd/Dockerfile .
docker-build-aggregator:
	docker build -t near-sffl-aggregator -f ./aggregator/cmd/Dockerfile .
docker-build-operator:
	docker build -t near-sffl-operator -f ./operator/cmd/Dockerfile .
docker-build-plugin:
	docker build -t near-sffl-operator-plugin -f ./plugin/cmd/Dockerfile .
docker-build-images: docker-build-indexer docker-build-relayer docker-build-aggregator docker-build-operator docker-build-plugin ## builds and publishes indexer, operator and aggregator docker images
docker-start-everything: docker-build-images ## starts aggregator and operator docker containers
	docker compose up

__CLI__: ##

cli-setup-operator: export OPERATOR_BLS_KEY_PASSWORD=$(OPERATOR_BLS_KEY_PASS)
cli-setup-operator: export OPERATOR_ECDSA_KEY_PASSWORD=$(OPERATOR_ECDSA_KEY_PASS)
cli-setup-operator: send-fund cli-register-operator-with-eigenlayer cli-register-operator-with-avs ## registers operator with eigenlayer and avs

cli-register-operator-with-eigenlayer: ## registers operator with delegationManager
	go run cli/main.go --config config-files/operator.anvil.yaml register-operator-with-eigenlayer

cli-register-operator-with-avs: ##
	go run cli/main.go --config config-files/operator.anvil.yaml register-operator-with-avs

cli-deregister-operator-with-avs: ##
	go run cli/main.go --config config-files/operator.anvil.yaml deregister-operator-with-avs

cli-print-operator-status: ##
	go run cli/main.go --config config-files/operator.anvil.yaml print-operator-status

send-fund: ## sends fund to the first operator saved in tests/keys/ecdsa/*
	cast send 0xD5A0359da7B310917d7760385516B2426E86ab7f --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6 --rpc-url http://127.0.0.1:8545

-----------------------------: ##
# We pipe all zapper logs through https://github.com/maoueh/zap-pretty so make sure to install it
# TODO: piping to zap-pretty only works when zapper environment is set to production, unsure why
____OFFCHAIN_SOFTWARE___: ##
start-aggregator: ##
	go run aggregator/cmd/main.go --config config-files/aggregator.yaml \
		--sffl-deployment ${DEPLOYMENT_FILES_DIR}/sffl_avs_deployment_output.json \
		--ecdsa-private-key ${AGGREGATOR_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

start-operator: export OPERATOR_BLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy
start-operator: export OPERATOR_ECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl
start-operator: ##
	go run operator/cmd/main.go --config config-files/operator.anvil.yaml \
		2>&1 | zap-pretty

start-indexer: ##
	cargo run -p indexer --release -- --home-dir ~/.near/localnet init --chain-id localnet
	cargo run -p indexer --release -- --home-dir ~/.near/localnet run --da-contract-ids da.test.near --rollup-ids 2 --rmq-address "amqp://127.0.0.1:5672"

start-test-relayer: ##
	CGO_LDFLAGS="-L ./relayer/libs ${CGO_LDFLAGS}" go run relayer/cmd/main.go run-args --rpc-url ws://127.0.0.1:8546 --da-account-id da.test.near --key-path ~/.near-credentials/localnet/da.test.near.json

run-plugin: ##
	go run plugin/cmd/main.go --config config-files/operator.anvil.yaml
-----------------------------: ##
_____HELPER_____: ##
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.3.0
	go generate ./...

tests-unit: ## runs all unit tests
	go test $$(go list ./... | grep -v /integration) -race -count=1 -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html

tests-contract: ## runs all forge tests
	cd contracts/evm && forge test --ffi

near-da-rpc-sys:
	rm -rf relayer/libs && \
	mkdir relayer/libs && \
	git clone https://github.com/taco-paco/rollup-data-availability.git && \
	cd rollup-data-availability && \
	git checkout c9ec12924b27e37b8c40e7ab1a051a64b363cfd6 && \
	make da-rpc-sys && \
	cp gopkg/da-rpc/lib/* ../relayer/libs && \
	cd .. && \
	rm -rf rollup-data-availability

# TODO: Currently we cannot use the race detector with `integration_test.go`
tests-integration: ## runs all integration tests
	go test ./tests/integration/integration_test.go -v -count=1
	go test ./tests/integration/registration_test.go -v -race -count=1

## runs slither for solidity files
## You can install Slither by following the guide at https://github.com/crytic/slither/tree/master?tab=readme-ov-file#how-to-install
slither:
	cd contracts/evm && slither .

## runs linter on all files
## TODO: For now, only Go files are linted
lint:
	golangci-lint run
