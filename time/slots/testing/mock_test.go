package testing

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)
