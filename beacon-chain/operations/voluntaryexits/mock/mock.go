package mock

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	eth "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
)

// PoolMock is a fake implementation of PoolManager.
type PoolMock struct {
	Exits []*eth.SignedVoluntaryExit
}

// PendingExits --
func (m *PoolMock) PendingExits() ([]*eth.SignedVoluntaryExit, error) {
	return m.Exits, nil
}

// ExitsForInclusion --
func (m *PoolMock) ExitsForInclusion(_ state.ReadOnlyBeaconState, _ primitives.Slot) ([]*eth.SignedVoluntaryExit, error) {
	return m.Exits, nil
}

// InsertVoluntaryExit --
func (m *PoolMock) InsertVoluntaryExit(exit *eth.SignedVoluntaryExit) {
	m.Exits = append(m.Exits, exit)
}

// MarkIncluded --
func (*PoolMock) MarkIncluded(_ *eth.SignedVoluntaryExit) {
	panic("implement me")
}
