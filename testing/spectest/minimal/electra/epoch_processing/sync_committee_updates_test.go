package epoch_processing

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/epoch_processing"
)

func TestMinimal_Electra_EpochProcessing_SyncCommitteeUpdates(t *testing.T) {
	epoch_processing.RunSyncCommitteeUpdatesTests(t, "minimal")
}
