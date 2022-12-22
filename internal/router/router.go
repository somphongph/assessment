package router

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// api := e.Group("")

	// 	expenseHandler := expense.NewHandler(db)
	// 	expense := api.Group("/expenses")
	// 	{
	// 		expense.POST("", expenseHandler.CreateExpenseHandler)
	// 	}

	log.Fatal(e.Start(os.Getenv("PORT")))

	return e
}
