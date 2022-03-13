package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/d0kur0/cui-server/database"

	"github.com/d0kur0/cui-server/graph/gqlutils"

	"github.com/d0kur0/cui-server/auth"

	"github.com/d0kur0/cui-server/graph/model"
)

func (r *mutationResolver) CreateService(ctx context.Context, props *model.CreateServiceProps) (*model.Service, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, gqlutils.AccessDeniedError(ctx)
	}

	newService := model.Service{
		Name:        props.Name,
		Description: props.Description,
		Price:       props.Price,
		UserID:      user.ID,
		CreatedAt:   time.Now(),
	}

	var serviceModel database.ServiceModel

	createdService, err := serviceModel.Create(newService)
	if err != nil {
		return nil, err
	}

	return createdService, nil
}

func (r *mutationResolver) UpdateService(ctx context.Context, props *model.UpdateServiceProps) (*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Services(ctx context.Context, count *int) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}
