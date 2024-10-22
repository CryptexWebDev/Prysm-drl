package validator_client_factory

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	beaconApi "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/beacon-api"
	grpcApi "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/grpc-api"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/client/iface"
	validatorHelpers "github.com/Dorol-Chain/Prysm-drl/v5/validator/helpers"
)

func NewValidatorClient(
	validatorConn validatorHelpers.NodeConnection,
	jsonRestHandler beaconApi.JsonRestHandler,
	opt ...beaconApi.ValidatorClientOpt,
) iface.ValidatorClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiValidatorClient(jsonRestHandler, opt...)
	} else {
		return grpcApi.NewGrpcValidatorClient(validatorConn.GetGrpcClientConn())
	}
}
