package impl

import (
	"context"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
)

func (u useCase) GetVehicleInTow(ctx context.Context, vehicleType, platNumber, townCode string) ([]*grpc.Vehicle, error) {
	res, err := u.clientGrpc.GetVehicleClient().GetVehicleInTown(ctx, &grpc.VehicleInTownRequest{
		TownCode: townCode,
		Filter: &grpc.Filter{
			VehicleType: vehicleType,
			PlatNumber:  platNumber,
		},
	})
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
