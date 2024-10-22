package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/electra/operations"
)

func TestMinimal_Electra_Operations_WithdrawalRequest(t *testing.T) {
	operations.RunWithdrawalRequestTest(t, "minimal")
}
