package main

import (
	"github.com/alehechka/validators-api/rest"
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
				Text: "idk",
				Type: "post",
			},
		},
	},
}

func main() {

	router := rest.SetupRouter()

	router.Run()

	// rest.MockRequest("POST", "/validate", validBody)
}
