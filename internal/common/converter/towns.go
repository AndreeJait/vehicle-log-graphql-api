package converter

import (
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
)

func ConvertGrpcTownToModelTownList(townListGrpc *grpc.TownList) []*model.Town {
	townsResult := make([]*model.Town, 0)
	for _, v := range townListGrpc.Lists {
		convert := ConvertGrpcTownToModelTown(*v)
		townsResult = append(townsResult, &convert)
	}
	return townsResult
}

func ConvertGrpcTownToModelTown(townGrpc grpc.Town) model.Town {
	return model.Town{
		ID:       &townGrpc.Id,
		TownName: &townGrpc.TownName,
		TownCode: &townGrpc.TownCode,
	}
}
