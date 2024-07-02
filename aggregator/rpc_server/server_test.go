package rpc_server

import (
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/NethermindEth/near-sffl/aggregator/mocks"
	"github.com/NethermindEth/near-sffl/core/types/messages"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProcessSignedCheckpointTaskResponse_Invalid_Params(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	agg := mocks.NewMockRpcAggregatorer(mockCtrl)
	logger, _ := logging.NewZapLogger(logging.Development)

	rpc := NewRpcServer("localhost:8080", agg, logger)

	var ignore bool
	t.Run("nil message", func(t *testing.T) {
		err := rpc.ProcessSignedCheckpointTaskResponse(nil, &ignore)

		assert.NotNil(t, err)
	})

	t.Run("nil signature", func(t *testing.T) {
		err := rpc.ProcessSignedCheckpointTaskResponse(&messages.SignedCheckpointTaskResponse{
			BlsSignature: bls.Signature{G1Point: nil},
		}, &ignore)

		assert.NotNil(t, err)
	})
}
