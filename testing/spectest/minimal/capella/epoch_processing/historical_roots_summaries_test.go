package epoch_processing

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/epoch_processing"
)

func TestMinimal_Capella_EpochProcessing_HistoricalSummariesUpdate(t *testing.T) {
	epoch_processing.RunHistoricalSummariesUpdateTests(t, "minimal")
}
