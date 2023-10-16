package impl

import grpcPB "vehicle-log-graphql-api/common/grpc/proto"

func (c client) GetVehicleClient() grpcPB.VehiclesClient {
	cl := grpcPB.NewVehiclesClient(c.cnn)
	return cl
}
