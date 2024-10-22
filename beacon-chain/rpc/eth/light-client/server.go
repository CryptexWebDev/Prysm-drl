package lightclient

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker     lookup.Blocker
	Stater      lookup.Stater
	HeadFetcher blockchain.HeadFetcher
	BeaconDB    db.HeadAccessDatabase
}
