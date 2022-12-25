package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/somphongph/assessment/internal/model"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateExpenseHandler(c echo.Context) error {
	e := Expense{}

	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	tags := e.Tags
	row := h.db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id",
		e.Title,
		e.Amount,
		e.Note,
		pq.Array(&tags),
	)
	err := row.Scan(&e.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}
