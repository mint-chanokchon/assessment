package expenses

import (
	"database/sql"
	"log"
	"os"
)

type Expense struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func open() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection Failure")
	}

	defer db.Close()
}
