package blockchain

import (
	"context"
	"testing"

	testDB "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/testing"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func TestService_getBlock(t *testing.T) {
	ctx := context.Background()
	beaconDB := testDB.SetupDB(t)
	s := setupBeaconChain(t, beaconDB)
	b1 := util.NewBeaconBlock()
	r1, err := b1.Block.HashTreeRoot()
	require.NoError(t, err)
	b2 := util.NewBeaconBlock()
	b2.Block.Slot = 100
	r2, err := b2.Block.HashTreeRoot()
	require.NoError(t, err)

	// block not found
	_, err = s.getBlock(ctx, [32]byte{})
	require.ErrorIs(t, err, errBlockNotFoundInCacheOrDB)

	// block in cache
	b, err := blocks.NewSignedBeaconBlock(b1)
	require.NoError(t, err)
	require.NoError(t, s.saveInitSyncBlock(ctx, r1, b))
	got, err := s.getBlock(ctx, r1)
	require.NoError(t, err)
	require.DeepEqual(t, b, got)

	// block in db
	b = util.SaveBlock(t, ctx, s.cfg.BeaconDB, b2)
	got, err = s.getBlock(ctx, r2)
	require.NoError(t, err)
	require.DeepEqual(t, b, got)
}

func TestService_hasBlockInInitSyncOrDB(t *testing.T) {
	ctx := context.Background()
	beaconDB := testDB.SetupDB(t)
	s := setupBeaconChain(t, beaconDB)
	b1 := util.NewBeaconBlock()
	r1, err := b1.Block.HashTreeRoot()
	require.NoError(t, err)
	b2 := util.NewBeaconBlock()
	b2.Block.Slot = 100
	r2, err := b2.Block.HashTreeRoot()
	require.NoError(t, err)

	// block not found
	require.Equal(t, false, s.hasBlockInInitSyncOrDB(ctx, [32]byte{}))

	// block in cache
	b, err := blocks.NewSignedBeaconBlock(b1)
	require.NoError(t, err)
	require.NoError(t, s.saveInitSyncBlock(ctx, r1, b))
	require.Equal(t, true, s.hasBlockInInitSyncOrDB(ctx, r1))

	// block in db
	util.SaveBlock(t, ctx, s.cfg.BeaconDB, b2)
	require.Equal(t, true, s.hasBlockInInitSyncOrDB(ctx, r2))
}
