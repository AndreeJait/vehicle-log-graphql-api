package impl

import grpcPB "vehicle-log-graphql-api/common/grpc/proto"

func (c client) GetLogClient() grpcPB.LogsClient {
	cl := grpcPB.NewLogsClient(c.cnn)
	return cl
}
