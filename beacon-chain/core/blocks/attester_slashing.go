package blocks

import (
	"context"
	"sort"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/helpers"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/container/slice"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1/attestation"
	"github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1/slashings"
	"github.com/Dorol-Chain/Prysm-drl/v5/time/slots"
	"github.com/pkg/errors"
)

// ProcessAttesterSlashings is one of the operations performed
// on each processed beacon block to slash attesters based on
// Casper FFG slashing conditions if any slashable events occurred.
//
// Spec pseudocode definition:
//
//	def process_attester_slashing(state: BeaconState, attester_slashing: AttesterSlashing) -> None:
//	 attestation_1 = attester_slashing.attestation_1
//	 attestation_2 = attester_slashing.attestation_2
//	 assert is_slashable_attestation_data(attestation_1.data, attestation_2.data)
//	 assert is_valid_indexed_attestation(state, attestation_1)
//	 assert is_valid_indexed_attestation(state, attestation_2)
//
//	 slashed_any = False
//	 indices = set(attestation_1.attesting_indices).intersection(attestation_2.attesting_indices)
//	 for index in sorted(indices):
//	     if is_slashable_validator(state.validators[index], get_current_epoch(state)):
//	         slash_validator(state, index)
//	         slashed_any = True
//	 assert slashed_any
func ProcessAttesterSlashings(
	ctx context.Context,
	beaconState state.BeaconState,
	slashings []ethpb.AttSlashing,
	slashFunc slashValidatorFunc,
) (state.BeaconState, error) {
	var err error
	for _, slashing := range slashings {
		beaconState, err = ProcessAttesterSlashing(ctx, beaconState, slashing, slashFunc)
		if err != nil {
			return nil, err
		}
	}
	return beaconState, nil
}

// ProcessAttesterSlashing processes individual attester slashing.
func ProcessAttesterSlashing(
	ctx context.Context,
	beaconState state.BeaconState,
	slashing ethpb.AttSlashing,
	slashFunc slashValidatorFunc,
) (state.BeaconState, error) {
	if err := VerifyAttesterSlashing(ctx, beaconState, slashing); err != nil {
		return nil, errors.Wrap(err, "could not verify attester slashing")
	}
	slashableIndices := SlashableAttesterIndices(slashing)
	sort.SliceStable(slashableIndices, func(i, j int) bool {
		return slashableIndices[i] < slashableIndices[j]
	})
	currentEpoch := slots.ToEpoch(beaconState.Slot())
	var err error
	var slashedAny bool
	var val state.ReadOnlyValidator
	for _, validatorIndex := range slashableIndices {
		val, err = beaconState.ValidatorAtIndexReadOnly(primitives.ValidatorIndex(validatorIndex))
		if err != nil {
			return nil, err
		}
		if helpers.IsSlashableValidator(val.ActivationEpoch(), val.WithdrawableEpoch(), val.Slashed(), currentEpoch) {
			beaconState, err = slashFunc(ctx, beaconState, primitives.ValidatorIndex(validatorIndex))
			if err != nil {
				return nil, errors.Wrapf(err, "could not slash validator index %d",
					validatorIndex)
			}
			slashedAny = true
		}
	}
	if !slashedAny {
		return nil, errors.New("unable to slash any validator despite confirmed attester slashing")
	}
	return beaconState, nil
}

// VerifyAttesterSlashing validates the attestation data in both attestations in the slashing object.
func VerifyAttesterSlashing(ctx context.Context, beaconState state.ReadOnlyBeaconState, slashing ethpb.AttSlashing) error {
	if slashing == nil {
		return errors.New("nil slashing")
	}
	if slashing.FirstAttestation() == nil || slashing.SecondAttestation() == nil {
		return errors.New("nil attestation")
	}
	if slashing.FirstAttestation().GetData() == nil || slashing.SecondAttestation().GetData() == nil {
		return errors.New("nil attestation data")
	}
	att1 := slashing.FirstAttestation()
	att2 := slashing.SecondAttestation()
	data1 := att1.GetData()
	data2 := att2.GetData()
	if !IsSlashableAttestationData(data1, data2) {
		return errors.New("attestations are not slashable")
	}
	if err := VerifyIndexedAttestation(ctx, beaconState, att1); err != nil {
		return errors.Wrap(err, "could not validate indexed attestation")
	}
	if err := VerifyIndexedAttestation(ctx, beaconState, att2); err != nil {
		return errors.Wrap(err, "could not validate indexed attestation")
	}
	return nil
}

// IsSlashableAttestationData verifies a slashing against the Casper Proof of Stake FFG rules.
//
// Spec pseudocode definition:
//
//	def is_slashable_attestation_data(data_1: AttestationData, data_2: AttestationData) -> bool:
//	 """
//	 Check if ``data_1`` and ``data_2`` are slashable according to Casper FFG rules.
//	 """
//	 return (
//	     # Double vote
//	     (data_1 != data_2 and data_1.target.epoch == data_2.target.epoch) or
//	     # Surround vote
//	     (data_1.source.epoch < data_2.source.epoch and data_2.target.epoch < data_1.target.epoch)
//	 )
func IsSlashableAttestationData(data1, data2 *ethpb.AttestationData) bool {
	if data1 == nil || data2 == nil || data1.Target == nil || data2.Target == nil || data1.Source == nil || data2.Source == nil {
		return false
	}
	isDoubleVote := !attestation.AttDataIsEqual(data1, data2) && data1.Target.Epoch == data2.Target.Epoch
	att1 := &ethpb.IndexedAttestation{Data: data1}
	att2 := &ethpb.IndexedAttestation{Data: data2}
	// Check if att1 is surrounding att2.
	isSurroundVote := slashings.IsSurround(att1, att2)
	return isDoubleVote || isSurroundVote
}

// SlashableAttesterIndices returns the intersection of attester indices from both attestations in this slashing.
func SlashableAttesterIndices(slashing ethpb.AttSlashing) []uint64 {
	if slashing == nil || slashing.FirstAttestation() == nil || slashing.SecondAttestation() == nil {
		return nil
	}
	indices1 := slashing.FirstAttestation().GetAttestingIndices()
	indices2 := slashing.SecondAttestation().GetAttestingIndices()
	return slice.IntersectionUint64(indices1, indices2)
}