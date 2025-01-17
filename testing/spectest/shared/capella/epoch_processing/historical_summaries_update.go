package epoch_processing

import (
	"path"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/epoch"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/utils"
)

// RunHistoricalSummariesUpdateTests executes "epoch_processing/historical_summaries_update" tests.
func RunHistoricalSummariesUpdateTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "capella", "epoch_processing/historical_summaries_update/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "capella", "epoch_processing/historical_summaries_update/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processHistoricalSummariesUpdateWrapper)
		})
	}
}

func processHistoricalSummariesUpdateWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessHistoricalDataUpdate(st)
	require.NoError(t, err, "Could not process historical summary updates")
	return st, nil
}
