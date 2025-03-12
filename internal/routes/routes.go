package routes

import (
	"github.com/UkralStul/TodoListTestTask/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, handler *handlers.DBHandler) {
	tasks := app.Group("/tasks")

	tasks.Get("/", handler.ReadAllTasks)
	tasks.Post("/", handler.CreateTask)
	tasks.Put("/:id", handler.UpdateTask)
	tasks.Delete("/:id", handler.DeleteTask)
}
