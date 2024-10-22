package shuffle

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/phase0/shuffling/core/shuffle"
)

func TestMinimal_Phase0_Shuffling_Core_Shuffle(t *testing.T) {
	shuffle.RunShuffleTests(t, "minimal")
}
