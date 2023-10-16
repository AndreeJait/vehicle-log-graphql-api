package impl

import (
	"vehicle-log-graphql-api/internal/handlers/town"
	townUC "vehicle-log-graphql-api/internal/usecase/town"
)

type handler struct {
	ucTown townUC.UseCase
}

func New(uc townUC.UseCase) town.Handler {
	return &handler{
		ucTown: uc,
	}
}
