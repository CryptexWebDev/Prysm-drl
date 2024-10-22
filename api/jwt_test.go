package api

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func TestGenerateRandomHexString(t *testing.T) {
	token, err := GenerateRandomHexString()
	require.NoError(t, err)
	require.NoError(t, ValidateAuthToken(token))
}
