package helpers

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	fieldparams "github.com/Dorol-Chain/Prysm-drl/v5/config/fieldparams"
	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/primitives"
)

// Uint64FromString converts a string into a uint64 representation.
func Uint64FromString(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// EpochFromString converts a string into Epoch.
func EpochFromString(str string) (primitives.Epoch, error) {
	e, err := Uint64FromString(str)
	if err != nil {
		return primitives.Epoch(e), err
	}
	return primitives.Epoch(e), nil
}

// SlotFromString converts a string into Slot.
func SlotFromString(str string) (primitives.Slot, error) {
	s, err := Uint64FromString(str)
	if err != nil {
		return primitives.Slot(s), err
	}
	return primitives.Slot(s), nil
}

// PubKeyFromHex takes in a hex string, verifies its length as 48 bytes, and converts that representation.
func PubKeyFromHex(str string) ([fieldparams.BLSPubkeyLength]byte, error) {
	pubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(str, "0x"))
	if err != nil {
		return [fieldparams.BLSPubkeyLength]byte{}, err
	}
	if len(pubKeyBytes) != 48 {
		return [fieldparams.BLSPubkeyLength]byte{}, fmt.Errorf("public key is not correct, 48-byte length: %s", str)
	}
	var pk [fieldparams.BLSPubkeyLength]byte
	copy(pk[:], pubKeyBytes[:fieldparams.BLSPubkeyLength])
	return pk, nil
}

// RootFromHex takes in a hex string, verifies its length as 32 bytes, and converts that representation.
func RootFromHex(str string) ([32]byte, error) {
	rootHexBytes, err := hex.DecodeString(strings.TrimPrefix(str, "0x"))
	if err != nil {
		return [32]byte{}, err
	}
	if len(rootHexBytes) != 32 {
		return [32]byte{}, fmt.Errorf("wrong root length, 32-byte length: %s", str)
	}
	var root [32]byte
	copy(root[:], rootHexBytes[:32])
	return root, nil
}

func RootToHexString(root []byte) (string, error) {
	// Nil signing roots are allowed in EIP-3076.
	if len(root) == 0 {
		return "", nil
	}
	if len(root) != 32 {
		return "", fmt.Errorf("wanted length 32, received %d", len(root))
	}
	return fmt.Sprintf("%#x", root), nil
}

func PubKeyToHexString(pubKey []byte) (string, error) {
	if len(pubKey) != 48 {
		return "", fmt.Errorf("wanted length 48, received %d", len(pubKey))
	}
	return fmt.Sprintf("%#x", pubKey), nil
}
