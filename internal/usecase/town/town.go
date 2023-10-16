package town

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

type UseCase interface {
	GetAllTown(ctx context.Context) ([]*model.Town, error)
}
