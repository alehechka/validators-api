package validators

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func UniqueField(fl validator.FieldLevel) bool {
	fmt.Println("Validating uniquefield...")
	// get the fields which need to be unique
	match := strings.Split(fl.Param(), " ")
	// check if value is a string
	if fl.Field().Kind() == reflect.String {
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
