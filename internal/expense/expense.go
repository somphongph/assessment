package expense

type Expense struct {
	Id     uint     `json:"id"`
	Title  string   `json:"title"`
	Amount uint     `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}
