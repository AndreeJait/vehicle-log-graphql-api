package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
	"vehicle-log-graphql-api/internal/common/converter"
)

func (u useCase) GetAllTown(ctx context.Context) ([]*model.Town, error) {
	townClient := u.clientGrpc.GetTownClient()

	towns, err := townClient.GetAllTown(ctx, &grpc.Empty{})
	if err != nil {
		return nil, err
	}
	converted := converter.ConvertGrpcTownToModelTownList(towns)
	return converted, nil
}
