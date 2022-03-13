package database

import (
	"errors"

	"github.com/d0kur0/cui-server/graph/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ServiceModel struct{}

func (s *ServiceModel) Create(serviceForCreate model.Service) (createdService *model.Service, err error) {
	err = s.validate(&serviceForCreate)
	if err != nil {
		return nil, err
	}

	createResult := db.Create(&serviceForCreate)
	if createResult.Error != nil {
		return nil, errors.New("internal error")
	}

	return &serviceForCreate, nil
}

func (s *ServiceModel) validate(service *model.Service) (err error) {
	return validation.ValidateStruct(service,
		validation.Field(&service.Name, validation.Required, validation.Length(1, 128)),
		validation.Field(&service.Description, validation.Required, validation.Length(0, 512)),
		validation.Field(&service.Price, validation.Required),
	)
}
