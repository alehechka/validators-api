package validators

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Validators is a global map of custom validators
var Validators = map[string]func(validator.FieldLevel) bool{
	"notoneof":    NotOneOf,
	"uniquefield": UniqueField,
}

// AddValidators adds selected validators from provided parameters or all validators if none are provided.
func AddValidators(bindings ...string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if len(bindings) == 0 {
			addAllValidators(v)
			return
		}
		addValidatorsByKeys(v, bindings)
	}
}

func addValidatorsByKeys(v *validator.Validate, bindings []string) {
	for _, binding := range bindings {
		if validatorFunc, exists := Validators[binding]; exists {
			v.RegisterValidation(binding, validatorFunc)
		}
	}
}

func addAllValidators(v *validator.Validate) {
	for key, validatorFunc := range Validators {
		v.RegisterValidation(key, validatorFunc)
	}
}

// AddValidator allows for one off validator functions to be registered safely. Will return a non-nil error if the global Validator object is unavailable.
func AddValidator(bindingKey string, validatorFunc func(fl validator.FieldLevel) bool) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation(bindingKey, validatorFunc)
		return nil
	}

	return errors.New("global validator engine not available")
}
