package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/altair/finality"
)

func TestMainnet_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
