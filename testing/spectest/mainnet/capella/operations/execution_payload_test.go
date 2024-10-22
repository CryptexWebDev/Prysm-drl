package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "mainnet")
}
