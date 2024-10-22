package cache

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	lru "github.com/hashicorp/golang-lru"
)

func BalanceCacheKey(st state.ReadOnlyBeaconState) (string, error) {
	return balanceCacheKey(st)
}

func MaxCheckpointStateSize() int {
	return maxCheckpointStateSize
}

func (c *CheckpointStateCache) Cache() *lru.Cache {
	return c.cache
}
