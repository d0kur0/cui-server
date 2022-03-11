package database

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/d0kur0/cui-server/graph/model"
	"github.com/d0kur0/cui-server/graph/scalars"
	"gorm.io/gorm"
)

type UserModel struct{}

func (u *UserModel) GetByToken(token string) (user *model.User) {
	result := db.
		Joins("join user_tokens on user_tokens.token = ?", token).
		First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return
}

func (u *UserModel) GetByEmail(email string) (user *model.User) {
	result := db.Where(&model.User{Email: scalars.Email(email)}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return
}

func (u *UserModel) GetByEmailAndPassword(email string, password string) (user *model.User) {
	return
}

func (u *UserModel) Create(userForCreate model.User) (createdUser *model.User, err error) {
	userWithEqualEmail := u.GetByEmail(string(userForCreate.Email))
	if userWithEqualEmail != nil {
		return nil, errors.New("email already in use")
	}

	hashOfPassword, err := bcrypt.GenerateFromPassword([]byte(userForCreate.Password), 14)
	if err != nil {
		log.Printf("error on UserModel.Create bcrypt.GenerateFromPassword; %s", err)
		return nil, errors.New("internal error")
	}

	userForCreate.Password = string(hashOfPassword)

	userCreateResult := db.Create(&userForCreate)
	if userCreateResult.Error != nil {
		log.Printf("error on SignUp db.Create(user); %s", err)
		return nil, errors.New("internal error")
	}

	return &userForCreate, nil
}
