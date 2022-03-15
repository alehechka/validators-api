package rest

func RegisterHandlers(router *gin.Engine) {
	router.POST("/validate", validate)
}
