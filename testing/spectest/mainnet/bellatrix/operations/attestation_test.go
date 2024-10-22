package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
