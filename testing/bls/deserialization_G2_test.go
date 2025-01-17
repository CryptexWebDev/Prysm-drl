package bls

import (
	"encoding/hex"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/crypto/bls"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/bls/utils"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/ghodss/yaml"
)

func TestDeserializationG2(t *testing.T) {
	t.Run("blst", testDeserializationG2)
}

func testDeserializationG2(t *testing.T) {
	fNames, fContent := utils.RetrieveFiles("deserialization_G2", t)

	for i, file := range fNames {
		content := fContent[i]
		t.Run(file, func(t *testing.T) {
			test := &DeserializationG2Test{}
			require.NoError(t, yaml.Unmarshal(content, test))
			rawKey, err := hex.DecodeString(test.Input.Signature)
			require.NoError(t, err)

			_, err = bls.SignatureFromBytes(rawKey)
			require.Equal(t, test.Output, err == nil)
			t.Log("Success")
		})
	}
}
