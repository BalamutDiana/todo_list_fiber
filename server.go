package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"

	"github.com/BalamutDiana/todo_list_fiber/pkg/database"
)

func main() {

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "1marvin2mode3",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {

	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todo")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}
func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello postHandler")
}
func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello putHandler")
}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello deleteHandler")
}
