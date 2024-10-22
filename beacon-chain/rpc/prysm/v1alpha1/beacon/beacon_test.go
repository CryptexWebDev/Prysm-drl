package beacon

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/cmd/beacon-chain/flags"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
)

func TestMain(m *testing.M) {
	// Use minimal config to reduce test setup time.
	prevConfig := params.BeaconConfig().Copy()
	defer params.OverrideBeaconConfig(prevConfig)
	params.OverrideBeaconConfig(params.MinimalSpecConfig())

	resetFlags := flags.Get()
	flags.Init(&flags.GlobalFlags{
		MinimumSyncPeers: 30,
	})
	defer func() {
		flags.Init(resetFlags)
	}()

	m.Run()
}
