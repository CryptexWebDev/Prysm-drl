// Package testing includes useful mocks for slot tickers in unit tests.
package testing

import "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"

// MockTicker defines a useful struct for mocking the Ticker interface
// from the slotutil package.
type MockTicker struct {
	Channel chan primitives.Slot
}

// C --
func (m *MockTicker) C() <-chan primitives.Slot {
	return m.Channel
}

// Done --
func (_ *MockTicker) Done() {}
