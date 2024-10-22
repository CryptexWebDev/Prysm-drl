package forkchoice

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Bellatrix_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Bellatrix)
}
