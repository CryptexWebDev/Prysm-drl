package fork_transition

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/fork"
)

func TestMainnet_Electra_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
