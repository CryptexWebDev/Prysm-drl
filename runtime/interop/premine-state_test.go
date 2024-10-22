package interop

import (
	"context"
	"math/big"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/time"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestPremineGenesis_Electra(t *testing.T) {
	one := uint64(1)

	genesis := types.NewBlockWithHeader(&types.Header{
		Time:          uint64(time.Now().Unix()),
		Extra:         make([]byte, 32),
		BaseFee:       big.NewInt(1),
		ExcessBlobGas: &one,
		BlobGasUsed:   &one,
	})
	_, err := NewPreminedGenesis(context.Background(), genesis.Time(), 10, 10, version.Electra, genesis)
	require.NoError(t, err)
}
