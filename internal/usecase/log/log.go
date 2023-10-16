package log

import (
	"context"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
)

type UseCase interface {
	LogInUser(ctx context.Context, vehicleType, platNumber, townCode string) error
	LogOutUser(ctx context.Context, vehicleType, platNumber, townCode string, timeOut string) (ResponseLog, error)
	GetReportLog(ctx context.Context, townCode, date string) (*grpc.TownLogReportList, error)
}
