package beacon

import (
	"context"

	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/container/slice"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SubmitProposerSlashing receives a proposer slashing object via
// RPC and injects it into the beacon node's operations pool.
// Submission into this pool does not guarantee inclusion into a beacon block.
func (bs *Server) SubmitProposerSlashing(
	ctx context.Context,
	req *ethpb.ProposerSlashing,
) (*ethpb.SubmitSlashingResponse, error) {
	beaconState, err := bs.HeadFetcher.HeadStateReadOnly(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not retrieve head state: %v", err)
	}
	if err := bs.SlashingsPool.InsertProposerSlashing(ctx, beaconState, req); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert proposer slashing into pool: %v", err)
	}
	if !features.Get().DisableBroadcastSlashings {
		if err := bs.Broadcaster.Broadcast(ctx, req); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not broadcast slashing object: %v", err)
		}
	}

	return &ethpb.SubmitSlashingResponse{
		SlashedIndices: []primitives.ValidatorIndex{req.Header_1.Header.ProposerIndex},
	}, nil
}

func (bs *Server) SubmitAttesterSlashing(ctx context.Context, req *ethpb.AttesterSlashing) (*ethpb.SubmitSlashingResponse, error) {
	return bs.submitAttesterSlashing(ctx, req)
}

// SubmitAttesterSlashingElectra receives an attester slashing object via
// RPC and injects it into the beacon node's operations pool.
// Submission into this pool does not guarantee inclusion into a beacon block.
func (bs *Server) SubmitAttesterSlashingElectra(ctx context.Context, req *ethpb.AttesterSlashingElectra) (*ethpb.SubmitSlashingResponse, error) {
	return bs.submitAttesterSlashing(ctx, req)
}

func (bs *Server) submitAttesterSlashing(ctx context.Context, slashing ethpb.AttSlashing) (*ethpb.SubmitSlashingResponse, error) {
	beaconState, err := bs.HeadFetcher.HeadStateReadOnly(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not retrieve head state: %v", err)
	}
	if err := bs.SlashingsPool.InsertAttesterSlashing(ctx, beaconState, slashing); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not insert attester slashing into pool: %v", err)
	}
	if !features.Get().DisableBroadcastSlashings {
		if err := bs.Broadcaster.Broadcast(ctx, slashing); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not broadcast slashing object: %v", err)
		}
	}
	indices := slice.IntersectionUint64(slashing.FirstAttestation().GetAttestingIndices(), slashing.SecondAttestation().GetAttestingIndices())
	slashedIndices := make([]primitives.ValidatorIndex, len(indices))
	for i, index := range indices {
		slashedIndices[i] = primitives.ValidatorIndex(index)
	}
	return &ethpb.SubmitSlashingResponse{
		SlashedIndices: slashedIndices,
	}, nil
}
