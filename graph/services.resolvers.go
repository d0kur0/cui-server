package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/d0kur0/cui-server/auth"

	"github.com/d0kur0/cui-server/graph/generated"
	"github.com/d0kur0/cui-server/graph/model"
)

func (r *queryResolver) Services(ctx context.Context, count *int) ([]*model.Service, error) {
	_ = auth.ForContext(ctx)

	return nil, errors.New("test")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
