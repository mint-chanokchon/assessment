package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

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
	e.PUT("/expenses/:id", expenses.Update)
	e.GET("/expenses", expenses.FindAll)

	go func() {
		err := e.Start(os.Getenv("PORT"))
		if err != nil {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	err := e.Shutdown(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}
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
