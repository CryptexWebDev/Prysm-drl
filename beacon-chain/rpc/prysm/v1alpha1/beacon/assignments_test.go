package beacon

import (
	"context"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"

	mock "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain/testing"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/helpers"
	dbTest "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/testing"
	doublylinkedtree "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	mockstategen "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen/mock"
	"github.com/Dorol-Chain/Prysm-drl/v5/cmd"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
	"github.com/Dorol-Chain/Prysm-drl/v5/time/slots"
)

func TestServer_ListAssignments_CannotRequestFutureEpoch(t *testing.T) {
	db := dbTest.SetupDB(t)
	ctx := context.Background()
	bs := &Server{
		BeaconDB:           db,
		GenesisTimeFetcher: &mock.ChainService{},
	}
	addDefaultReplayerBuilder(bs, db)

	wanted := errNoEpochInfoError
	_, err := bs.ListValidatorAssignments(
		ctx,
		&ethpb.ListValidatorAssignmentsRequest{
			QueryFilter: &ethpb.ListValidatorAssignmentsRequest_Epoch{
				Epoch: slots.ToEpoch(bs.GenesisTimeFetcher.CurrentSlot()) + 1,
			},
		},
	)
	assert.ErrorContains(t, wanted, err)
}

func TestServer_ListAssignments_Pagination_InputOutOfRange(t *testing.T) {
	helpers.ClearCache()
	db := dbTest.SetupDB(t)
	ctx := context.Background()
	count := 100
	validators := make([]*ethpb.Validator, 0, count)
	for i := 0; i < count; i++ {
		pubKey := make([]byte, params.BeaconConfig().BLSPubkeyLength)
		withdrawalCred := make([]byte, 32)
		binary.LittleEndian.PutUint64(pubKey, uint64(i))
		validators = append(validators, &ethpb.Validator{
			PublicKey:             pubKey,
			WithdrawalCredentials: withdrawalCred,
			ExitEpoch:             params.BeaconConfig().FarFutureEpoch,
			EffectiveBalance:      params.BeaconConfig().MaxEffectiveBalance,
			ActivationEpoch:       0,
		})
	}

	blk := util.NewBeaconBlock()
	blockRoot, err := blk.Block.HashTreeRoot()
	require.NoError(t, err)

	s, err := util.NewBeaconState()
	require.NoError(t, err)
	require.NoError(t, s.SetValidators(validators))
	require.NoError(t, db.SaveState(ctx, s, blockRoot))
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, blockRoot))

	bs := &Server{
		BeaconDB: db,
		HeadFetcher: &mock.ChainService{
			State: s,
		},
		FinalizationFetcher: &mock.ChainService{
			FinalizedCheckPoint: &ethpb.Checkpoint{
				Epoch: 0,
			},
		},
		GenesisTimeFetcher: &mock.ChainService{},
		StateGen:           stategen.New(db, doublylinkedtree.New()),
		ReplayerBuilder:    mockstategen.NewReplayerBuilder(mockstategen.WithMockState(s)),
	}

	wanted := fmt.Sprintf("page start %d >= list %d", 500, count)
	_, err = bs.ListValidatorAssignments(context.Background(), &ethpb.ListValidatorAssignmentsRequest{
		PageToken:   strconv.Itoa(2),
		QueryFilter: &ethpb.ListValidatorAssignmentsRequest_Genesis{Genesis: true},
	})
	assert.ErrorContains(t, wanted, err)
}

func TestServer_ListAssignments_Pagination_ExceedsMaxPageSize(t *testing.T) {
	bs := &Server{}
	exceedsMax := int32(cmd.Get().MaxRPCPageSize + 1)

	wanted := fmt.Sprintf("Requested page size %d can not be greater than max size %d", exceedsMax, cmd.Get().MaxRPCPageSize)
	req := &ethpb.ListValidatorAssignmentsRequest{
		PageToken: strconv.Itoa(0),
		PageSize:  exceedsMax,
	}
	_, err := bs.ListValidatorAssignments(context.Background(), req)
	assert.ErrorContains(t, wanted, err)
}

func TestServer_ListAssignments_Pagination_DefaultPageSize_NoArchive(t *testing.T) {
	helpers.ClearCache()
	db := dbTest.SetupDB(t)
	ctx := context.Background()
	count := 500
	validators := make([]*ethpb.Validator, 0, count)
	for i := 0; i < count; i++ {
		pubKey := make([]byte, params.BeaconConfig().BLSPubkeyLength)
		withdrawalCred := make([]byte, 32)
		binary.LittleEndian.PutUint64(pubKey, uint64(i))
		// Mark the validators with index divisible by 3 inactive.
		if i%3 == 0 {
			validators = append(validators, &ethpb.Validator{
				PublicKey:             pubKey,
				WithdrawalCredentials: withdrawalCred,
				ExitEpoch:             0,
				ActivationEpoch:       0,
				EffectiveBalance:      params.BeaconConfig().MaxEffectiveBalance,
			})
		} else {
			validators = append(validators, &ethpb.Validator{
				PublicKey:             pubKey,
				WithdrawalCredentials: withdrawalCred,
				ExitEpoch:             params.BeaconConfig().FarFutureEpoch,
				EffectiveBalance:      params.BeaconConfig().MaxEffectiveBalance,
				ActivationEpoch:       0,
			})
		}
	}

	b := util.NewBeaconBlock()
	blockRoot, err := b.Block.HashTreeRoot()
	require.NoError(t, err)

	s, err := util.NewBeaconState()
	require.NoError(t, err)
	require.NoError(t, s.SetValidators(validators))
	require.NoError(t, db.SaveState(ctx, s, blockRoot))
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, blockRoot))

	bs := &Server{
		BeaconDB: db,
		HeadFetcher: &mock.ChainService{
			State: s,
		},
		FinalizationFetcher: &mock.ChainService{
			FinalizedCheckPoint: &ethpb.Checkpoint{
				Epoch: 0,
			},
		},
		GenesisTimeFetcher: &mock.ChainService{},
		StateGen:           stategen.New(db, doublylinkedtree.New()),
		ReplayerBuilder:    mockstategen.NewReplayerBuilder(mockstategen.WithMockState(s)),
	}

	res, err := bs.ListValidatorAssignments(context.Background(), &ethpb.ListValidatorAssignmentsRequest{
		QueryFilter: &ethpb.ListValidatorAssignmentsRequest_Genesis{Genesis: true},
	})
	require.NoError(t, err)

	// Construct the wanted assignments.
	var wanted []*ethpb.ValidatorAssignments_CommitteeAssignment

	activeIndices, err := helpers.ActiveValidatorIndices(ctx, s, 0)
	require.NoError(t, err)
	assignments, err := helpers.CommitteeAssignments(context.Background(), s, 0, activeIndices[0:params.BeaconConfig().DefaultPageSize])
	require.NoError(t, err)
	proposerSlots, err := helpers.ProposerAssignments(ctx, s, 0)
	require.NoError(t, err)
	for _, index := range activeIndices[0:params.BeaconConfig().DefaultPageSize] {
		val, err := s.ValidatorAtIndex(index)
		require.NoError(t, err)
		wanted = append(wanted, &ethpb.ValidatorAssignments_CommitteeAssignment{
			BeaconCommittees: assignments[index].Committee,
			CommitteeIndex:   assignments[index].CommitteeIndex,
			AttesterSlot:     assignments[index].AttesterSlot,
			ProposerSlots:    proposerSlots[index],
			PublicKey:        val.PublicKey,
			ValidatorIndex:   index,
		})
	}
	assert.DeepSSZEqual(t, wanted, res.Assignments, "Did not receive wanted assignments")
}

func TestServer_ListAssignments_FilterPubkeysIndices_NoPagination(t *testing.T) {
	helpers.ClearCache()
	db := dbTest.SetupDB(t)

	ctx := context.Background()
	count := 100
	validators := make([]*ethpb.Validator, 0, count)
	withdrawCreds := make([]byte, 32)
	for i := 0; i < count; i++ {
		pubKey := make([]byte, params.BeaconConfig().BLSPubkeyLength)
		binary.LittleEndian.PutUint64(pubKey, uint64(i))
		val := &ethpb.Validator{
			PublicKey:             pubKey,
			WithdrawalCredentials: withdrawCreds,
			ExitEpoch:             params.BeaconConfig().FarFutureEpoch,
		}
		validators = append(validators, val)
	}

	b := util.NewBeaconBlock()
	blockRoot, err := b.Block.HashTreeRoot()
	require.NoError(t, err)
	s, err := util.NewBeaconState()
	require.NoError(t, err)
	require.NoError(t, s.SetValidators(validators))
	require.NoError(t, db.SaveState(ctx, s, blockRoot))
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, blockRoot))

	bs := &Server{
		BeaconDB: db,
		FinalizationFetcher: &mock.ChainService{
			FinalizedCheckPoint: &ethpb.Checkpoint{
				Epoch: 0,
			},
		},
		GenesisTimeFetcher: &mock.ChainService{},
		StateGen:           stategen.New(db, doublylinkedtree.New()),
		ReplayerBuilder:    mockstategen.NewReplayerBuilder(mockstategen.WithMockState(s)),
	}

	pubKey1 := make([]byte, params.BeaconConfig().BLSPubkeyLength)
	binary.LittleEndian.PutUint64(pubKey1, 1)
	pubKey2 := make([]byte, params.BeaconConfig().BLSPubkeyLength)
	binary.LittleEndian.PutUint64(pubKey2, 2)
	req := &ethpb.ListValidatorAssignmentsRequest{PublicKeys: [][]byte{pubKey1, pubKey2}, Indices: []primitives.ValidatorIndex{2, 3}}
	res, err := bs.ListValidatorAssignments(context.Background(), req)
	require.NoError(t, err)

	// Construct the wanted assignments.
	var wanted []*ethpb.ValidatorAssignments_CommitteeAssignment

	activeIndices, err := helpers.ActiveValidatorIndices(ctx, s, 0)
	require.NoError(t, err)
	assignments, err := helpers.CommitteeAssignments(context.Background(), s, 0, activeIndices[1:4])
	require.NoError(t, err)
	proposerSlots, err := helpers.ProposerAssignments(ctx, s, 0)
	require.NoError(t, err)
	for _, index := range activeIndices[1:4] {
		val, err := s.ValidatorAtIndex(index)
		require.NoError(t, err)
		wanted = append(wanted, &ethpb.ValidatorAssignments_CommitteeAssignment{
			BeaconCommittees: assignments[index].Committee,
			CommitteeIndex:   assignments[index].CommitteeIndex,
			AttesterSlot:     assignments[index].AttesterSlot,
			ProposerSlots:    proposerSlots[index],
			PublicKey:        val.PublicKey,
			ValidatorIndex:   index,
		})
	}

	assert.DeepEqual(t, wanted, res.Assignments, "Did not receive wanted assignments")
}

func TestServer_ListAssignments_CanFilterPubkeysIndices_WithPagination(t *testing.T) {
	helpers.ClearCache()
	db := dbTest.SetupDB(t)
	ctx := context.Background()
	count := 100
	validators := make([]*ethpb.Validator, 0, count)
	withdrawCred := make([]byte, 32)
	for i := 0; i < count; i++ {
		pubKey := make([]byte, params.BeaconConfig().BLSPubkeyLength)
		binary.LittleEndian.PutUint64(pubKey, uint64(i))
		val := &ethpb.Validator{
			PublicKey:             pubKey,
			WithdrawalCredentials: withdrawCred,
			ExitEpoch:             params.BeaconConfig().FarFutureEpoch,
		}
		validators = append(validators, val)
	}

	b := util.NewBeaconBlock()
	blockRoot, err := b.Block.HashTreeRoot()
	require.NoError(t, err)
	s, err := util.NewBeaconState()
	require.NoError(t, err)
	util.SaveBlock(t, ctx, db, b)
	require.NoError(t, s.SetValidators(validators))
	require.NoError(t, db.SaveState(ctx, s, blockRoot))
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, blockRoot))

	bs := &Server{
		BeaconDB: db,
		FinalizationFetcher: &mock.ChainService{
			FinalizedCheckPoint: &ethpb.Checkpoint{
				Epoch: 0,
			},
		},
		GenesisTimeFetcher: &mock.ChainService{},
		StateGen:           stategen.New(db, doublylinkedtree.New()),
	}

	addDefaultReplayerBuilder(bs, db)

	req := &ethpb.ListValidatorAssignmentsRequest{Indices: []primitives.ValidatorIndex{1, 2, 3, 4, 5, 6}, PageSize: 2, PageToken: "1"}
	res, err := bs.ListValidatorAssignments(context.Background(), req)
	require.NoError(t, err)

	// Construct the wanted assignments.
	var assignments []*ethpb.ValidatorAssignments_CommitteeAssignment

	activeIndices, err := helpers.ActiveValidatorIndices(ctx, s, 0)
	require.NoError(t, err)
	as, err := helpers.CommitteeAssignments(context.Background(), s, 0, activeIndices[3:5])
	require.NoError(t, err)
	proposalSlots, err := helpers.ProposerAssignments(ctx, s, 0)
	require.NoError(t, err)
	for _, index := range activeIndices[3:5] {
		val, err := s.ValidatorAtIndex(index)
		require.NoError(t, err)
		assignments = append(assignments, &ethpb.ValidatorAssignments_CommitteeAssignment{
			BeaconCommittees: as[index].Committee,
			CommitteeIndex:   as[index].CommitteeIndex,
			AttesterSlot:     as[index].AttesterSlot,
			ProposerSlots:    proposalSlots[index],
			PublicKey:        val.PublicKey,
			ValidatorIndex:   index,
		})
	}

	wantedRes := &ethpb.ValidatorAssignments{
		Assignments:   assignments,
		TotalSize:     int32(len(req.Indices)),
		NextPageToken: "2",
	}

	assert.DeepEqual(t, wantedRes, res, "Did not get wanted assignments")

	// Test the wrap around scenario.
	assignments = nil
	req = &ethpb.ListValidatorAssignmentsRequest{Indices: []primitives.ValidatorIndex{1, 2, 3, 4, 5, 6}, PageSize: 5, PageToken: "1"}
	res, err = bs.ListValidatorAssignments(context.Background(), req)
	require.NoError(t, err)
	as, err = helpers.CommitteeAssignments(context.Background(), s, 0, activeIndices[6:7])
	require.NoError(t, err)
	proposalSlots, err = helpers.ProposerAssignments(ctx, s, 0)
	require.NoError(t, err)
	for _, index := range activeIndices[6:7] {
		val, err := s.ValidatorAtIndex(index)
		require.NoError(t, err)
		assignments = append(assignments, &ethpb.ValidatorAssignments_CommitteeAssignment{
			BeaconCommittees: as[index].Committee,
			CommitteeIndex:   as[index].CommitteeIndex,
			AttesterSlot:     as[index].AttesterSlot,
			ProposerSlots:    proposalSlots[index],
			PublicKey:        val.PublicKey,
			ValidatorIndex:   index,
		})
	}

	wantedRes = &ethpb.ValidatorAssignments{
		Assignments:   assignments,
		TotalSize:     int32(len(req.Indices)),
		NextPageToken: "",
	}

	assert.DeepEqual(t, wantedRes, res, "Did not receive wanted assignments")
}
