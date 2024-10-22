package sync

import (
	"context"
	"fmt"

	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed"
	opfeed "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/core/feed/operation"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"google.golang.org/protobuf/proto"
)

func (s *Service) blobSubscriber(ctx context.Context, msg proto.Message) error {
	b, ok := msg.(blocks.VerifiedROBlob)
	if !ok {
		return fmt.Errorf("message was not type blocks.ROBlob, type=%T", msg)
	}

	s.setSeenBlobIndex(b.Slot(), b.ProposerIndex(), b.Index)

	if err := s.cfg.chain.ReceiveBlob(ctx, b); err != nil {
		return err
	}

	s.cfg.operationNotifier.OperationFeed().Send(&feed.Event{
		Type: opfeed.BlobSidecarReceived,
		Data: &opfeed.BlobSidecarReceivedData{
			Blob: &b,
		},
	})

	return nil
}
