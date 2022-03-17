package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

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
