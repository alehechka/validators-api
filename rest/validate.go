package rest

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Body struct {
	FirstName   string   `json:"firstName" binding:"required"`
	LastName    string   `json:"lastName" binding:"required,unique=FirstName"`
	Email       string   `json:"email" binding:"required,email"`
	Phone       string   `json:"phone" binding:"required,e164"`
	CountryCode string   `json:"countryCode" binding:"required,iso3166_1_alpha2"`
	ProductCode string   `json:"productCode" binding:"required,startswith=PC,len=10"`
	Width       int      `json:"width" binding:"required,gte=1,lte=100,gtfield=Height"`
	Height      int      `json:"height" binding:"required,gte=-1,lte=100"`
	User        User     `json:"user" binding:"required"`
	Things      []string `json:"things" binding:"required,unique"`
}

type User struct {
	Name     string    `json:"name" binding:"required,min=3"`
	Age      uint      `json:"age" binding:"required,min=18"`
	Comments []Comment `json:"comments" binding:"required,min=1,dive"`
}

type Comment struct {
	Text string `json:"text" binding:"required,max=255,notoneof=idk idc"`
	Type string `json:"type" binding:"required,oneof=post nested"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "notoneof":
		return "Value cannot be one of the following: " + fe.Param()
	case "unique":
		if fe.Kind() == reflect.Slice {
			return "All values in array must be unique"
		}
		return "Value must be different than " + fe.Param()
	}

	return "Unknown error - " + fe.Error()
}

func createErrors(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}
	return nil
}

func AddValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		addNotOneOfBinding(v)
		addUniqueBinding(v)
	}
}

func addNotOneOfBinding(v *validator.Validate) {
	v.RegisterValidation("notoneof", func(fl validator.FieldLevel) bool {
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
	})
}

func addUniqueBinding(v *validator.Validate) {
	v.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
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
		case reflect.Slice:
			length := fl.Field().Len()
			hash := make(map[interface{}]bool, 0)
			for i := 0; i < length; i++ {
				value := fl.Field().Index(i)
				if _, exists := hash[value.Interface()]; exists {
					return false
				}
				hash[value.Interface()] = true
			}
		}
		return true
	})
}

func validate(c *gin.Context) {
	var body Body
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": createErrors(err)})
		return
	}

	c.JSON(http.StatusAccepted, &body)
}
