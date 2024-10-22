package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/phase0/finality"
)

func TestMainnet_Phase0_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
