package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
	"vehicle-log-graphql-api/internal/common/converter"
)

func (u useCase) GetAllTypes(ctx context.Context) ([]*model.Types, error) {
	results := make([]*model.Types, 0)
	types, err := u.clientGrpc.GetTypeClient().GetAllType(ctx, &grpc.EmptyType{})
	if err != nil {
		return results, err
	}
	return converter.ConvertGrpcTypesToModelTypesList(types), nil
}
