package database

import (
	"context"
	"fmt"
	"github.com/UkralStul/TodoListTestTask/internal/config"
	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}

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
