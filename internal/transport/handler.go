package transport

import (
	"github.com/gofiber/fiber/v2"
)

type Todos interface {
	GetTodos(ctx *fiber.Ctx) []string
	InsertTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx)
	DeleteTodo(ctx *fiber.Ctx)
}

type Handler struct {
	todosService Todos
}

func NewHandler(td Todos) *Handler {
	return &Handler{
		todosService: td,
	}
}

func (h *Handler) IndexHandler(ctx *fiber.Ctx) error {

	todos := h.todosService.GetTodos(ctx)
	return ctx.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func (h *Handler) PostHandler(ctx *fiber.Ctx) error {
	err := h.todosService.InsertTodo(ctx)
	if err != nil {
		return err
	}
	return ctx.Redirect("/")
}

func (h *Handler) PutHandler(ctx *fiber.Ctx) error {
	h.todosService.UpdateTodo(ctx)
	return ctx.Redirect("/")
}

func (h *Handler) DeleteHandler(ctx *fiber.Ctx) error {
	h.todosService.DeleteTodo(ctx)
	return ctx.SendString("deleted")
}
