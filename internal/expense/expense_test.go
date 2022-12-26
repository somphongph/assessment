//go:build unit
// +build unit

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

	bodyFailed = bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": "79",
		"note": 20,
		"tags": "food", "beverage"
	}`)
)

func TestCreateExpenseHandler(t *testing.T) {
	t.Run("Create Expense Success", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		tags := expense.Tags
		mockedSql := "INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
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
	})

	t.Run("Bind Data Expense Failed", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		tags := expense.Tags
		mockedSql := "INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
		mockedRow := sqlmock.NewRows([]string{"id"}).AddRow(1)

		mock.ExpectQuery(regexp.QuoteMeta(mockedSql)).
			WithArgs(expense.Title, expense.Amount, expense.Note, pq.Array(&tags)).
			WillReturnRows((mockedRow))

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bodyFailed)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewHandler(db)

		// Assertions
		if assert.NoError(t, h.CreateExpenseHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Create Expense Failed", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		tags := expense.Tags
		mockedSql := "INSERT INTO expenses_failed (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id"
		mockedRow := sqlmock.NewRows([]string{"id"}).AddRow(1)

		mock.ExpectQuery(regexp.QuoteMeta(mockedSql)).
			WithArgs(expense.Title, expense.Amount, expense.Note, pq.Array(&tags)).
			WillReturnRows((mockedRow))

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", bodyFailed)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewHandler(db)

		// Assertions
		if assert.NoError(t, h.CreateExpenseHandler(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestGetByIdExpenseHandler(t *testing.T) {
	t.Run("Get Expense by ID Success", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		id := 1
		tags := expense.Tags
		mockedSql := "SELECT id, title, amount, note, tags FROM expenses WHERE id = $1"
		mockedRow := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
			AddRow(id, expense.Title, expense.Amount, expense.Note, pq.Array(&tags))

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
	})
}

func TestGetAllExpenseHandler(t *testing.T) {
	t.Run("Get All Expense Success", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		tags := expense.Tags
		mockedSql := "SELECT id, title, amount, note, tags FROM expenses"
		mockedRow := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
			AddRow(1, expense.Title, expense.Amount, expense.Note, pq.Array(&tags)).
			AddRow(2, expense.Title, expense.Amount, expense.Note, pq.Array(&tags))

		mock.ExpectPrepare(regexp.QuoteMeta(mockedSql)).ExpectQuery().
			WillReturnRows((mockedRow))

		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewHandler(db)

		// Assertions
		if assert.NoError(t, h.GetAllExpenseHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUpdateExpenseHandler(t *testing.T) {
	t.Run("Update Expense Success", func(t *testing.T) {
		// Mock
		db, mock, _ := sqlmock.New()

		id := 1
		tags := expense.Tags
		mockedSql := "UPDATE expenses SET title = $2, amount = $3, note = $4, tags = $5 WHERE id = $1"
		mockedRow := sqlmock.NewResult(1, 1)

		mock.ExpectPrepare(regexp.QuoteMeta(mockedSql)).ExpectExec().
			WithArgs(strconv.Itoa(id), expense.Title, expense.Amount, expense.Note, pq.Array(&tags)).
			WillReturnResult(mockedRow)

		// Setup
		b, err := json.Marshal(expense)
		if err != nil {
			fmt.Println(err)
			return
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(id))
		h := NewHandler(db)

		// Assertions
		if assert.NoError(t, h.UpdateExpenseHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
