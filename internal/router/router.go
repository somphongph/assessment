package router

import (
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/somphongph/assessment/internal/expense"
)

func NewRouter() *echo.Echo {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	// create a new echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	expenseHandler := expense.NewHandler(db)
	expenseHandler.InitDB()
	expense := e.Group("/expenses")
	{
		expense.POST("", expenseHandler.CreateExpenseHandler)
		expense.GET("/:id", expenseHandler.GetExpenseHandler)
		expense.PUT("/:id", expenseHandler.UpdateExpenseHandler)
	}

	log.Fatal(e.Start(os.Getenv("PORT")))

	return e
}
