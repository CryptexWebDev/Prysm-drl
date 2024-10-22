// Package beacon defines a gRPC beacon service implementation, providing
// useful endpoints for checking fetching chain-specific data such as
// blocks, committees, validators, assignments, and more.
package beacon

import (
	"context"
	"time"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache"
	blockfeed "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/block"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/operation"
	statefeed "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/execution"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/attestations"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/slashings"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/p2p"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/rpc/core"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/sync"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
)

// Server defines a server implementation of the gRPC Beacon Chain service,
// providing RPC endpoints to access data relevant to the Ethereum beacon chain.
type Server struct {
	BeaconDB                    db.ReadOnlyDatabase
	Ctx                         context.Context
	ChainStartFetcher           execution.ChainStartFetcher
	HeadFetcher                 blockchain.HeadFetcher
	CanonicalFetcher            blockchain.CanonicalFetcher
	FinalizationFetcher         blockchain.FinalizationFetcher
	DepositFetcher              cache.DepositFetcher
	BlockFetcher                execution.POWBlockFetcher
	GenesisTimeFetcher          blockchain.TimeFetcher
	StateNotifier               statefeed.Notifier
	BlockNotifier               blockfeed.Notifier
	AttestationNotifier         operation.Notifier
	Broadcaster                 p2p.Broadcaster
	AttestationsPool            attestations.Pool
	SlashingsPool               slashings.PoolManager
	ChainStartChan              chan time.Time
	ReceivedAttestationsBuffer  chan *ethpb.Attestation
	CollectedAttestationsBuffer chan []*ethpb.Attestation
	StateGen                    stategen.StateManager
	SyncChecker                 sync.Checker
	ReplayerBuilder             stategen.ReplayerBuilder
	OptimisticModeFetcher       blockchain.OptimisticModeFetcher
	CoreService                 *core.Service
}
