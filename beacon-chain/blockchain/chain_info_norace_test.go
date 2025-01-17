package blockchain

import (
	"context"
	"testing"

	testDB "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/testing"
	doublylinkedtree "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func TestHeadSlot_DataRace(t *testing.T) {
	beaconDB := testDB.SetupDB(t)
	s := &Service{
		cfg: &config{BeaconDB: beaconDB},
	}
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	st, _ := util.DeterministicGenesisState(t, 1)
	wait := make(chan struct{})
	go func() {
		defer close(wait)
		require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))
	}()
	s.HeadSlot()
	<-wait
}

func TestHeadRoot_DataRace(t *testing.T) {
	beaconDB := testDB.SetupDB(t)
	s := &Service{
		cfg:  &config{BeaconDB: beaconDB, StateGen: stategen.New(beaconDB, doublylinkedtree.New())},
		head: &head{root: [32]byte{'A'}},
	}
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	wait := make(chan struct{})
	st, _ := util.DeterministicGenesisState(t, 1)
	go func() {
		defer close(wait)
		require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))

	}()
	_, err = s.HeadRoot(context.Background())
	require.NoError(t, err)
	<-wait
}

func TestHeadBlock_DataRace(t *testing.T) {
	beaconDB := testDB.SetupDB(t)
	wsb, err := blocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlock{Block: &ethpb.BeaconBlock{Body: &ethpb.BeaconBlockBody{}}})
	require.NoError(t, err)
	s := &Service{
		cfg:  &config{BeaconDB: beaconDB, StateGen: stategen.New(beaconDB, doublylinkedtree.New())},
		head: &head{block: wsb},
	}
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	wait := make(chan struct{})
	st, _ := util.DeterministicGenesisState(t, 1)
	go func() {
		defer close(wait)
		require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))

	}()
	_, err = s.HeadBlock(context.Background())
	require.NoError(t, err)
	<-wait
}

func TestHeadState_DataRace(t *testing.T) {
	beaconDB := testDB.SetupDB(t)
	s := &Service{
		cfg: &config{BeaconDB: beaconDB, StateGen: stategen.New(beaconDB, doublylinkedtree.New())},
	}
	b, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlock())
	require.NoError(t, err)
	wait := make(chan struct{})
	st, _ := util.DeterministicGenesisState(t, 1)
	root := bytesutil.ToBytes32(bytesutil.PadTo([]byte{'s'}, 32))
	require.NoError(t, beaconDB.SaveGenesisBlockRoot(context.Background(), root))
	require.NoError(t, beaconDB.SaveState(context.Background(), st, root))
	go func() {
		defer close(wait)
		require.NoError(t, s.saveHead(context.Background(), [32]byte{}, b, st))

	}()
	_, err = s.HeadState(context.Background())
	require.NoError(t, err)
	<-wait
}
