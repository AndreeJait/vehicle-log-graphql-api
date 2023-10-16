package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	"vehicle-log-graphql-api/internal/common/converter"
	"vehicle-log-graphql-api/internal/common/utils"
)

func (h handler) GetReportLogged(ctx context.Context, filter model.FilterReportLogged) ([]*model.TownLogReport, error) {
	res, err := h.ucLog.GetReportLog(ctx, filter.TownCode, utils.PtrStr(filter.Date))
	if err != nil {
		return nil, err
	}
	return converter.ConvertGrpcReportToReportGraphQlList(res), err
}
