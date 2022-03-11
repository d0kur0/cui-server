package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/d0kur0/cui-server/graph/model"
)

func (r *mutationResolver) CreateService(ctx context.Context, props *model.CreateServiceProps) (*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateService(ctx context.Context, props *model.UpdateServiceProps) (*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Services(ctx context.Context, count *int) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}
