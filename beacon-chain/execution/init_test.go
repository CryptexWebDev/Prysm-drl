package execution

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/config/params"
)

func init() {
	// Override network name so that hardcoded genesis files are not loaded.
	if err := params.SetActive(params.MainnetTestConfig()); err != nil {
		panic(err)
	}
}
