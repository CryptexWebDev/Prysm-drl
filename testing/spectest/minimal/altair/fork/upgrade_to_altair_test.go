package fork

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/altair/fork"
)

func TestMinimal_Altair_UpgradeToAltair(t *testing.T) {
	fork.RunUpgradeToAltair(t, "minimal")
}
