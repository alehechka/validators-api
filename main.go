package main

import (
	"github.com/alehechka/validators-api/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rest.AddValidators()

	rest.RegisterHandlers(router)

	router.Run()
}
