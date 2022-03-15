package rest

type Body struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required,e164"`
	CountryCode string `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

func validate(c *gin.Context) {
	var body Body
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("%#v\n", body)
}
