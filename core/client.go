package core

import (
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
)

func createEthClient(rpcUrl string, retryCount int, retryInterval time.Duration, reinitializeInterval time.Duration, logger sdklogging.Logger) (eth.Client, <-chan eth.Client, func(error), error) {
	clientC := make(chan eth.Client)

	reinitializeTicker := time.NewTicker(reinitializeInterval)
	reinitializeTicker.Stop()

	reinitializeClient := func() {
		logger.Info("Reinitializing eth client")
		client, err := eth.NewClient(rpcUrl)
		if err != nil {
			logger.Error("Error reinitializing eth client", "err", err)
		} else {
			reinitializeTicker.Stop()
			clientC <- client
		}
	}

	client, err := eth.NewClient(rpcUrl)
	if err != nil {
		logger.Error("Error creating initial eth client", "err", err)
		return nil, nil, nil, err
	}

	onClientError := func(err error) {
		logger.Error("onClientError triggered", "err", err)

		for i := 0; i < retryCount; i++ {
			<-time.After(retryInterval)
			newClient, err := eth.NewClient(rpcUrl)
			if err == nil {
				logger.Info("Eth client recovered after retry")
				clientC <- newClient
				return
			}
		}

		reinitializeTicker.Reset(reinitializeInterval)
	}

	go func() {
		for range reinitializeTicker.C {
			reinitializeClient()
		}
	}()

	return client, clientC, onClientError, nil
}
