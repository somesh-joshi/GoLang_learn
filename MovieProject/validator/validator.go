package validator

import (
	"github.com/go-playground/validator/v10"
)

func Validator(data interface{}) error {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
