package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "minimal")
}
