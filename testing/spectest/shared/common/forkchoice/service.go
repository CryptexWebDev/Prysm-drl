package forkchoice

import (
	"context"
	"math/big"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain/kzg"
	mock "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/blockchain/testing"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache/depositsnapshot"
	coreTime "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/time"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/filesystem"
	testDB "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/testing"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/forkchoice"
	doublylinkedtree "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/forkchoice/doubly-linked-tree"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/operations/attestations"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/startup"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state/stategen"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	payloadattribute "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/payload-attribute"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
	pb "github.com/Dorol-Chain/Prysm-drl/v5/proto/engine/v1"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func startChainService(t testing.TB,
	st state.BeaconState,
	block interfaces.ReadOnlySignedBeaconBlock,
	engineMock *engineMock,
	clockSync *startup.ClockSynchronizer,
) (*blockchain.Service, *stategen.State, forkchoice.ForkChoicer) {
	ctx := context.Background()
	db := testDB.SetupDB(t)
	require.NoError(t, db.SaveBlock(ctx, block))
	r, err := block.Block().HashTreeRoot()
	require.NoError(t, err)
	require.NoError(t, db.SaveGenesisBlockRoot(ctx, r))

	cp := &ethpb.Checkpoint{
		Epoch: coreTime.CurrentEpoch(st),
		Root:  r[:],
	}
	require.NoError(t, db.SaveState(ctx, st, r))
	require.NoError(t, db.SaveJustifiedCheckpoint(ctx, cp))
	require.NoError(t, db.SaveFinalizedCheckpoint(ctx, cp))
	attPool, err := attestations.NewService(ctx, &attestations.Config{
		Pool: attestations.NewPool(),
	})
	require.NoError(t, err)

	depositCache, err := depositsnapshot.New()
	require.NoError(t, err)

	fc := doublylinkedtree.New()
	sg := stategen.New(db, fc)
	opts := append([]blockchain.Option{},
		blockchain.WithExecutionEngineCaller(engineMock),
		blockchain.WithFinalizedStateAtStartUp(st),
		blockchain.WithDatabase(db),
		blockchain.WithAttestationService(attPool),
		blockchain.WithForkChoiceStore(fc),
		blockchain.WithStateGen(sg),
		blockchain.WithStateNotifier(&mock.MockStateNotifier{}),
		blockchain.WithAttestationPool(attestations.NewPool()),
		blockchain.WithDepositCache(depositCache),
		blockchain.WithTrackedValidatorsCache(cache.NewTrackedValidatorsCache()),
		blockchain.WithPayloadIDCache(cache.NewPayloadIDCache()),
		blockchain.WithClockSynchronizer(clockSync),
		blockchain.WithBlobStorage(filesystem.NewEphemeralBlobStorage(t)),
		blockchain.WithSyncChecker(mock.MockChecker{}),
		blockchain.WithBlobStorage(filesystem.NewEphemeralBlobStorage(t)),
	)
	service, err := blockchain.NewService(context.Background(), opts...)
	require.NoError(t, err)
	// force start kzg context here until Deneb fork epoch is decided
	require.NoError(t, kzg.Start())
	require.NoError(t, service.StartFromSavedState(st))
	return service, sg, fc
}

type engineMock struct {
	powBlocks       map[[32]byte]*ethpb.PowBlock
	latestValidHash []byte
	payloadStatus   error
}

func (m *engineMock) GetPayload(context.Context, [8]byte, primitives.Slot) (*blocks.GetPayloadResponse, error) {
	return nil, nil
}
func (m *engineMock) GetPayloadV2(context.Context, [8]byte) (*pb.ExecutionPayloadCapella, error) {
	return nil, nil
}
func (m *engineMock) ForkchoiceUpdated(context.Context, *pb.ForkchoiceState, payloadattribute.Attributer) (*pb.PayloadIDBytes, []byte, error) {
	return nil, m.latestValidHash, m.payloadStatus
}

func (m *engineMock) NewPayload(context.Context, interfaces.ExecutionData, []common.Hash, *common.Hash, *pb.ExecutionRequests) ([]byte, error) {
	return m.latestValidHash, m.payloadStatus
}

func (m *engineMock) ForkchoiceUpdatedV2(context.Context, *pb.ForkchoiceState, payloadattribute.Attributer) (*pb.PayloadIDBytes, []byte, error) {
	return nil, m.latestValidHash, m.payloadStatus
}

func (m *engineMock) LatestExecutionBlock(context.Context) (*pb.ExecutionBlock, error) {
	return nil, nil
}

func (m *engineMock) ExecutionBlockByHash(_ context.Context, hash common.Hash, _ bool) (*pb.ExecutionBlock, error) {
	b, ok := m.powBlocks[bytesutil.ToBytes32(hash.Bytes())]
	if !ok {
		return nil, nil
	}

	td := new(big.Int).SetBytes(bytesutil.ReverseByteOrder(b.TotalDifficulty))
	tdHex := hexutil.EncodeBig(td)
	return &pb.ExecutionBlock{
		Header: gethtypes.Header{
			ParentHash: common.BytesToHash(b.ParentHash),
		},
		TotalDifficulty: tdHex,
		Hash:            common.BytesToHash(b.BlockHash),
	}, nil
}

func (m *engineMock) GetTerminalBlockHash(context.Context, uint64) ([]byte, bool, error) {
	return nil, false, nil
}