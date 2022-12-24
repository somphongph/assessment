package expense

import (
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
