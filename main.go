package main

func main() {
	router := gin.Default()

	rest.RegisterHandlers(router)

	router.Run()
}
