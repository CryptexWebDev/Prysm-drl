package attestations

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)
