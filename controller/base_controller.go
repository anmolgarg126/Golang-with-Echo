package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var Echo = echo.New()

// custom middleware

// type MiddlewareFunc func(HandleFunc) HandleFunc
// type HandlerFunc func(Context) error

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom middleware")
		return next(c)
	}
}

func preServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom pre middleware")
		c.Request().URL.Path = "/Anmol"
		fmt.Println("%+v\n", c.Request())
		return next(c)
	}
}

func Start() {
	port := os.Getenv("MY_GO_APP_PORT")
	if port == "" {
		port = "8080"
	}

	// .pre() will always execute first before .use(middleware) and middleware declared in GET/POST/PUT/DELETE method
	// I understood .pre() like this:   It excutes the middleware before the router. So before the request hits the router,
	//  you can make changes to it. In the example it is show to be used for the www-to-no-www redirect. the http://www.xyz.tld is changed to http://xyz.tld and than hits the router. At least this is how did understand it from the docs.

	// e.Pre(preServerMessage)
	// e.Use(serverMessage) // common middleware for all endpoints

	Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	createRoutesForUserController()

	Echo.Logger.Print(fmt.Sprintf("listening on port %s", port))
	Echo.Logger.Fatal(Echo.Start(fmt.Sprintf(":%s", port)))
	// e.Logger.Fatal(e.Start(":8080"))
}
