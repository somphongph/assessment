package main

import (
	"fmt"
	"os"

	"github.com/somphongph/assessment/internal/expense"
	"github.com/somphongph/assessment/internal/router"
)

func main() {
	expense.InitDB()

	router.NewRouter()

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
