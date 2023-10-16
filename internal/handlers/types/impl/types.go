package impl

import (
	"vehicle-log-graphql-api/internal/handlers/types"
	typesUseCase "vehicle-log-graphql-api/internal/usecase/types"
)

type handler struct {
	uc typesUseCase.UseCase
}

func New(uc typesUseCase.UseCase) types.Handler {
	return &handler{
		uc: uc,
	}
}
