package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	id := c.Param("id")
	db := open()
	defer db.Close()

	return c.JSON(http.StatusOK, id)
}
