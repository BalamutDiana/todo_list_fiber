package repository

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	Item string
}

type Todo struct {
	db *sql.DB
}

func NewTodos(db *sql.DB) *Todo {
	return &Todo{db}
}

func (td *Todo) GetTodos(ctx *fiber.Ctx) []string {
	var res string
	var todos []string
	rows, err := td.db.Query("SELECT * FROM todo")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		ctx.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return todos
}

func (td *Todo) InsertTodo(ctx *fiber.Ctx) error {
	newTodo := Item{}
	if err := ctx.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return ctx.SendString(err.Error())
	}

	if newTodo.Item != "" {
		_, err := td.db.Exec("INSERT into todo VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
	return nil
}

func (td *Todo) UpdateTodo(ctx *fiber.Ctx) {
	olditem := ctx.Query("olditem")
	newitem := ctx.Query("newitem")
	td.db.Exec("UPDATE todo SET item=$1 WHERE item=$2", newitem, olditem)
}

func (td *Todo) DeleteTodo(ctx *fiber.Ctx) {
	todoToDelete := ctx.Query("item")
	td.db.Exec("DELETE from todo WHERE item=$1", todoToDelete)
}
