package stateutil

import (
	"testing"

	"github.com/Dorol-Chain/Prysm-drl/v5/testing/assert"
)

func TestArraysTreeRoot_OnlyPowerOf2(t *testing.T) {
	_, err := ArraysRoot([][]byte{}, 1)
	assert.NoError(t, err)
	_, err = ArraysRoot([][]byte{}, 4)
	assert.NoError(t, err)
	_, err = ArraysRoot([][]byte{}, 8)
	assert.NoError(t, err)
	_, err = ArraysRoot([][]byte{}, 10)
	assert.ErrorContains(t, "hash layer is a non power of 2", err)
}

func TestArraysTreeRoot_ZeroLength(t *testing.T) {
	_, err := ArraysRoot([][]byte{}, 0)
	assert.ErrorContains(t, "zero leaves provided", err)
}
