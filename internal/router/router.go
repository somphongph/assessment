package router

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("apidesign")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("45678")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	expenseHandler := expense.NewHandler(db)
	expenseHandler.InitDB()
	expense := e.Group("/expenses")
	{
		expense.GET("/:id", expenseHandler.GetByIdExpenseHandler)
		expense.GET("", expenseHandler.GetAllExpenseHandler)
		expense.POST("", expenseHandler.CreateExpenseHandler)
		expense.PUT("/:id", expenseHandler.UpdateExpenseHandler)
	}

	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return e
}
