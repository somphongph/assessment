package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/somphongph/assessment/internal/model"
)

func (h *Handler) GetAllExpenseHandler(c echo.Context) error {
	stmt, err := h.db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: "can't prepare query expenses statment:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: "can't query all expenses:" + err.Error()})
	}

	expenses := []Expense{}
	for rows.Next() {
		var e Expense
		err = rows.Scan(&e.Id, &e.Title, &e.Amount, &e.Note, &e.Tags)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.Err{Message: "can't scan expenses:" + err.Error()})
		}
		expenses = append(expenses, e)
	}

	return c.JSON(http.StatusOK, expenses)
}
