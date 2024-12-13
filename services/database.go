package services

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func DbConnect(ctx context.Context, dbUrl string) (*pgx.Conn, error) {



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
