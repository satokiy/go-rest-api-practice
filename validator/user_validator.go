package validator

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-rest-api/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

// UserValidate implements IUserValidator
func (*userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("email must be between 1 and 30 characters"),
			is.Email.Error("email is invalid"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(4, 30).Error("password must be between 4 and 30 characters"),
		),
	)
}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}
