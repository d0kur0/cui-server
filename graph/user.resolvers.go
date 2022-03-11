package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/d0kur0/cui-server/auth"
	"github.com/d0kur0/cui-server/database"
	"github.com/d0kur0/cui-server/graph/generated"
	"github.com/d0kur0/cui-server/graph/model"
)

func (r *mutationResolver) SignIn(ctx context.Context, props model.SignInProps) (*model.SafeUser, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	var userModel database.UserModel
	var userTokenModel database.UserTokenModel

	signedUser, err := userModel.GetByEmailAndPassword(props.Email, props.Password)
	if err != nil {
		return nil, err
	}

	createdToken, err := userTokenModel.Create(signedUser.ID)
	if err != nil {
		return nil, err
	}

	var tokens []*model.UserToken
	tokens = append(tokens, createdToken)
	signedUser.Tokens = tokens

	return userModel.GetSafeStruct(signedUser), nil
}

func (r *mutationResolver) SignUp(ctx context.Context, props model.SignUpProps) (*model.SafeUser, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	var userModel database.UserModel
	var userTokenModel database.UserTokenModel

	newUser := model.User{
		Name:      props.Name,
		Email:     props.Email,
		Password:  props.Password,
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

	return userModel.GetSafeStruct(&newUser), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.SafeUser, error) {
	return nil, errors.New("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
