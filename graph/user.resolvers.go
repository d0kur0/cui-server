package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/d0kur0/cui-server/auth"
	"github.com/d0kur0/cui-server/database"
	"github.com/d0kur0/cui-server/graph/generated"
	"github.com/d0kur0/cui-server/graph/model"
)

func (r *mutationResolver) SignIn(ctx context.Context, props model.SignInProps) (*model.UserToken, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	return nil, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, props model.SignUpProps) (*model.User, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	var userModel database.UserModel
	var userTokenModel database.UserTokenModel

	newUser := model.User{
		Name:      props.Name,
		Email:     props.Email,
		Password:  string(props.Password),
		CreatedAt: time.Now(),
	}

	createdUser, err := userModel.Create(newUser)
	if err != nil {
		return nil, err
	}

	createdToken, err := userTokenModel.Create(createdUser.ID)
	if err != nil {
		return nil, err
	}

	var tokens []*model.UserToken
	tokens = append(tokens, createdToken)
	newUser.Tokens = tokens

	return &newUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
