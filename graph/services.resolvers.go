package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/d0kur0/cui-server/auth"
	"github.com/d0kur0/cui-server/graph/model"
)

func (r *queryResolver) Services(ctx context.Context, count *int) ([]*model.Service, error) {
	user := auth.ForContext(ctx)
	_ = user

	return nil, nil
}
