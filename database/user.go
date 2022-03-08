package database

import (
	"errors"

	"github.com/d0kur0/cui-server/graph/model"
)

func ValidateAndGetUser(authToken string) (user *model.User, err error) {
	db := GetDB()

	db.
		Joins("", db.Where(model.UserToken{Token: authToken})).
		First(&user, "token = ?", authToken)

	if user == nil {
		err = errors.New("authorization token is invalid")
		return
	}

	return
}
