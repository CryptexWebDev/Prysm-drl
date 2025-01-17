//go:build !fuzz

package helpers

import "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/cache"

func CommitteeCache() *cache.CommitteeCache {
	return committeeCache
}

func SyncCommitteeCache() *cache.SyncCommitteeCache {
	return syncCommitteeCache
}

func ProposerIndicesCache() *cache.ProposerIndicesCache {
	return proposerIndicesCache
}
