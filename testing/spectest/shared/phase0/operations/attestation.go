package operations

import (
	"testing"

	b "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	common "github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/operations"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func blockWithAttestation(attestationSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	att := &ethpb.Attestation{}
	if err := att.UnmarshalSSZ(attestationSSZ); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlock()
	b.Block.Body = &ethpb.BeaconBlockBody{Attestations: []*ethpb.Attestation{att}}
	return blocks.NewSignedBeaconBlock(b)
}

// RunAttestationTest executes "operations/attestation" tests.
func RunAttestationTest(t *testing.T, config string) {
	common.RunAttestationTest(t, config, version.String(version.Phase0), blockWithAttestation, b.ProcessAttestationsNoVerifySignature, sszToState)
}
