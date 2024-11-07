deploy-regular:
	forge script Deploy.s.sol --rpc-url https://holesky.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META
	forge script Deploy.s.sol --rpc-url https://polygon-amoy.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META
	forge script Deploy.s.sol --rpc-url https://arbitrum-sepolia.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META

deploy-simple:
	forge script DeploySimple.s.sol --rpc-url https://holesky.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META
	forge script DeploySimple.s.sol --rpc-url https://polygon-amoy.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META
	forge script DeploySimple.s.sol --rpc-url https://arbitrum-sepolia.infura.io/v3/a7d4a3dd6f774049bce5d61651549421 --broadcast --account META

deploy: deploy-simple deploy-regular

