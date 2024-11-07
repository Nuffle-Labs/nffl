.PHONY: deploy wire e2e e2e-all e2e-simple e2e-simple-all

deploy:
	npx hardhat lz:deploy
	@echo "================================================================"
	@echo "=== Deployed the OApps for testing (both regular and simple) ==="
	@echo "================================================================"

wire:
	npx hardhat lz:oapp:wire --oapp-config layerzero.config.ts
	npx hardhat lz:oapp:wire --oapp-config layerzerosimple.config.ts
	@echo "============================================================="
	@echo "=== Wired the OApps for testing (both regular and simple) ==="
	@echo "============================================================="

e2e:
	npx hardhat e2e --source holesky --target arbitrum-sepolia

e2e-all: e2e
	npx hardhat e2e --source holesky --target polygon-amoy
	npx hardhat e2e --source arbitrum-sepolia --target polygon-amoy

e2e-simple:
	npx hardhat e2e:simple --source holesky --target arbitrum-sepolia

e2e-simple-all: e2e-simple
	npx hardhat e2e:simple --source holesky --target polygon-amoy
	npx hardhat e2e:simple --source arbitrum-sepolia --target polygon-amoy
