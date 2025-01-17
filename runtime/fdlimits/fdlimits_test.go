package fdlimits_test

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/fdlimits"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
	gethLimit "github.com/ethereum/go-ethereum/common/fdlimit"
)

func TestSetMaxFdLimits(t *testing.T) {
	assert.NoError(t, fdlimits.SetMaxFdLimits())

	curr, err := gethLimit.Current()
	assert.NoError(t, err)

	max, err := gethLimit.Maximum()
	assert.NoError(t, err)

	assert.Equal(t, max, curr, "current and maximum file descriptor limits do not match up.")

}
