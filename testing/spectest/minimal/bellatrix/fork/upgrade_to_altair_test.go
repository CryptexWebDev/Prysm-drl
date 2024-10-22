package fork

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/bellatrix/fork"
)

func TestMinimal_Bellatrix_UpgradeToBellatrix(t *testing.T) {
	fork.RunUpgradeToBellatrix(t, "minimal")
}
