package impl

import (
	"google.golang.org/grpc"
	grpcCommon "vehicle-log-graphql-api/internal/common/grpc"
)

type client struct {
	cnn *grpc.ClientConn
}

func New(cnn *grpc.ClientConn) grpcCommon.Client {
	return &client{
		cnn: cnn,
	}
}
