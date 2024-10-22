package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/operations"
)

func TestMainnet_Electra_Operations_Withdrawals(t *testing.T) {
	operations.RunWithdrawalsTest(t, "mainnet")
}
