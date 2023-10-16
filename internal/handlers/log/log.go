package log

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

type Handler interface {
	LogInUser(ctx context.Context, vehicleType, platNumber, townCode string) error
	LogOutUser(ctx context.Context, vehicleType, platNumber, townCode string) (*model.LogOutResponse, error)
	GetReportLogged(ctx context.Context, filter model.FilterReportLogged) ([]*model.TownLogReport, error)
}
