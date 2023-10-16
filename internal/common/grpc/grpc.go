package grpc

import (
	grpcPB "vehicle-log-graphql-api/common/grpc/proto"
)

type Client interface {
	GetTownClient() grpcPB.TownsClient
	GetLogClient() grpcPB.LogsClient
	GetVehicleClient() grpcPB.VehiclesClient
	GetTypeClient() grpcPB.TypesClient
}
