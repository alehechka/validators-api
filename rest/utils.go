package rest

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

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
			v.RegisterValidation(binding, validatorFunc.(func(fl validator.FieldLevel) bool))
		}
	}
}

func addAllValidators(v *validator.Validate) {
	for key, validatorFunc := range Validators {
		v.RegisterValidation(key, validatorFunc.(func(fl validator.FieldLevel) bool))
	}
}

var Validators = map[string]interface{}{
	"notoneof":    NotOneOf,
	"uniquefield": UniqueField,
}

func addNotOneOfBinding(v *validator.Validate) {
	v.RegisterValidation("notoneof", NotOneOf)
}

func NotOneOf(fl validator.FieldLevel) bool {
	// split values using ` `. eg. notoneof=bob rob job
	match := strings.Split(fl.Param(), " ")
	// convert field value to string
	value := fl.Field().String()
	for _, s := range match {
		// match value with struct filed tag
		if s == value {
			return false
		}
	}
	return true
}

func addUniqueBinding(v *validator.Validate) {
	v.RegisterValidation("uniquefield", UniqueField)
}

func UniqueField(fl validator.FieldLevel) bool {
	// get the fields which need to be unique
	match := strings.Split(fl.Param(), " ")
	// check if value is a string
	switch fl.Field().Kind() {
	case reflect.String:
		// value of the field
		value := fl.Field().String()
		for _, s := range match {
			// access to struct and getting value by field name
			fs := fl.Top().FieldByName(s)
			// check only for string validation
			if fs.Kind() == reflect.String {
				// check value of both fields
				if value == fs.String() {
					return false
				}
			}
		}
	}
	return true
}
