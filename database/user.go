package database

import (
	"github.com/d0kur0/cui-server/graph/model"
)

func ValidateAndGetUser(authToken string) (user *model.User) {
	db.
		Joins("join user_tokens on user_tokens.token = ?", authToken).
		First(&user)

	return
}
