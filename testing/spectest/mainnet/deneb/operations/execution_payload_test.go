package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "mainnet")
}
