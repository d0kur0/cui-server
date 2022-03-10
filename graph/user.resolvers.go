package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/d0kur0/cui-server/auth"
	"github.com/d0kur0/cui-server/database"
	"github.com/d0kur0/cui-server/graph/generated"
	"github.com/d0kur0/cui-server/graph/model"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) SignIn(ctx context.Context, props model.SignInProps) (*model.UserToken, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	return nil, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, props model.SignUpProps) (*model.UserToken, error) {
	user := auth.ForContext(ctx)
	if user != nil {
		return nil, errors.New("already signed")
	}

	hashOfPassword, err := bcrypt.GenerateFromPassword([]byte(props.Password), 14)
	if err != nil {
		log.Printf("error on signUp bcrypt.GenerateFromPassword; %s", err)
		return nil, errors.New("internal error")
	}

	newUser := model.User{
		Name:      props.Name,
		Email:     props.Email,
		Password:  string(hashOfPassword),
		CreatedAt: time.Now(),
	}

	db := database.GetDB()
	userCreateResult := db.Create(&newUser)

	if userCreateResult.Error != nil {
		log.Printf("error on SignUp db.Create(user); %s", err)
		return nil, errors.New("internal error")
	}

	tokenValue, err := ksuid.NewRandom()
	if err != nil {
		log.Printf("error on SignUp ksuid.NewRandom; %s", err)
		return nil, errors.New("internal error")
	}

	newToken := model.UserToken{
		UserID:    newUser.ID,
		Token:     tokenValue.String(),
		CreatedAt: time.Now(),
	}

	tokenCreateResult := db.Create(&newToken)
	if tokenCreateResult.Error != nil {
		log.Printf("error on SignUp db.Create(userToken); %s", err)
		return nil, errors.New("internal error")
	}

	return &newToken, nil
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
