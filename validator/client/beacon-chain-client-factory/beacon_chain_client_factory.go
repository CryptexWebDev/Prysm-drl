package beacon_chain_client_factory

import (
	"github.com/Dorol-Chain/Prysm-drl/v5/config/features"
	beaconApi "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/beacon-api"
	grpcApi "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/grpc-api"
	"github.com/Dorol-Chain/Prysm-drl/v5/validator/client/iface"
	nodeClientFactory "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/node-client-factory"
	validatorHelpers "github.com/Dorol-Chain/Prysm-drl/v5/validator/helpers"
)

func NewChainClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.ChainClient {
	grpcClient := grpcApi.NewGrpcChainClient(validatorConn.GetGrpcClientConn())
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewBeaconApiChainClientWithFallback(jsonRestHandler, grpcClient)
	} else {
		return grpcClient
	}
}

func NewPrysmChainClient(validatorConn validatorHelpers.NodeConnection, jsonRestHandler beaconApi.JsonRestHandler) iface.PrysmChainClient {
	if features.Get().EnableBeaconRESTApi {
		return beaconApi.NewPrysmChainClient(jsonRestHandler, nodeClientFactory.NewNodeClient(validatorConn, jsonRestHandler))
	} else {
		return grpcApi.NewGrpcPrysmChainClient(validatorConn.GetGrpcClientConn())
	}
}