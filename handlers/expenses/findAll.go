package expenses

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func FindAll(c echo.Context) error {
	db := open()
	defer db.Close()

	querString := `SELECT * FROM expenses`
	rows, err := db.Query(querString)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	var expenses []Expense
	for rows.Next() {
		var expense Expense

		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		expenses = append(expenses, expense)
	}

	return c.JSON(http.StatusOK, expenses)
}
