package beacon_api

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/Dorol-Chain/Prysm-drl/v5/api/server/structs"
	ethpb "github.com/Dorol-Chain/Prysm-drl/v5/proto/prysm/v1alpha1"
	"github.com/pkg/errors"
)

func (c *beaconApiValidatorClient) submitValidatorRegistrations(ctx context.Context, registrations []*ethpb.SignedValidatorRegistrationV1) error {
	const endpoint = "/eth/v1/validator/register_validator"

	jsonRegistration := make([]*structs.SignedValidatorRegistration, len(registrations))

	for index, registration := range registrations {
		jsonRegistration[index] = structs.SignedValidatorRegistrationFromConsensus(registration)
	}

	marshalledJsonRegistration, err := json.Marshal(jsonRegistration)
	if err != nil {
		return errors.Wrap(err, "failed to marshal registration")
	}

	return c.jsonRestHandler.Post(ctx, endpoint, nil, bytes.NewBuffer(marshalledJsonRegistration), nil)
}
