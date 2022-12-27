package expense

import "database/sql"

type Expense struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount uint     `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Handler struct {
	db *sql.DB
}
