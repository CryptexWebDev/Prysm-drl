package fork_helper

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/bellatrix/fork"
)

func TestMainnet_Bellatrix_UpgradeToBellatrix(t *testing.T) {
	fork.RunUpgradeToBellatrix(t, "mainnet")
}
