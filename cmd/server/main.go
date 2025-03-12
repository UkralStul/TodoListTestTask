package main

import (
	"context"
	"github.com/UkralStul/TodoListTestTask/internal/config"
	"github.com/UkralStul/TodoListTestTask/internal/database"
	"github.com/UkralStul/TodoListTestTask/internal/handlers"
	"github.com/UkralStul/TodoListTestTask/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	err = database.InitTasksTable(context.Background(), db)
	if err != nil {
		log.Fatalf("Failed to init tasks table: %v", err)
	}

	handler := &handlers.DBHandler{Conn: db}

	app := fiber.New()

	routes.SetupRouter(app, handler)

	log.Fatal(app.Listen(":3000"))
}
