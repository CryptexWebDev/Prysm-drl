package forkchoice

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func TestBuilderTick(t *testing.T) {
	st, err := util.NewBeaconState()
	require.NoError(t, err)
	blk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	builder := NewBuilder(t, st, blk)
	builder.Tick(t, 10)

	require.Equal(t, int64(10), builder.lastTick)
}

func TestBuilderInvalidBlock(t *testing.T) {
	st, err := util.NewBeaconState()
	require.NoError(t, err)
	blk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	builder := NewBuilder(t, st, blk)
	blk, err = blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	blk.SetSlot(2)
	require.NoError(t, err)
	builder.InvalidBlock(t, blk)
}

func TestPoWBlock(t *testing.T) {
	st, err := util.NewBeaconState()
	require.NoError(t, err)
	blk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	builder := NewBuilder(t, st, blk)
	builder.PoWBlock(&ethpb.PowBlock{BlockHash: []byte{1, 2, 3}})

	require.Equal(t, 1, len(builder.execMock.powBlocks))
}
