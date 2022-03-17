package rest

import (
	"github.com/alehechka/validators-api/validators"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	validators.AddValidators()

	RegisterHandlers(router)

	return router
}

func RegisterHandlers(router *gin.Engine) {
	router.POST("/validate", validate)
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
}
