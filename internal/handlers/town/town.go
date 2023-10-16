package town

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

type Handler interface {
	GetaAllTown(ctx context.Context) ([]*model.Town, error)
}
