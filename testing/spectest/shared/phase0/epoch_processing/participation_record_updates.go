package epoch_processing

import (
	"path"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/epoch"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/utils"
)

// RunParticipationRecordUpdatesTests executes "epoch_processing/participation_record_updates" tests.
func RunParticipationRecordUpdatesTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "phase0", "epoch_processing/participation_record_updates/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "phase0", "epoch_processing/participation_record_updates/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			RunEpochOperationTest(t, folderPath, processParticipationRecordUpdatesWrapper)
		})
	}
}

func processParticipationRecordUpdatesWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessParticipationRecordUpdates(st)
	require.NoError(t, err, "Could not process final updates")
	return st, nil
}
