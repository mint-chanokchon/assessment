package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/mint-chanokchon/assessment/handlers/expenses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(initExpensesTable)

	e.POST("/expenses", expenses.Create)
	e.GET("/expenses/:id", expenses.FindById)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func initExpensesTable(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			c.Error(err)
			log.Fatal(err)
			return err
		}

		defer db.Close()

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT, amount FLOAT, note TEXT, tags TEXT[] );`)
		if err != nil {
			c.Error(err)
			log.Fatal(err)
			return err
		}

		log.Println("Create successful")

		return next(c)
	}
}
