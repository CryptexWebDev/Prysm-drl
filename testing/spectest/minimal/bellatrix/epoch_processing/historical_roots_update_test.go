package epoch_processing

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/bellatrix/epoch_processing"
)

func TestMinimal_Bellatrix_EpochProcessing_HistoricalRootsUpdate(t *testing.T) {
	epoch_processing.RunHistoricalRootsUpdateTests(t, "minimal")
}
