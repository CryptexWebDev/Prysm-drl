package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_Deposit(t *testing.T) {
	operations.RunDepositTest(t, "minimal")
}
