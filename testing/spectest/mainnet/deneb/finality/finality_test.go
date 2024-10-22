package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/finality"
)

func TestMainnet_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}
