package expense

type Expense struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Amount uint     `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}
