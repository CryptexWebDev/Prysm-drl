package kzg

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
