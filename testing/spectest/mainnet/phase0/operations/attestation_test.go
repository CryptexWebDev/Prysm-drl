package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}
