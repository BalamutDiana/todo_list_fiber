package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"

	"github.com/BalamutDiana/todo_list_fiber/internal/transport"
	"github.com/BalamutDiana/todo_list_fiber/pkg/database"
)

func main() {

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "password",
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
		return transport.IndexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return transport.PostHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return transport.PutHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return transport.DeleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}
