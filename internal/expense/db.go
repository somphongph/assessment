package expense

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitDB() {
	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := h.db.Exec(createTb); err != nil {
		log.Fatal("can't create table", err)
	}

}
