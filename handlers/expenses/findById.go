package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func FindById(c echo.Context) error {
	id := c.Param("id")
	db := open()
	defer db.Close()

	queryString := `SELECT * FROM expenses WHERE id=$1`
	row := db.QueryRow(queryString, id)

	var expense Expense

	err := row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(expense.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, expense)
}
