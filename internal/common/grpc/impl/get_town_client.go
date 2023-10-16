package impl

import grpcPB "vehicle-log-graphql-api/common/grpc/proto"

func (c client) GetTownClient() grpcPB.TownsClient {
	cl := grpcPB.NewTownsClient(c.cnn)
	return cl
}
