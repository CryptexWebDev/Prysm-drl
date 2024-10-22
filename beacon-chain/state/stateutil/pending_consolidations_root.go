package stateutil

import (
	fieldparams "github.com/Dorol-Chain/Prysm-drl/v5/config/fieldparams"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/ssz"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
)

func PendingConsolidationsRoot(slice []*ethpb.PendingConsolidation) ([32]byte, error) {
	return ssz.SliceRoot(slice, fieldparams.PendingConsolidationsLimit)
}
