package database

import (
	"context"
	"fmt"
	"github.com/UkralStul/TodoListTestTask/internal/config"
	"github.com/UkralStul/TodoListTestTask/internal/models"
	"github.com/jackc/pgx/v5"
)

// Подключение к базе данных
func Connect(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}

// Создание таблицы с тасками если еще не создана
func InitTasksTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS tasks (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            description TEXT,
            status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
            created_at TIMESTAMP DEFAULT NOW(),
            updated_at TIMESTAMP DEFAULT NOW()
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create table tasks: %v", err)
	}
	return nil
}

// Создание таски
func CreateTask(ctx context.Context, conn *pgx.Conn, task models.Task) (int, error) {
	var id int
	err := conn.QueryRow(ctx, `
        INSERT INTO tasks (title, description, status)
        VALUES ($1, $2, $3)
        RETURNING id
    `, task.Title, task.Description, task.Status).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %v", err)
	}
	return id, nil
}

// Получение одной таски по id
func GetTask(ctx context.Context, conn *pgx.Conn, id int) (models.Task, error) {
	var task models.Task
	err := conn.QueryRow(ctx, `
        SELECT id, title, description, status, created_at, updated_at
        FROM tasks
        WHERE id = $1
    `, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to get task: %v", err)
	}
	return task, nil
}

// Получение всех тасок
func GetAllTasks(ctx context.Context, conn *pgx.Conn) ([]models.Task, error) {
	rows, err := conn.Query(ctx, `
        SELECT id, title, description, status, created_at, updated_at
        FROM tasks
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %v", err)
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}
	return tasks, nil
}

// Обновление таски
func UpdateTask(ctx context.Context, conn *pgx.Conn, id int, task models.Task) error {
	_, err := conn.Exec(ctx, `
        UPDATE tasks
        SET title = $1, description = $2, status = $3
        WHERE id = $4
    `, task.Title, task.Description, task.Status, id)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	return nil
}

// Удаление таски
func DeleteTask(ctx context.Context, conn *pgx.Conn, id int) error {
	_, err := conn.Exec(ctx, `
        DELETE FROM tasks
        WHERE id = $1
    `, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}
	return nil
}
