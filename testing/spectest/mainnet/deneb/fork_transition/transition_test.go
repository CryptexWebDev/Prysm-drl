package fork_transition

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/fork"
)

func TestMainnet_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}
