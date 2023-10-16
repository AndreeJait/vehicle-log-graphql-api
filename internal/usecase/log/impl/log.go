package impl

import (
	"vehicle-log-graphql-api/config"
	"vehicle-log-graphql-api/internal/common/grpc"
	"vehicle-log-graphql-api/internal/common/publisher"
	"vehicle-log-graphql-api/internal/repository/redis"
	"vehicle-log-graphql-api/internal/usecase/log"
)

type useCase struct {
	clientGrpc grpc.Client
	cfg        *config.Config
	redis      redis.Redis
	pubs       publisher.Publisher
}

func New(clientRpc grpc.Client, cfg *config.Config, redis redis.Redis, pubs publisher.Publisher) log.UseCase {
	return &useCase{
		clientGrpc: clientRpc,
		cfg:        cfg,
		redis:      redis,
		pubs:       pubs,
	}
}
