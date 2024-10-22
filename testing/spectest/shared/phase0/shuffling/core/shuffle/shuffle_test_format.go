package shuffle

import "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"

// TestCase --
type TestCase struct {
	Seed    string                      `yaml:"seed"`
	Count   uint64                      `yaml:"count"`
	Mapping []primitives.ValidatorIndex `yaml:"mapping"`
}
