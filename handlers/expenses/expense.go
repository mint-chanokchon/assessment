package expenses

import (
	"database/sql"
	"log"
	"os"
)

type Expense struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func open() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection Failure")
	}

	return db
}

func (e *Expense) isEmpty() bool {
	return len(e.Title) == 0 || len(e.Note) == 0
}
