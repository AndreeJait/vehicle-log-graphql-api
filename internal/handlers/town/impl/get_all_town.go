package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

func (h handler) GetaAllTown(ctx context.Context) ([]*model.Town, error) {
	return h.ucTown.GetAllTown(ctx)
}
