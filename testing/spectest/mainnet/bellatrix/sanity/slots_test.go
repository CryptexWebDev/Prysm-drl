package sanity

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/bellatrix/sanity"
)

func TestMainnet_Bellatrix_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "mainnet")
}
