package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/altair/finality"
)

func TestMinimal_Altair_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
