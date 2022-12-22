package router

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/somphongph/assessment/internal/expense"
)

func NewRouter() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ex := e.Group("/expenses")
	{
		ex.POST("", expense.CreateExpenseHandler)
	}

	log.Fatal(e.Start(os.Getenv("PORT")))

	return e
}
