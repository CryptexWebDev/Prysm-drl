package fork

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/fork"
)

func TestMinimal_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "minimal")
}
