package debug

import (
	"context"
	"fmt"
	"math"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/helpers"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/filters"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
	pbrpc "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1/attestation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetBlock in an ssz-encoded format by block root.
func (ds *Server) GetBlock(
	ctx context.Context,
	req *pbrpc.BlockRequestByRoot,
) (*pbrpc.SSZResponse, error) {
	root := bytesutil.ToBytes32(req.BlockRoot)
	signedBlock, err := ds.BeaconDB.Block(ctx, root)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not retrieve block by root: %v", err)
	}
	if signedBlock == nil || signedBlock.IsNil() {
		return &pbrpc.SSZResponse{Encoded: make([]byte, 0)}, nil
	}
	encoded, err := signedBlock.MarshalSSZ()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not marshal block: %v", err)
	}
	return &pbrpc.SSZResponse{
		Encoded: encoded,
	}, nil
}

// GetInclusionSlot of an attestation in block.
func (ds *Server) GetInclusionSlot(ctx context.Context, req *pbrpc.InclusionSlotRequest) (*pbrpc.InclusionSlotResponse, error) {
	ds.GenesisTimeFetcher.CurrentSlot()

	// Attestation has one epoch to get included in the chain. This blocks users from requesting too soon.
	epochBack := primitives.Slot(0)
	if ds.GenesisTimeFetcher.CurrentSlot() > params.BeaconConfig().SlotsPerEpoch {
		epochBack = ds.GenesisTimeFetcher.CurrentSlot() - params.BeaconConfig().SlotsPerEpoch
	}
	if epochBack < req.Slot {
		return nil, fmt.Errorf("attestation has one epoch window, please request slot older than %d", epochBack)
	}

	// Attestation could be in blocks between slot + 1 to slot + epoch_duration.
	startSlot := req.Slot + params.BeaconConfig().MinAttestationInclusionDelay
	endSlot := req.Slot + params.BeaconConfig().SlotsPerEpoch

	filter := filters.NewFilter().SetStartSlot(startSlot).SetEndSlot(endSlot)
	blks, _, err := ds.BeaconDB.Blocks(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not retrieve blocks: %v", err)
	}

	inclusionSlot := primitives.Slot(math.MaxUint64)
	targetStates := make(map[[32]byte]state.ReadOnlyBeaconState)
	for _, blk := range blks {
		for _, a := range blk.Block().Body().Attestations() {
			tr := bytesutil.ToBytes32(a.GetData().Target.Root)
			s, ok := targetStates[tr]
			if !ok {
				s, err = ds.StateGen.StateByRoot(ctx, tr)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Could not retrieve state: %v", err)
				}
				targetStates[tr] = s
			}
			c, err := helpers.BeaconCommitteeFromState(ctx, s, a.GetData().Slot, a.GetData().CommitteeIndex)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Could not get committee: %v", err)
			}
			indices, err := attestation.AttestingIndices(a, c)
			if err != nil {
				return nil, err
			}
			for _, i := range indices {
				if req.Id == i && req.Slot == a.GetData().Slot {
					inclusionSlot = blk.Block().Slot()
					break
				}
			}
		}
	}

	return &pbrpc.InclusionSlotResponse{Slot: inclusionSlot}, nil
}