package expense

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	expenseJSON = `{"title":"buy a new phone","amount":39000,"note":"buy a new phone"}`
)

func TestCreateExpense(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mock.ExpectBegin()
	// // // mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "expenses" ("title", "amount", "note")
	// 	VALUES ($1, $2, $3) RETURNING id`)).
	// 	WithArgs("buy a new phone", 39000, "buy a new phone").
	// 	WillReturnResult(sqlmock.NewResult(1, 1))
	// // mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// // mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(expenseJSON))
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
