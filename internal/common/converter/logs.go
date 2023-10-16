package converter

import (
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpcPB "vehicle-log-graphql-api/common/grpc/proto"
	"vehicle-log-graphql-api/internal/usecase/log"
)

func ConvertResponseLogToGqlModel(res log.ResponseLog) *model.LogOutResponse {
	return &model.LogOutResponse{
		TownCode:    &res.TownCode,
		TotalTime:   &res.TotalTime,
		TimeIn:      &res.TimeIn,
		TimeOut:     &res.TimeOut,
		DateAt:      &res.DateAt,
		DateOutAt:   &res.DateOutAt,
		PlatNumber:  &res.PlatNumber,
		VehicleType: &res.VehicleType,
		TownName:    res.TownName,
	}
}

func ConvertGrpcReportToReportGraphQlList(listGrpc *grpcPB.TownLogReportList) []*model.TownLogReport {
	modelGQls := make([]*model.TownLogReport, 0)
	for _, v := range listGrpc.List {
		convert := ConvertGrpcReportToReportGraphQl(v)
		modelGQls = append(modelGQls, &convert)
	}
	return modelGQls
}

func ConvertGrpcReportToReportGraphQl(item *grpcPB.TownLogReport) model.TownLogReport {
	totalLogged := int(item.TotalLogged)
	quantity := int(item.Quantity)
	return model.TownLogReport{
		TownCode:    &item.TownCode,
		VehicleType: &item.VehicleType,
		TotalLogged: &totalLogged,
		Quantity:    &quantity,
	}
}
