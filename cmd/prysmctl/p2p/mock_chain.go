package p2p

import (
	"time"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/forkchoice"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/time/slots"
)

type mockChain struct {
	currentFork     *ethpb.Fork
	genesisValsRoot [32]byte
	genesisTime     time.Time
}

func (m *mockChain) ForkChoicer() forkchoice.ForkChoicer {
	return nil
}

func (m *mockChain) CurrentFork() *ethpb.Fork {
	return m.currentFork
}

func (m *mockChain) GenesisValidatorsRoot() [32]byte {
	return m.genesisValsRoot
}

func (m *mockChain) GenesisTime() time.Time {
	return m.genesisTime
}

func (m *mockChain) CurrentSlot() primitives.Slot {
	return slots.SinceGenesis(m.genesisTime)
}
