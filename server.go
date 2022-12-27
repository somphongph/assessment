package main

import (
	"crypto/subtle"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/somphongph/assessment/internal/router"
)

func main() {
	// create a new echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("apidesign")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("45678")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	// Router
	router.NewRouter(e)

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
