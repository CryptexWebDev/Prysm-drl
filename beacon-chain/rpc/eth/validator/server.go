package validator

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/builder"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/operation"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/attestations"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/synccommittee"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/p2p"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/core"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/eth/rewards"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/lookup"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/sync"
	eth "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Validator service,
// providing RPC endpoints intended for validator clients.
type Server struct {
	HeadFetcher            blockchain.HeadFetcher
	TimeFetcher            blockchain.TimeFetcher
	SyncChecker            sync.Checker
	AttestationsPool       attestations.Pool
	PeerManager            p2p.PeerManager
	Broadcaster            p2p.Broadcaster
	Stater                 lookup.Stater
	OptimisticModeFetcher  blockchain.OptimisticModeFetcher
	SyncCommitteePool      synccommittee.Pool
	V1Alpha1Server         eth.BeaconNodeValidatorServer
	ChainInfoFetcher       blockchain.ChainInfoFetcher
	BeaconDB               db.HeadAccessDatabase
	BlockBuilder           builder.BlockBuilder
	OperationNotifier      operation.Notifier
	CoreService            *core.Service
	BlockRewardFetcher     rewards.BlockRewardsFetcher
	TrackedValidatorsCache *cache.TrackedValidatorsCache
	PayloadIDCache         *cache.PayloadIDCache
}
