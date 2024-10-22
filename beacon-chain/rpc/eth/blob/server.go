package blob

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker lookup.Blocker
}
