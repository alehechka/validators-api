package rest_test

import (
	"testing"

	"github.com/alehechka/validators-api/rest"
	"github.com/stretchr/testify/assert"
)

var validBody = rest.Body{
	FirstName:   "John",
	LastName:    "Doe",
	Email:       "johndoe@example.com",
	Phone:       "+11234567890",
	CountryCode: "US",
	ProductCode: "PC12345678",
	Width:       99,
	Height:      1,
	Things:      []string{"thing1", "thing2"},
	User: rest.User{
		Name: "John",
		Age:  19,
		Comments: []rest.Comment{
			{
				Text: "solid",
				Type: "post",
			},
		},
	},
}

func Test_Validate(t *testing.T) {
	w := rest.MockRequest("POST", "/validate", validBody)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
