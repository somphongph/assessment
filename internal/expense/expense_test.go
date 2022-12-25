package expense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	expense = Expense{
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}
)

func TestCreateExpenseHandler(t *testing.T) {
	// Mock
	db, mock, _ := sqlmock.New()

	tags := expense.Tags
	mockedSql := `INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id`
	mockedRow := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(mockedSql)).
		WithArgs(expense.Title, expense.Amount, expense.Note, pq.Array(&tags)).
		WillReturnRows((mockedRow))

	// Setup
	b, err := json.Marshal(expense)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHandler(db)

	// Assertions
	if assert.NoError(t, h.CreateExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetByIdExpenseHandler(t *testing.T) {
	// Mock
	db, mock, _ := sqlmock.New()

	id := 1
	tags := expense.Tags
	mockedSql := `SELECT id, title, amount, note, tags FROM expenses WHERE id = $1`
	mockedRow := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(id, expense.Title, expense.Amount, expense.Note, (*pq.StringArray)(&tags))

	mock.ExpectPrepare(regexp.QuoteMeta(mockedSql)).ExpectQuery().
		WithArgs(strconv.Itoa(id)).
		WillReturnRows((mockedRow))

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	h := NewHandler(db)

	// Assertions
	if assert.NoError(t, h.GetByIdExpenseHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
