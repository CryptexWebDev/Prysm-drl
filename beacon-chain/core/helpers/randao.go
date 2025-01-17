package helpers

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/state"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
	"github.com/Dorol-Chain/Prysm-drl/v5/crypto/bls"
	"github.com/Dorol-Chain/Prysm-drl/v5/crypto/hash"
	"github.com/Dorol-Chain/Prysm-drl/v5/encoding/bytesutil"
)

// Seed returns the randao seed used for shuffling of a given epoch.
//
// Spec pseudocode definition:
//
//	def get_seed(state: BeaconState, epoch: Epoch, domain_type: DomainType) -> Bytes32:
//	  """
//	  Return the seed at ``epoch``.
//	  """
//	  mix = get_randao_mix(state, Epoch(epoch + EPOCHS_PER_HISTORICAL_VECTOR - MIN_SEED_LOOKAHEAD - 1))  # Avoid underflow
//	  return hash(domain_type + uint_to_bytes(epoch) + mix)
func Seed(state state.ReadOnlyBeaconState, epoch primitives.Epoch, domain [bls.DomainByteLength]byte) ([32]byte, error) {
	// See https://github.com/ethereum/consensus-specs/pull/1296 for
	// rationale on why offset has to look down by 1.
	lookAheadEpoch := epoch + params.BeaconConfig().EpochsPerHistoricalVector -
		params.BeaconConfig().MinSeedLookahead - 1

	randaoMix, err := RandaoMix(state, lookAheadEpoch)
	if err != nil {
		return [32]byte{}, err
	}
	seed := append(domain[:], bytesutil.Bytes8(uint64(epoch))...)
	seed = append(seed, randaoMix...)

	seed32 := hash.Hash(seed)

	return seed32, nil
}

// RandaoMix returns the randao mix (xor'ed seed)
// of a given slot. It is used to shuffle validators.
//
// Spec pseudocode definition:
//
//	def get_randao_mix(state: BeaconState, epoch: Epoch) -> Bytes32:
//	 """
//	 Return the randao mix at a recent ``epoch``.
//	 """
//	 return state.randao_mixes[epoch % EPOCHS_PER_HISTORICAL_VECTOR]
func RandaoMix(state state.ReadOnlyBeaconState, epoch primitives.Epoch) ([]byte, error) {
	return state.RandaoMixAtIndex(uint64(epoch % params.BeaconConfig().EpochsPerHistoricalVector))
}
