package coverage

import "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"

// AvailableBlocker can be used to check whether there is a finalized block in the db for the given slot.
// This interface is typically fulfilled by backfill.Store.
type AvailableBlocker interface {
	AvailableBlock(primitives.Slot) bool
}
