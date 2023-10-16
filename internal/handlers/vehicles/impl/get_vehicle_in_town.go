package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	"vehicle-log-graphql-api/internal/common/converter"
)

func (h handler) GetVehicleInTown(ctx context.Context, platNumber, vehicleType, townCode string) ([]*model.VehicleTown, error) {
	lists, err := h.vUc.GetVehicleInTow(ctx, vehicleType, platNumber, townCode)
	if err != nil {
		return nil, err
	}
	return converter.ConvertFromGrpcVehicleToModelVehicleList(lists), err
}
