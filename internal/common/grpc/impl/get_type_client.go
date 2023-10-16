package impl

import grpcPB "vehicle-log-graphql-api/common/grpc/proto"

func (c client) GetTypeClient() grpcPB.TypesClient {
	cl := grpcPB.NewTypesClient(c.cnn)
	return cl
}
