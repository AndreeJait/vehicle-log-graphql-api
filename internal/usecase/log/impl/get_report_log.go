package impl

import (
	"context"
	grpcPB "vehicle-log-graphql-api/common/grpc/proto"
	"vehicle-log-graphql-api/internal/common/utils"
)

func (u useCase) GetReportLog(ctx context.Context, townCode, date string) (*grpcPB.TownLogReportList, error) {
	if date == "" {
		timeNow := utils.NowJkrt()
		date = timeNow.Format(utils.DefaultFormatDate)
	}
	res, err := u.clientGrpc.GetLogClient().ReportLogged(ctx, &grpcPB.GetLogReport{
		TownCode: townCode,
		Date:     date,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
