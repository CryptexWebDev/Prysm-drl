package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	common "github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/operations"
)

func RunBlockHeaderTest(t *testing.T, config string) {
	common.RunBlockHeaderTest(t, config, version.String(version.Electra), sszToBlock, sszToState)
}
