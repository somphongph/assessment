package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/assessment/internal/model"
)

func (h *Handler) GetExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := h.db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: "can't prepare query user statment:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	e := Expense{}
	err = row.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, &e.Tags)
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, model.Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, model.Err{Message: "can't scan user:" + err.Error()})
	}
}
