package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/somphongph/assessment/internal/model"
)

func (h *Handler) UpdateExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	e := Expense{}
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	stmt, err := h.db.Prepare("UPDATE expenses SET title = $2, amount = $3, note = $4, tags = $5 WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Message: "can't prepare statment update:" + err.Error()})
	}

	tags := e.Tags
	if _, err := stmt.Exec(id, e.Title, e.Amount, e.Note, pq.Array(&tags)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Err{Message: "error execute update:" + err.Error()})
	}

	return c.JSON(http.StatusOK, e)
}
