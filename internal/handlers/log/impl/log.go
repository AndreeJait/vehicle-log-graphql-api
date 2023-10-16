package impl

import (
	"vehicle-log-graphql-api/internal/handlers/log"
	logUc "vehicle-log-graphql-api/internal/usecase/log"
)

type handler struct {
	ucLog logUc.UseCase
}

func New(ucLog logUc.UseCase) log.Handler {
	return &handler{
		ucLog: ucLog,
	}
}
