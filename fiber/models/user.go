package models

import "github.com/go-playground/validator/v10"

type Job struct {
	Type   string `validate:"required,min=3,max=32"`
	Salary int    `validate:"required,number"`
}

type User struct {
	Name string `validate:"required,min=3,max=32"`
	// use `*bool` here otherwise the validation will fail for `false` values
	// Ref: https://github.com/go-playground/validator/issues/319#issuecomment-339222389
	IsActive *bool  `validate:"required"`
	Email    string `validate:"required,email,min=6,max=32"`
	Job      Job    `validate:"dive"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
