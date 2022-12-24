package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}
