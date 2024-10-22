package beacon

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	beacondb "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/p2p"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/core"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/sync"
)

type Server struct {
	SyncChecker           sync.Checker
	HeadFetcher           blockchain.HeadFetcher
	TimeFetcher           blockchain.TimeFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	CanonicalHistory      *stategen.CanonicalHistory
	BeaconDB              beacondb.ReadOnlyDatabase
	Stater                lookup.Stater
	ChainInfoFetcher      blockchain.ChainInfoFetcher
	FinalizationFetcher   blockchain.FinalizationFetcher
	CoreService           *core.Service
	Broadcaster           p2p.Broadcaster
	BlobReceiver          blockchain.BlobReceiver
}
