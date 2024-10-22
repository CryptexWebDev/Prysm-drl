package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
