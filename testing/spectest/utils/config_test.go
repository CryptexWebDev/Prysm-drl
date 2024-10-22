package utils

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func TestConfig(t *testing.T) {
	require.NoError(t, SetConfig(t, "minimal"))
	require.Equal(t, primitives.Slot(8), params.BeaconConfig().SlotsPerEpoch)
	require.NoError(t, SetConfig(t, "mainnet"))
	require.Equal(t, primitives.Slot(32), params.BeaconConfig().SlotsPerEpoch)
}
