package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Todos interface {
	GetTodos(ctx *fiber.Ctx) []string
	InsertTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
}

type Handler struct {
	todosService Todos
}

func NewHandler(td Todos) *Handler {
	return &Handler{
		todosService: td,
	}
}

func (h *Handler) InitRouter(engine *html.Engine) *fiber.App {

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return h.IndexHandler(ctx)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return h.PostHandler(ctx)
	})

	app.Put("/update", func(ctx *fiber.Ctx) error {
		return h.PutHandler(ctx)
	})

	app.Delete("/delete", func(ctx *fiber.Ctx) error {
		return h.DeleteHandler(ctx)
	})

	return app
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
	err := h.todosService.UpdateTodo(ctx)
	if err != nil {
		return err
	}
	return ctx.Redirect("/")
}

func (h *Handler) DeleteHandler(ctx *fiber.Ctx) error {
	err := h.todosService.DeleteTodo(ctx)
	if err != nil {
		return err
	}
	return ctx.SendString("deleted")
}
