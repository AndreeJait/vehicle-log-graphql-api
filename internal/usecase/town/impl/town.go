package impl

import (
	"vehicle-log-graphql-api/config"
	"vehicle-log-graphql-api/internal/common/grpc"
	"vehicle-log-graphql-api/internal/repository/redis"
	"vehicle-log-graphql-api/internal/usecase/town"
)

type useCase struct {
	clientGrpc grpc.Client
	cfg        *config.Config
	redis      redis.Redis
}

func New(clientRpc grpc.Client, cfg *config.Config, redis redis.Redis) town.UseCase {
	return &useCase{
		redis:      redis,
		clientGrpc: clientRpc,
		cfg:        cfg,
	}
}
