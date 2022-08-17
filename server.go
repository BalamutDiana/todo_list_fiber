package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todos := repository.NewTodos(db)
	handler := transport.NewHandler(todos)
	engine := html.New("./views", ".html")
	app := handler.InitRouter(engine)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}
