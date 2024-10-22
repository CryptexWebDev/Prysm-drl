package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	common "github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/operations"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func blockWithBlsChange(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	c := &ethpb.SignedBLSToExecutionChange{}
	if err := c.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockCapella()
	b.Block.Body = &ethpb.BeaconBlockBodyCapella{BlsToExecutionChanges: []*ethpb.SignedBLSToExecutionChange{c}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunBLSToExecutionChangeTest(t *testing.T, config string) {
	common.RunBLSToExecutionChangeTest(t, config, version.String(version.Capella), blockWithBlsChange, sszToState)
}
