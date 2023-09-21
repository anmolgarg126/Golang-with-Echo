package controller

import (
	"echo_framework/service"

	"github.com/labstack/echo/v4/middleware"
)

func createRoutesForUserController() {

	// Middleware
	// Echo.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Timeout: 5 * time.Millisecond,
	// }))

	Echo.POST("/users", service.SaveUser, middleware.BodyLimit("1K")) // middleware provided by echo
	Echo.GET("/users/id", service.GetUserById)
	Echo.GET("/users/jobs", service.FindByNumberOfJobsAUserHad)
	Echo.GET("/users/:experience/:salary", service.FindUserByExperienceAndSalary)
	Echo.GET("/users/having-names", service.FindUserHavingNameField)
	Echo.GET("/users/organisation", service.FindUserWorkedInGivenOrganisation)
	Echo.GET("/users/name", service.GetUserByName)
	Echo.GET("/users/all", service.GetAllUsers)
	// Echo.GET("/users/all", service.GetAllUsers, middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Skipper: Skipper,
	// 	ErrorHandler: func(err error, e echo.Context) error {
	// 		// you can handle your error here, the returning error will be
	// 		// passed down the middleware chain
	// 		return err
	// 	},
	// 	Timeout: 5 * time.Millisecond,
	// }))
	// e.GET("/users/all", getAllUsers, serverMessage) // route level middleware
	// e.PUT("/users/:id", updateUser)
	Echo.DELETE("/users", service.DeleteUser)

}
