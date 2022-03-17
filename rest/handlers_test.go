package rest_test

import (
	"testing"

	"github.com/alehechka/validators-api/rest"
	"github.com/stretchr/testify/assert"
)

func Test_Ping(t *testing.T) {
	w := rest.MockRequest("GET", "/ping", nil)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
