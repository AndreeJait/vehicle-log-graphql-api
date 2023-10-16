package converter

import (
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
)

func ConvertGrpcTypesToModelTypesList(typesListGrpc *grpc.TypeList) []*model.Types {
	typesResult := make([]*model.Types, 0)
	for _, v := range typesListGrpc.List {
		convert := ConvertGrpcTypesToModelTypes(*v)
		typesResult = append(typesResult, &convert)
	}
	return typesResult
}

func ConvertGrpcTypesToModelTypes(typesGrpc grpc.Type) model.Types {
	return model.Types{
		ID:       &typesGrpc.Id,
		TypeName: &typesGrpc.TypeName,
	}
}
