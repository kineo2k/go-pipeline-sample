package server

import (
	validator2 "github.com/go-playground/validator"
)

type RestApiValidator struct {
	validator *validator2.Validate
}

func (rav *RestApiValidator) Validate(i interface{}) error {
	return rav.validator.Struct(i)
}
