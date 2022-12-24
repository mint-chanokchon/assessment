package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func Create(c echo.Context) error {
	db := open()
	defer db.Close()

	var expense Expense
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if expense.isEmpty() {
		return c.JSON(http.StatusBadRequest, "incomplete information")
	}

	queryString := `INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id`
	row := db.QueryRow(queryString, expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags))

	err = row.Scan(&expense.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, expense)
}
