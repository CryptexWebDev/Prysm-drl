package peers_test

import (
	"io"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/cmd/beacon-chain/flags"
	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)

	resetCfg := features.InitWithReset(&features.Flags{
		EnablePeerScorer: true,
	})
	defer resetCfg()

	resetFlags := flags.Get()
	flags.Init(&flags.GlobalFlags{
		BlockBatchLimit:            64,
		BlockBatchLimitBurstFactor: 10,
		BlobBatchLimit:             8,
		BlobBatchLimitBurstFactor:  2,
	})
	defer func() {
		flags.Init(resetFlags)
	}()
	m.Run()
}