package epoch_processing

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/epoch_processing"
)

func TestMainnet_Electra_EpochProcessing_PendingConsolidations(t *testing.T) {
	epoch_processing.RunPendingConsolidationsTests(t, "mainnet")
}
