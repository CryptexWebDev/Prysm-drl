package iface

import (
	"context"

	"github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/validator"
	"github.com/pkg/errors"
)

var ErrNotSupported = errors.New("endpoint not supported")

type ValidatorCount struct {
	Status string
	Count  uint64
}

// PrysmChainClient defines an interface required to implement all the prysm specific custom endpoints.
type PrysmChainClient interface {
	ValidatorCount(context.Context, string, []validator.Status) ([]ValidatorCount, error)
}
