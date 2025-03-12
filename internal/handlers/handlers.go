package handlers

import (
	"context"
	"github.com/UkralStul/TodoListTestTask/internal/database"
	"github.com/UkralStul/TodoListTestTask/internal/models"
	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v5"
	"strconv"
)

type DBHandler struct {
	conn *pgx.Conn
}

// Хендлер создания таски
func (h *DBHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}

	if task.Status != "" && task.Status != "new" && task.Status != "in_progress" && task.Status != "done" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status: must be 'new', 'in_progress', or 'done'",
		})
	}

	id, err := database.CreateTask(context.Background(), h.conn, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

// Обновление таски
func (h *DBHandler) UpdateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := database.UpdateTask(context.Background(), h.conn, task.ID, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": task})
}

// Получение всех тасок
func (h *DBHandler) ReadAllTasks(c *fiber.Ctx) error {
	tasks, err := database.GetAllTasks(context.Background(), h.conn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"tasks": tasks})
}

// Удаление таски по id
func (h *DBHandler) DeleteTask(c *fiber.Ctx) error {
	idFromParam := c.Params("id")

	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	err = database.DeleteTask(context.Background(), h.conn, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.SendStatus(fiber.StatusNoContent)
	return nil
}
