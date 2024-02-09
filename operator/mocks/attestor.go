package mocks

import (
	"context"
	"github.com/NethermindEth/near-sffl/core/types"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	"github.com/NethermindEth/near-sffl/operator/attestor"
)

type MockAttestor struct {
	consumer   *MockConsumer
	blsKeypair *bls.KeyPair
	operatorId bls.OperatorId

	signedRootC chan types.SignedStateRootUpdateMessage
}

func NewMockAttestor(blsKeypair *bls.KeyPair, operatorId bls.OperatorId) *MockAttestor {
	consumer := NewMockConsumer()
	return &MockAttestor{
		blsKeypair:  blsKeypair,
		operatorId:  operatorId,
		consumer:    consumer,
		signedRootC: make(chan types.SignedStateRootUpdateMessage),
	}
}

func (mockAttestor *MockAttestor) Start(ctx context.Context) error {
	go func() {
		mqBlockC := mockAttestor.consumer.GetBlockStream()
		for {
			mqBlock := <-mqBlockC

			message := servicemanager.StateRootUpdateMessage{
				RollupId:    mqBlock.RollupId,
				BlockHeight: mqBlock.Block.Header().Number.Uint64(),
				Timestamp:   mqBlock.Block.Header().Time,
				StateRoot:   mqBlock.Block.Header().Root,
			}
			signature, err := attestor.SignStateRootUpdateMessage(mockAttestor.blsKeypair, &message)
			if err != nil {
				panic(err)
			}

			signedStateRootUpdateMessage := types.SignedStateRootUpdateMessage{
				Message:      message,
				BlsSignature: *signature,
				OperatorId:   mockAttestor.operatorId,
			}

			mockAttestor.signedRootC <- signedStateRootUpdateMessage
		}
	}()

	return nil
}

func (mockAttestor *MockAttestor) Close() error { return mockAttestor.consumer.Close() }

func (mockAttestor *MockAttestor) GetSingedRootC() <-chan types.SignedStateRootUpdateMessage {
	return mockAttestor.signedRootC
}

func (mockAttestor *MockAttestor) MockGetConsumer() *MockConsumer {
	return mockAttestor.consumer
}
