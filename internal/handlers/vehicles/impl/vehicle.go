package impl

import (
	"vehicle-log-graphql-api/internal/handlers/vehicles"
	vehicleUC "vehicle-log-graphql-api/internal/usecase/vehicle"
)

type handler struct {
	vUc vehicleUC.UseCase
}

func New(vUc vehicleUC.UseCase) vehicles.Handler {
	return &handler{
		vUc: vUc,
	}
}
