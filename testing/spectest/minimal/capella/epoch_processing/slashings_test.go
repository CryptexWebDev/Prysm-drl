package epoch_processing

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/epoch_processing"
)

func TestMinimal_Capella_EpochProcessing_Slashings(t *testing.T) {
	epoch_processing.RunSlashingsTests(t, "minimal")
}
