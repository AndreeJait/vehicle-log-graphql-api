package vehicles

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

type Handler interface {
	GetVehicleInTown(ctx context.Context, platNumber, vehicleType, townCode string) ([]*model.VehicleTown, error)
}
