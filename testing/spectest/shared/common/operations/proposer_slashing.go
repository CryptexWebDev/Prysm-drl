package operations

import (
	"context"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/blocks"
	v "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/validators"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
)

func RunProposerSlashingTest(t *testing.T, config string, fork string, block blockWithSSZObject, sszToState SSZToState) {
	runSlashingTest(t, config, fork, "proposer_slashing", block, sszToState, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
		return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), v.SlashValidator)
	})
}
