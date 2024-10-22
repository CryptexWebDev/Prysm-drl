package fork_transition

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/altair/fork"
)

func TestMinimal_Altair_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
