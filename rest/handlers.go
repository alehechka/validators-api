package rest

import "github.com/gin-gonic/gin"

func RegisterHandlers(router *gin.Engine) {
	router.POST("/validate", validate)
}
