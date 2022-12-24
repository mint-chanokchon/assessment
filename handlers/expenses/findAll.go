package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindAll(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
