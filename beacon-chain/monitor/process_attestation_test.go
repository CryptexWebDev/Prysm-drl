package monitor

import (
	"bytes"
	"context"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
	"github.com/prysmaticlabs/go-bitfield"
	"github.com/sirupsen/logrus"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func TestGetAttestingIndices(t *testing.T) {
	ctx := context.Background()
	beaconState, _ := util.DeterministicGenesisState(t, 256)
	att := &ethpb.Attestation{
		Data: &ethpb.AttestationData{
			Slot:           1,
			CommitteeIndex: 0,
		},
		AggregationBits: bitfield.Bitlist{0b11, 0b1},
	}
	attestingIndices, err := attestingIndices(ctx, beaconState, att)
	require.NoError(t, err)
	require.DeepEqual(t, attestingIndices, []uint64{0xc, 0x2})

}

func TestProcessIncludedAttestationTwoTracked(t *testing.T) {
	hook := logTest.NewGlobal()
	s := setupService(t)
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	require.NoError(t, state.SetSlot(2))
	require.NoError(t, state.SetCurrentParticipationBits(bytes.Repeat([]byte{0xff}, 13)))

	att := &ethpb.Attestation{
		Data: &ethpb.AttestationData{
			Slot:            1,
			CommitteeIndex:  0,
			BeaconBlockRoot: bytesutil.PadTo([]byte("hello-world"), 32),
			Source: &ethpb.Checkpoint{
				Epoch: 0,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
			Target: &ethpb.Checkpoint{
				Epoch: 1,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
		},
		AggregationBits: bitfield.Bitlist{0b11, 0b1},
	}
	s.processIncludedAttestation(context.Background(), state, att)
	wanted1 := "\"Attestation included\" balanceChange=0 correctHead=true correctSource=true correctTarget=true head=0x68656c6c6f2d inclusionSlot=2 newBalance=32000000000 prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=2"
	wanted2 := "\"Attestation included\" balanceChange=100000000 correctHead=true correctSource=true correctTarget=true head=0x68656c6c6f2d inclusionSlot=2 newBalance=32000000000 prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=12"
	require.LogsContain(t, hook, wanted1)
	require.LogsContain(t, hook, wanted2)
}

func TestProcessUnaggregatedAttestationStateNotCached(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	hook := logTest.NewGlobal()
	ctx := context.Background()

	s := setupService(t)
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	require.NoError(t, state.SetSlot(2))
	header := state.LatestBlockHeader()
	participation := []byte{0xff, 0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	require.NoError(t, state.SetCurrentParticipationBits(participation))

	att := &ethpb.Attestation{
		Data: &ethpb.AttestationData{
			Slot:            1,
			CommitteeIndex:  0,
			BeaconBlockRoot: header.GetStateRoot(),
			Source: &ethpb.Checkpoint{
				Epoch: 0,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
			Target: &ethpb.Checkpoint{
				Epoch: 1,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
		},
		AggregationBits: bitfield.Bitlist{0b11, 0b1},
	}
	s.processUnaggregatedAttestation(ctx, att)
	require.LogsContain(t, hook, "Skipping unaggregated attestation due to state not found in cache")
	logrus.SetLevel(logrus.InfoLevel)
}

func TestProcessUnaggregatedAttestationStateCached(t *testing.T) {
	ctx := context.Background()
	hook := logTest.NewGlobal()

	s := setupService(t)
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	participation := []byte{0xff, 0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	require.NoError(t, state.SetCurrentParticipationBits(participation))

	var root [32]byte
	copy(root[:], "hello-world")

	att := &ethpb.Attestation{
		Data: &ethpb.AttestationData{
			Slot:            1,
			CommitteeIndex:  0,
			BeaconBlockRoot: root[:],
			Source: &ethpb.Checkpoint{
				Epoch: 0,
				Root:  root[:],
			},
			Target: &ethpb.Checkpoint{
				Epoch: 1,
				Root:  root[:],
			},
		},
		AggregationBits: bitfield.Bitlist{0b11, 0b1},
	}
	require.NoError(t, s.config.StateGen.SaveState(ctx, root, state))
	s.processUnaggregatedAttestation(context.Background(), att)
	wanted1 := "\"Processed unaggregated attestation\" head=0x68656c6c6f2d prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=2"
	wanted2 := "\"Processed unaggregated attestation\" head=0x68656c6c6f2d prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=12"
	require.LogsContain(t, hook, wanted1)
	require.LogsContain(t, hook, wanted2)
}

func TestProcessAggregatedAttestationStateNotCached(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	hook := logTest.NewGlobal()
	ctx := context.Background()

	s := setupService(t)
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	require.NoError(t, state.SetSlot(2))
	header := state.LatestBlockHeader()
	participation := []byte{0xff, 0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	require.NoError(t, state.SetCurrentParticipationBits(participation))

	att := &ethpb.AggregateAttestationAndProof{
		AggregatorIndex: 2,
		Aggregate: &ethpb.Attestation{
			Data: &ethpb.AttestationData{
				Slot:            1,
				CommitteeIndex:  0,
				BeaconBlockRoot: header.GetStateRoot(),
				Source: &ethpb.Checkpoint{
					Epoch: 0,
					Root:  bytesutil.PadTo([]byte("hello-world"), 32),
				},
				Target: &ethpb.Checkpoint{
					Epoch: 1,
					Root:  bytesutil.PadTo([]byte("hello-world"), 32),
				},
			},
			AggregationBits: bitfield.Bitlist{0b11, 0b1},
		},
	}
	s.processAggregatedAttestation(ctx, att)
	require.LogsContain(t, hook, "\"Processed attestation aggregation\" aggregatorIndex=2 beaconBlockRoot=0x000000000000 prefix=monitor slot=1 sourceRoot=0x68656c6c6f2d targetRoot=0x68656c6c6f2d")
	require.LogsContain(t, hook, "Skipping aggregated attestation due to state not found in cache")
	logrus.SetLevel(logrus.InfoLevel)
}

func TestProcessAggregatedAttestationStateCached(t *testing.T) {
	hook := logTest.NewGlobal()
	ctx := context.Background()
	s := setupService(t)
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	participation := []byte{0xff, 0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	require.NoError(t, state.SetCurrentParticipationBits(participation))

	var root [32]byte
	copy(root[:], "hello-world")

	att := &ethpb.AggregateAttestationAndProof{
		AggregatorIndex: 2,
		Aggregate: &ethpb.Attestation{
			Data: &ethpb.AttestationData{
				Slot:            1,
				CommitteeIndex:  0,
				BeaconBlockRoot: root[:],
				Source: &ethpb.Checkpoint{
					Epoch: 0,
					Root:  root[:],
				},
				Target: &ethpb.Checkpoint{
					Epoch: 1,
					Root:  root[:],
				},
			},
			AggregationBits: bitfield.Bitlist{0b10, 0b1},
		},
	}

	require.NoError(t, s.config.StateGen.SaveState(ctx, root, state))
	s.processAggregatedAttestation(ctx, att)
	require.LogsContain(t, hook, "\"Processed attestation aggregation\" aggregatorIndex=2 beaconBlockRoot=0x68656c6c6f2d prefix=monitor slot=1 sourceRoot=0x68656c6c6f2d targetRoot=0x68656c6c6f2d")
	require.LogsContain(t, hook, "\"Processed aggregated attestation\" head=0x68656c6c6f2d prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=2")
	require.LogsDoNotContain(t, hook, "\"Processed aggregated attestation\" head=0x68656c6c6f2d prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=12")
}

func TestProcessAttestations(t *testing.T) {
	hook := logTest.NewGlobal()
	s := setupService(t)
	ctx := context.Background()
	state, _ := util.DeterministicGenesisStateAltair(t, 256)
	require.NoError(t, state.SetSlot(2))
	require.NoError(t, state.SetCurrentParticipationBits(bytes.Repeat([]byte{0xff}, 13)))

	att := &ethpb.Attestation{
		Data: &ethpb.AttestationData{
			Slot:            1,
			CommitteeIndex:  0,
			BeaconBlockRoot: bytesutil.PadTo([]byte("hello-world"), 32),
			Source: &ethpb.Checkpoint{
				Epoch: 0,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
			Target: &ethpb.Checkpoint{
				Epoch: 1,
				Root:  bytesutil.PadTo([]byte("hello-world"), 32),
			},
		},
		AggregationBits: bitfield.Bitlist{0b11, 0b1},
	}

	block := &ethpb.BeaconBlockAltair{
		Slot: 2,
		Body: &ethpb.BeaconBlockBodyAltair{
			Attestations: []*ethpb.Attestation{att},
		},
	}

	wrappedBlock, err := blocks.NewBeaconBlock(block)
	require.NoError(t, err)
	s.processAttestations(ctx, state, wrappedBlock)
	wanted1 := "\"Attestation included\" balanceChange=0 correctHead=true correctSource=true correctTarget=true head=0x68656c6c6f2d inclusionSlot=2 newBalance=32000000000 prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=2"
	wanted2 := "\"Attestation included\" balanceChange=100000000 correctHead=true correctSource=true correctTarget=true head=0x68656c6c6f2d inclusionSlot=2 newBalance=32000000000 prefix=monitor slot=1 source=0x68656c6c6f2d target=0x68656c6c6f2d validatorIndex=12"
	require.LogsContain(t, hook, wanted1)
	require.LogsContain(t, hook, wanted2)

}
