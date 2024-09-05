package utils

import (
	"food-delivery/internal/dto"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(val interface{}) []*dto.InvalidParams {
	err := validate.Struct(val)
	if err == nil {
		return nil
	}

	var invalidParams []*dto.InvalidParams
	for _, err := range err.(validator.ValidationErrors) {
		invalidParams = append(invalidParams, &dto.InvalidParams{
			Name:   err.Field(),
			Reason: err.ActualTag(),
		})
	}

	return invalidParams
}
