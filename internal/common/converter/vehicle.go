package converter

import (
	"vehicle-log-graphql-api/common/graphql/graph/model"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
)

func ConvertFromGrpcVehicleToModelVehicleList(vehicleGrpcs []*grpc.Vehicle) []*model.VehicleTown {
	results := make([]*model.VehicleTown, 0)
	for _, v := range vehicleGrpcs {
		convert := ConvertFromGrpcVehicleToModelVehicle(*v)
		results = append(results, &convert)
	}
	return results
}

func ConvertFromGrpcVehicleToModelVehicle(v grpc.Vehicle) model.VehicleTown {
	return model.VehicleTown{
		ID:           &v.Id,
		PlatNumber:   &v.PlatNumber,
		VehicleType:  &v.VehicleType,
		TimeCheckIn:  &v.TimeCheckIn,
		TimeCheckOut: &v.TimeCheckOut,
		DateCheckIn:  &v.DateCheckIn,
		DateCheckOut: &v.DateCheckOut,
	}
}
