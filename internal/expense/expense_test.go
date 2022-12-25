package expense

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	expenseJSON = `{"title":"buy a new phone","amount":39000,"note":"buy a new phone","tags":["food", "beverage"]}`
)

func TestCreateExpense(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := `INSERT INTO "expenses" ("title", "amount", "note", "tags") VALUES ($1,$2,$3,$4) RETURNING "expenses"."id"`
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs("buy a new phone", 39000, "buy a new phone", pq.Array([]string{"food", "beverage"})).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHandler(db)

	// Assertions
	if assert.NoError(t, h.CreateExpenseHandler(c)) {
		// assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, expenseJSON, rec.Body.String())
	}
}

func TestGetExpense(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expense := &Expense{
		Id:     1,
		Title:  "",
		Amount: 39000,
		Note:   "",
		Tags:   []string{"food", "beverage"},
	}

	query := `SELECT id, title, amount, note, tags FROM "expenses" WHERE (id = $1)`
	rows := sqlmock.
		NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(expense.Id, expense.Title, expense.Amount, expense.Note, expense.Tags)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(expense.Id).
		WillReturnRows(rows)

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := NewHandler(db)

	fmt.Println(expenseJSON)
	fmt.Println(rec.Body.String())

	// Assertions
	if assert.NoError(t, h.GetExpenseHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expenseJSON, rec.Body.String())
	}
}
