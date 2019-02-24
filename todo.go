package main

import (
	"fmt"

	"demo-go-postgres/handlers"

	"github.com/go-pg/pg"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// For displaying the generated SQL in the command-line.
type dbLogger struct{}

func main() {
	db := initDB()
	defer db.Close()

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	port := ":8000"
	fmt.Print("App running at http://localhost", port)
	e.Run(standard.New(port))
}

func initDB() *pg.DB {
	options := &pg.Options{
		User:     "logan",
		Password: "xmen",
		Database: "marvel",
	}
	db := pg.Connect(options)
	db.AddQueryHook(dbLogger{})
	err := migrate(db)
	if err != nil {
		panic(err)
	}
	return db
}

// Create the table
func migrate(db *pg.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Tasks(id BIGSERIAL PRIMARY KEY, name text)`)

	if err != nil {
		return err
	}
	return nil
}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

// Show the generated SQL in the command-line.
func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
