package fork

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/fork"
)

func TestMinimal_Capella_UpgradeToCapella(t *testing.T) {
	fork.RunUpgradeToCapella(t, "minimal")
}
