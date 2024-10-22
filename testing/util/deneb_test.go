package util

import (
	"testing"

	fieldparams "github.com/Dorol-Chain/Prysm-drl/v5/config/fieldparams"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func TestInclusionProofs(t *testing.T) {
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, 0, fieldparams.MaxBlobsPerBlock)
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
