package keymanager_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
	"github.com/Dorol-Chain/Prysm-drl/v5/testing/require"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager/derived"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager/local"
	remoteweb3signer "github.com/Dorol-Chain/Prysm-drl/v5/validator/keymanager/remote-web3signer"
)

var (
	_ = keymanager.IKeymanager(&local.Keymanager{})
	_ = keymanager.IKeymanager(&derived.Keymanager{})

	// More granular assertions.
	_ = keymanager.KeysFetcher(&local.Keymanager{})
	_ = keymanager.KeysFetcher(&derived.Keymanager{})
	_ = keymanager.Importer(&local.Keymanager{})
	_ = keymanager.Importer(&derived.Keymanager{})
	_ = keymanager.Deleter(&local.Keymanager{})
	_ = keymanager.Deleter(&derived.Keymanager{})

	_ = keymanager.PublicKeyAdder(&remoteweb3signer.Keymanager{})
	_ = keymanager.PublicKeyDeleter(&remoteweb3signer.Keymanager{})
)

func TestKeystoreContainsPath(t *testing.T) {
	keystore := keymanager.Keystore{}
	encoded, err := json.Marshal(keystore)

	require.NoError(t, err, "Unexpected error marshalling keystore")
	assert.Equal(t, true, strings.Contains(string(encoded), "path"))
}
