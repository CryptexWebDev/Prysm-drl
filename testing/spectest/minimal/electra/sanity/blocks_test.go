package sanity

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/sanity"
)

func TestMinimal_Electra_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "sanity/blocks/pyspec_tests")
}
