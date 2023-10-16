package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

// LogIn is the resolver for the LogIn field.
func (r *mutationResolver) LogIn(ctx context.Context, input *model.LogInRequest) (*bool, error) {
	result := true
	err := r.LogHandler.LogInUser(ctx, input.VehicleType, input.PlatNumber, input.TownCode)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// LogOut is the resolver for the LogOut field.
func (r *mutationResolver) LogOut(ctx context.Context, input *model.LogOutRequest) (*model.LogOutResponse, error) {
	result, err := r.LogHandler.LogOutUser(ctx, *input.VehicleType, input.PlatNumber, input.TownCode)
	if err != nil {
		return result, err
	}
	return result, nil
}

// ReportLogged is the resolver for the ReportLogged field.
func (r *queryResolver) ReportLogged(ctx context.Context, filter model.FilterReportLogged) ([]*model.TownLogReport, error) {
	result, err := r.LogHandler.GetReportLogged(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
