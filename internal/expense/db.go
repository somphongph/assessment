package expense

import (
	"log"

	_ "github.com/lib/pq"
)

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
