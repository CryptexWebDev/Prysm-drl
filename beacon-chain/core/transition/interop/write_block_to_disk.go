package interop

import (
	"fmt"
	"os"
	"path"

	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/interfaces"
	"github.com/Dorol-Chain/Prysm-drl/v5/io/file"
)

// WriteBlockToDisk as a block ssz. Writes to temp directory. Debug!
func WriteBlockToDisk(block interfaces.ReadOnlySignedBeaconBlock, failed bool) {
	if !features.Get().WriteSSZStateTransitions {
		return
	}

	filename := fmt.Sprintf("beacon_block_%d.ssz", block.Block().Slot())
	if failed {
		filename = "failed_" + filename
	}
	fp := path.Join(os.TempDir(), filename)
	log.Warnf("Writing block to disk at %s", fp)
	enc, err := block.MarshalSSZ()
	if err != nil {
		log.WithError(err).Error("Failed to ssz encode block")
		return
	}
	if err := file.WriteFile(fp, enc); err != nil {
		log.WithError(err).Error("Failed to write to disk")
	}
}
