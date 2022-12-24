package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
