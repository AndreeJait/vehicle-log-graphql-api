package impl

import (
	"context"
	"vehicle-log-graphql-api/common/graphql/graph/model"
)

func (h handler) GetAllTypes(ctx context.Context) ([]*model.Types, error) {
	return h.uc.GetAllTypes(ctx)
}
