package random

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/sanity"
)

func TestMinimal_Electra_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}
