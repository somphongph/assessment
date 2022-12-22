package expense

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
	"github.com/somphongph/assessment/internal/model"
)

func CreateExpenseHandler(c echo.Context) error {
	e := Expense{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	tags := e.Tags
	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id",
		e.Title,
		e.Amount,
		e.Note,
		pq.Array(&tags),
	)
	err = row.Scan(&e.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}
