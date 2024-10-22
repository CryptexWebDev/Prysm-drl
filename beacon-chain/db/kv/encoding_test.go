package kv

import (
	"context"
	"testing"

	testpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/testing"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
)

func Test_encode_handlesNilFromFunction(t *testing.T) {
	foo := func() *testpb.Puzzle {
		return nil
	}
	_, err := encode(context.Background(), foo())
	require.ErrorContains(t, "cannot encode nil message", err)
}
