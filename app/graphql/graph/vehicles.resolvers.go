package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
	"vehicle-log-graphql-api/internal/common/utils"
)

// GetVehicleInTown is the resolver for the GetVehicleInTown field.
func (r *queryResolver) GetVehicleInTown(ctx context.Context, filter *model.FilterVehicleTown) ([]*model.VehicleTown, error) {
	res, err := r.VehicleHandler.GetVehicleInTown(ctx, utils.PtrStr(filter.PlatNumber), utils.PtrStr(filter.VehicleType), filter.TownCode)
	if err != nil {
		return nil, err
	}
	return res, err
}