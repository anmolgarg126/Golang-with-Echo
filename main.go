package main

import (
	"echo_framework/config"
	"echo_framework/controller"
)

// "github.com/go-ozzo/ozzo-validation/v4"

// "github.com/go-playground/validator/v10"
// "github.com/labstack/echo/v4"
// "github.com/go-ozzo/ozzo-validation/v4/is"

// var v *validator.Validate

func init() {
	config.LoadConfig()
}

func main() {

	controller.Start()

}

// func (u User) Validate() error {
// 	return validation.ValidateStruct(&u,
// 		// Street cannot be empty, and the length must between 5 and 50
// 		validation.Field(&u.Name, validation.Required, validation.Length(5, 50)),
// 		// City cannot be empty, and the length must between 5 and 50
// 		validation.Field(&u.Id, validation.Required, validation.Length(5, 50)),
// 		// // State cannot be empty, and must be a string consisting of two letters in upper case
// 		// validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
// 		// // State cannot be empty, and must be a string consisting of five digits
// 		// validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
// 	)
// }

// e.POST("/users", func(c echo.Context) error {
// 	u := new(User)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusCreated, u)
// 	// or
// 	// return c.XML(http.StatusCreated, u)
// })
