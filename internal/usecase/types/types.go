package types

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

type UseCase interface {
	GetAllTypes(ctx context.Context) ([]*model.Types, error)
}
