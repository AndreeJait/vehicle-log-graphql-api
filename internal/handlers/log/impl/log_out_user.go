package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	"vehicle-log-graphql-api/internal/common/converter"
	"vehicle-log-graphql-api/internal/common/utils"
)

func (h handler) LogOutUser(ctx context.Context, vehicleType, platNumber, townCode string) (*model.LogOutResponse, error) {
	timeOut := utils.NowJkrt()
	res, err := h.ucLog.LogOutUser(ctx, vehicleType, platNumber, townCode, timeOut.Format(utils.DefaultFormatTime))
	if err != nil {
		return nil, err
	}
	return converter.ConvertResponseLogToGqlModel(res), nil
}
