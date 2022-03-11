package database

import (
	"errors"
	"log"
	"time"

	"github.com/d0kur0/cui-server/graph/model"
	"github.com/segmentio/ksuid"
)

type UserTokenModel struct{}

func (ut *UserTokenModel) Create(userID int) (createdToken *model.UserToken, err error) {
	tokenValue, err := ksuid.NewRandom()
	if err != nil {
		log.Printf("error on UserTokenModel.Create ksuid.NewRandom; %s", err)
		return nil, errors.New("internal error")
	}

	newToken := model.UserToken{
		UserID:    userID,
		Token:     tokenValue.String(),
		CreatedAt: time.Now(),
	}

	result := db.Create(&newToken)
	if result.Error != nil {
		return nil, err
	}

	return &newToken, nil
}
