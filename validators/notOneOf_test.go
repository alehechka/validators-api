package validators_test

import (
	"testing"

	"github.com/alehechka/validators-api/validators"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

func Test_NotOneOf_Match(t *testing.T) {
	err := validators.AddValidator("notoneof", validators.NotOneOf)
	assert.Nil(t, err)

	var body = struct {
		Name string `json:"first" binding:"required,notoneof=idk idc"`
	}{
		Name: "idk",
	}

	err = binding.Validator.ValidateStruct(body)
	assert.NotNil(t, err)
}

func Test_NotOneOf_NoMatch(t *testing.T) {
	err := validators.AddValidator("uniquefield", validators.UniqueField)
	assert.Nil(t, err)

	var body = struct {
		Name string `json:"first" binding:"required,notoneof=idk idc"`
	}{
		Name: "John",
	}

	err = binding.Validator.ValidateStruct(body)
	assert.Nil(t, err)
}
