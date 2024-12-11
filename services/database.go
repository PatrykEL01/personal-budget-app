package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DbConnect(ctx context.Context) (*pgx.Conn, error) {
	dir, _ := os.Getwd()
	log.Printf("Current working directory: %s\n", dir)

	// Load .env file
	err := godotenv.Load("/app/.env")
	// Load .env file from root directory if not found in /app
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file from both /app/.env and .env: %v", err)
		}
	}
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return nil, errors.New("Error loading .env file")
	}

	// Initialize environment variables
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Println("DATABASE_URL is not set")
		return nil, errors.New("DATABASE_URL is not set")
	}

	// Connect to the database
	log.Printf("Connecting to database: %s\n", dbUrl)
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to the database successfully")
	return conn, nil
}

func InitializeSchema(ctx context.Context, conn *pgx.Conn) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS personal_budget (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			amount NUMERIC(10, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)`

	_, err := conn.Exec(ctx, createTableQuery)
	if err != nil {
		return err

	}

	return nil

}
