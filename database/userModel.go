package database

import (
	"encoding/json"
	"errors"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"golang.org/x/crypto/bcrypt"

	"github.com/d0kur0/cui-server/graph/model"
	"gorm.io/gorm"
)

type UserModel struct{}

func (u *UserModel) GetSafeStruct(user *model.User) (safeUser *model.SafeUser) {
	userAsJson, err := json.Marshal(user)
	if err != nil {
		return nil
	}

	if err = json.Unmarshal(userAsJson, &safeUser); err != nil {
		return nil
	}

	return
}

func (u *UserModel) GetByToken(token string) (user *model.User) {
	const joinQuery = `"join user_tokens on user_tokens.token = ?"`

	result := db.Joins(joinQuery, token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return
}

func (u *UserModel) GetByEmail(email string) (user *model.User) {
	result := db.Where(&model.User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return
}

func (u *UserModel) GetByEmailAndPassword(email string, password string) (user *model.User, err error) {
	result := db.Where(model.User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user with this email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	return
}

func (u *UserModel) validate(user *model.User) (err error) {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 128)),
		validation.Field(&user.Name, validation.Required, validation.Length(2, 32)),
	)
}

func (u *UserModel) Create(userForCreate model.User) (createdUser *model.User, err error) {
	err = u.validate(&userForCreate)
	if err != nil {
		return nil, err
	}

	userWithEqualEmail := u.GetByEmail(userForCreate.Email)
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
