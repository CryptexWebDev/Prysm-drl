package state_native

import (
	"fmt"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/blocks"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	"github.com/Dorol-Chain/Prysm-drl/v5/runtime/version"
)

// LatestExecutionPayloadHeader of the beacon state.
func (b *BeaconState) LatestExecutionPayloadHeader() (interfaces.ExecutionData, error) {
	if b.version < version.Bellatrix {
		return nil, errNotSupported("LatestExecutionPayloadHeader", b.version)
	}

	b.lock.RLock()
	defer b.lock.RUnlock()

	switch b.version {
	case version.Bellatrix:
		return blocks.WrappedExecutionPayloadHeader(b.latestExecutionPayloadHeader.Copy())
	case version.Capella:
		return blocks.WrappedExecutionPayloadHeaderCapella(b.latestExecutionPayloadHeaderCapella.Copy())
	case version.Deneb, version.Electra:
		return blocks.WrappedExecutionPayloadHeaderDeneb(b.latestExecutionPayloadHeaderDeneb.Copy())
	default:
		return nil, fmt.Errorf("unsupported version (%s) for latest execution payload header", version.String(b.version))
	}
}