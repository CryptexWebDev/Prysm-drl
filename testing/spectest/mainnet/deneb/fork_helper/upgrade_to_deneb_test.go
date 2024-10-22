package fork_helper

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/fork"
)

func TestMainnet_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "mainnet")
}
