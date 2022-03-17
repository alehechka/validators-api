package validators_test

import (
	"testing"

	"github.com/alehechka/validators-api/validators"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

func Test_UniqueField(t *testing.T) {
	err := validators.AddValidator("uniquefield", validators.UniqueField)
	assert.Nil(t, err)

	var body = struct {
		First string `json:"first" binding:"required"`
		Last  string `json:"last" binding:"required,uniquefield=First"`
	}{
		First: "John",
		Last:  "Doe",
	}

	err = binding.Validator.ValidateStruct(body)
	assert.Nil(t, err)
}

func Test_UniqueField_NotUnique(t *testing.T) {
	err := validators.AddValidator("uniquefield", validators.UniqueField)
	assert.Nil(t, err)

	var body = struct {
		First string `json:"first" binding:"required"`
		Last  string `json:"last" binding:"required,uniquefield=First"`
	}{
		First: "John",
		Last:  "John",
	}

	err = binding.Validator.ValidateStruct(body)
	assert.NotNil(t, err)
}
