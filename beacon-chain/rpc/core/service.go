package core

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache"
	opfeed "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/operation"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/synccommittee"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/p2p"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/sync"
)

type Service struct {
	BeaconDB              db.ReadOnlyDatabase
	ChainInfoFetcher      blockchain.ChainInfoFetcher
	HeadFetcher           blockchain.HeadFetcher
	FinalizedFetcher      blockchain.FinalizationFetcher
	GenesisTimeFetcher    blockchain.TimeFetcher
	SyncChecker           sync.Checker
	Broadcaster           p2p.Broadcaster
	SyncCommitteePool     synccommittee.Pool
	OperationNotifier     opfeed.Notifier
	AttestationCache      *cache.AttestationCache
	StateGen              stategen.StateManager
	P2P                   p2p.Broadcaster
	ReplayerBuilder       stategen.ReplayerBuilder
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
}
