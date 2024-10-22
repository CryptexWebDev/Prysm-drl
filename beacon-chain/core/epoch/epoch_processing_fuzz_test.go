package epoch

import (
	"testing"

	state_native "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/state-native"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	fuzz "github.com/google/gofuzz"
)

func TestFuzzFinalUpdates_10000(t *testing.T) {
	fuzzer := fuzz.NewWithSeed(0)
	base := &ethpb.BeaconState{}

	for i := 0; i < 10000; i++ {
		fuzzer.Fuzz(base)
		s, err := state_native.InitializeFromProtoUnsafePhase0(base)
		require.NoError(t, err)
		_, err = ProcessFinalUpdates(s)
		_ = err
	}
}
