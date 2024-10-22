package fork_helper

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/altair/fork"
)

func TestMainnet_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "mainnet")
}
