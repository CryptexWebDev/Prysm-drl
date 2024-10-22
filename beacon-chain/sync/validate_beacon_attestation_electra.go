package sync

import (
	"context"
	"fmt"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/monitoring/tracing/trace"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// validateCommitteeIndexElectra implements the following checks from the spec:
//   - [REJECT] len(committee_indices) == 1, where committee_indices = get_committee_indices(attestation).
//   - [REJECT] attestation.data.index == 0
func validateCommitteeIndexElectra(ctx context.Context, a ethpb.Att) (primitives.CommitteeIndex, pubsub.ValidationResult, error) {
	_, span := trace.StartSpan(ctx, "sync.validateCommitteeIndexElectra")
	defer span.End()
	_, ok := a.(*ethpb.AttestationElectra)
	if !ok {
		return 0, pubsub.ValidationIgnore, fmt.Errorf("attestation has wrong type (expected %T, got %T)", &ethpb.AttestationElectra{}, a)
	}
	committeeIndex, err := a.GetCommitteeIndex()
	if err != nil {
		return 0, pubsub.ValidationReject, err
	}
	return committeeIndex, pubsub.ValidationAccept, nil
}
