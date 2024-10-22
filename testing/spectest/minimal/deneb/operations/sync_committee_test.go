package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/deneb/operations"
)

func TestMinimal_Deneb_Operations_SyncCommittee(t *testing.T) {
	operations.RunSyncCommitteeTest(t, "minimal")
}
