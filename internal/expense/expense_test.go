package expense

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	expenseBody = bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags":["food", "beverage"]
	}`)
)

func TestCreateExpenseHandler(t *testing.T) {
	db, mock, _ := sqlmock.New()

	ex := Expense{
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	tags := ex.Tags
	mockedSql := `INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id`
	mockedRow := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(mockedSql)).
		WithArgs(ex.Title, ex.Amount, ex.Note, pq.Array(&tags)).
		WillReturnRows((mockedRow))

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", expenseBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHandler(db)

	// Assertions
	if assert.NoError(t, h.CreateExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
