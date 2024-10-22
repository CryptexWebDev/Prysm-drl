package finality

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/finality"
)

func TestMinimal_Deneb_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}
