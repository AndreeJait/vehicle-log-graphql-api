package vehicle

import (
	"context"
	grpcPB "vehicle-log-graphql-api/common/grpc/proto"
)

type UseCase interface {
	GetVehicleInTow(ctx context.Context, vehicleType, platNumber, townCode string) ([]*grpcPB.Vehicle, error)
}
