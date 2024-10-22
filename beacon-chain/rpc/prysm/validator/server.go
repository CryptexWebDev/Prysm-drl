package validator

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/core"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
)

type Server struct {
	BeaconDB            db.ReadOnlyDatabase
	Stater              lookup.Stater
	CanonicalFetcher    blockchain.CanonicalFetcher
	FinalizationFetcher blockchain.FinalizationFetcher
	ChainInfoFetcher    blockchain.ChainInfoFetcher
	CoreService         *core.Service
}
