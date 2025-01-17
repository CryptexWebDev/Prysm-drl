// Package stateutils contains useful tools for faster computation
// of state transitions using maps to represent validators instead
// of slices.
package stateutils

import (
	fieldparams "github.com/Dorol-Chain/Prysm-drl/v5/config/fieldparams"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
)

// ValidatorIndexMap builds a lookup map for quickly determining the index of
// a validator by their public key.
func ValidatorIndexMap(validators []*ethpb.Validator) map[[fieldparams.BLSPubkeyLength]byte]primitives.ValidatorIndex {
	m := make(map[[fieldparams.BLSPubkeyLength]byte]primitives.ValidatorIndex, len(validators))
	if validators == nil {
		return m
	}
	for idx, record := range validators {
		if record == nil {
			continue
		}
		key := bytesutil.ToBytes48(record.PublicKey)
		m[key] = primitives.ValidatorIndex(idx)
	}
	return m
}
