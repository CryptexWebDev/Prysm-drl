package block

import "github.com/Dorol-Chain/Prysm-drl/v5/async/event"

// Notifier interface defines the methods of the service that provides block updates to consumers.
type Notifier interface {
	BlockFeed() *event.Feed
}
