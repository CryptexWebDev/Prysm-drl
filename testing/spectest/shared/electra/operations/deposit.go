package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/electra"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	common "github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/operations"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/util"
)

func blockWithDeposit(ssz []byte) (interfaces.SignedBeaconBlock, error) {
	d := &ethpb.Deposit{}
	if err := d.UnmarshalSSZ(ssz); err != nil {
		return nil, err
	}
	b := util.NewBeaconBlockElectra()
	b.Block.Body = &ethpb.BeaconBlockBodyElectra{Deposits: []*ethpb.Deposit{d}}
	return blocks.NewSignedBeaconBlock(b)
}

func RunDepositTest(t *testing.T, config string) {
	common.RunDepositTest(t, config, version.String(version.Electra), blockWithDeposit, electra.ProcessDeposits, sszToState)
}