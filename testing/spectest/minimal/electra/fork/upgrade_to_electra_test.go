package fork

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/fork"
)

func TestMinimal_UpgradeToElectra(t *testing.T) {
	fork.RunUpgradeToElectra(t, "minimal")
}
