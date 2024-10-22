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

func blockWithAttesterSlashing(asSSZ []byte) (interfaces.SignedBeaconBlock, error) {
	as := &ethpb.AttesterSlashing{}
	if err := as.UnmarshalSSZ(asSSZ); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockAltair()
	b.Block.Body = &ethpb.BeaconBlockBodyAltair{AttesterSlashings: []*ethpb.AttesterSlashing{as}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunAttesterSlashingTest(t *testing.T, config string) {
	common.RunAttesterSlashingTest(t, config, version.String(version.Altair), blockWithAttesterSlashing, sszToState)
}
