package actions

import (
	"encoding/json"
	"log"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	optypes "github.com/NethermindEth/near-sffl/operator/types"
)

func readNodeConfig(configPath string) (*optypes.NodeConfig, error) {
	nodeConfig := optypes.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return nil, err
	}

	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	return &nodeConfig, nil
}
