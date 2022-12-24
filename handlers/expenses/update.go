package expenses

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func Update(c echo.Context) error {
	id := c.Param("id")
	db := open()
	defer db.Close()

	var expense Expense

	queryString := `SELECT * FROM expenses WHERE id=$1`
	row := db.QueryRow(queryString, id)
	err := row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	err = c.Bind(&expense)
	expense.Id = id
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if expense.isEmpty() {
		return c.JSON(http.StatusBadRequest, "incomplete information")
	}

	queryString = `UPDATE expenses SET title=$1, amount=$2, note=$3, tags=$4 where id=$5`
	_, err = db.Exec(queryString, expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, expense)
}
