package graph

import (
	"vehicle-log-graphql-api/internal/handlers/log"
	"vehicle-log-graphql-api/internal/handlers/town"
	"vehicle-log-graphql-api/internal/handlers/types"
	"vehicle-log-graphql-api/internal/handlers/vehicles"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TownHandler    town.Handler
	TypeHandler    types.Handler
	LogHandler     log.Handler
	VehicleHandler vehicles.Handler
}
