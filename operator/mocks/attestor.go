package mocks

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/NethermindEth/near-sffl/core"
	messages "github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/NethermindEth/near-sffl/operator/attestor"
)

type MockAttestor struct {
	consumer   *MockConsumer
	blsKeypair *bls.KeyPair
	operatorId eigentypes.OperatorId

	signedRootC chan messages.SignedStateRootUpdateMessage
}

var _ core.Metricable = (*MockAttestor)(nil)

func NewMockAttestor(blsKeypair *bls.KeyPair, operatorId eigentypes.OperatorId) *MockAttestor {
	consumer := NewMockConsumer()
	return &MockAttestor{
		blsKeypair:  blsKeypair,
		operatorId:  operatorId,
		consumer:    consumer,
		signedRootC: make(chan messages.SignedStateRootUpdateMessage),
	}
}

func (mockAttestor *MockAttestor) WithMetrics(_ *prometheus.Registry) error {
	return nil
}

func (mockAttestor *MockAttestor) Start(ctx context.Context) error {
	go func() {
		mqBlockC := mockAttestor.consumer.GetBlockStream()
		for {
			mqBlock := <-mqBlockC

			message := messages.StateRootUpdateMessage{
				RollupId:    mqBlock.RollupId,
				BlockHeight: mqBlock.Block.Header().Number.Uint64(),
				Timestamp:   mqBlock.Block.Header().Time,
				StateRoot:   mqBlock.Block.Header().Root,
				// TODO: get below fields from mqBlock
				NearDaTransactionId: [32]byte{1},
				NearDaCommitment:    [32]byte{2},
			}
			signature, err := attestor.SignStateRootUpdateMessage(mockAttestor.blsKeypair, &message)
			if err != nil {
				panic(err)
			}

			signedStateRootUpdateMessage := messages.SignedStateRootUpdateMessage{
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

func (mockAttestor *MockAttestor) GetSignedRootC() <-chan messages.SignedStateRootUpdateMessage {
	return mockAttestor.signedRootC
}

func (mockAttestor *MockAttestor) MockGetConsumer() *MockConsumer {
	return mockAttestor.consumer
}
