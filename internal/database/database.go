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
