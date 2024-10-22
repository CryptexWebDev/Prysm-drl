package iface

import (
	"context"

	"github.com/Dorol-Chain/Prysm-drl/v5/api/client/beacon"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
)

type NodeClient interface {
	SyncStatus(ctx context.Context, in *empty.Empty) (*ethpb.SyncStatus, error)
	Genesis(ctx context.Context, in *empty.Empty) (*ethpb.Genesis, error)
	Version(ctx context.Context, in *empty.Empty) (*ethpb.Version, error)
	Peers(ctx context.Context, in *empty.Empty) (*ethpb.Peers, error)
	HealthTracker() *beacon.NodeHealthTracker
}
