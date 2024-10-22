package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/finality"
)

func TestMainnet_Capella_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
