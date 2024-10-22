package validator

import (
	"testing"

	consensusblocks "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
)

func TestUnblinder_UnblindBlobSidecars_InvalidBundle(t *testing.T) {
	wBlock, err := consensusblocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockDeneb{
		Block: &ethpb.BeaconBlockDeneb{
			Body: &ethpb.BeaconBlockBodyDeneb{},
		},
		Signature: nil,
	})
	assert.NoError(t, err)
	_, err = unblindBlobsSidecars(wBlock, nil)
	assert.NoError(t, err)

	wBlock, err = consensusblocks.NewSignedBeaconBlock(&ethpb.SignedBeaconBlockDeneb{
		Block: &ethpb.BeaconBlockDeneb{
			Body: &ethpb.BeaconBlockBodyDeneb{
				BlobKzgCommitments: [][]byte{[]byte("a"), []byte("b")},
			},
		},
		Signature: nil,
	})
	assert.NoError(t, err)
	_, err = unblindBlobsSidecars(wBlock, nil)
	assert.ErrorContains(t, "no valid bundle provided", err)

}
