package impl

import (
	"context"
	log2 "log"
	grpc "vehicle-log-graphql-api/common/grpc/proto"
	"vehicle-log-graphql-api/internal/common/publisher"
	"vehicle-log-graphql-api/internal/usecase/log"
)

func (u useCase) LogOutUser(ctx context.Context, vehicleType, platNumber, townCode string, timeOut string) (log.ResponseLog, error) {
	res, err := u.clientGrpc.GetLogClient().LogOutVehicle(ctx, &grpc.LogOutRequest{
		TownCode:   townCode,
		PlatNumber: platNumber,
		TimeOut:    timeOut,
	})

	if err != nil {
		return log.ResponseLog{}, err
	}

	go u.doBackgroundCheckoutStatus(vehicleType, platNumber, townCode)

	return log.ResponseLog{
		TownName:    res.TownName,
		TotalTime:   int(res.TotalTime),
		TimeIn:      res.TimeIn,
		TimeOut:     res.TimeOut,
		DateAt:      res.DateAt,
		PlatNumber:  res.PlatNumber,
		VehicleType: res.VehicleType,
		DateOutAt:   res.DateOutAt,
		TownCode:    res.TownCode,
	}, err
}

func (u useCase) doBackgroundCheckoutStatus(vehicleType, platNumber, townCode string) {
	err := u.pubs.LogOut(context.Background(), publisher.MessageLogOut{
		TownCode:    townCode,
		PlatNumber:  platNumber,
		VehicleType: vehicleType,
	})
	if err != nil {
		log2.Println(err)
	}
}
