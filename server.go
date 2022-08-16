package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"

	"github.com/BalamutDiana/todo_list_fiber/internal/repository"
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
	todos := repository.NewTodos(db)
	handler := transport.NewHandler(todos)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return handler.IndexHandler(ctx)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return handler.PostHandler(ctx)
	})

	app.Put("/update", func(ctx *fiber.Ctx) error {
		return handler.PutHandler(ctx)
	})

	app.Delete("/delete", func(ctx *fiber.Ctx) error {
		return handler.DeleteHandler(ctx)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}
