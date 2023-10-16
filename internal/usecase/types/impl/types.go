package impl

import (
	"vehicle-log-graphql-api/config"
	"vehicle-log-graphql-api/internal/common/grpc"
	"vehicle-log-graphql-api/internal/repository/redis"
	"vehicle-log-graphql-api/internal/usecase/types"
)

type useCase struct {
	clientGrpc grpc.Client
	cfg        *config.Config
	redis      redis.Redis
}

func New(clientGrpc grpc.Client, cfg *config.Config, redis redis.Redis) types.UseCase {
	return &useCase{
		clientGrpc: clientGrpc,
		cfg:        cfg,
		redis:      redis,
	}
}
