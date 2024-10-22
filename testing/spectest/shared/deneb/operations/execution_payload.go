package operations

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	common "github.com/Dorol-Chain/Prysm-drl/v5/testing/spectest/shared/common/operations"
)

func RunExecutionPayloadTest(t *testing.T, config string) {
	common.RunExecutionPayloadTest(t, config, version.String(version.Deneb), sszToBlockBody, sszToState)
}
