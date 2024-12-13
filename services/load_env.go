package services

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"errors"
)


func LoadEnv() error {
	dir, _ := os.Getwd()
	log.Printf("Current working directory: %s\n", dir)

	possiblePaths := []string{"/app/.env", ".env"}

	var lastErr error
	for _, path := range possiblePaths {
		err := godotenv.Load(path)
		if err == nil {
			log.Printf("Successfully loaded .env file from: %s\n", path)
			return nil
		}
		lastErr = err
		log.Printf("Failed to load .env from: %s, error: %v\n", path, err)
	}

	if lastErr != nil {
		return errors.New("failed to load .env file from any of the specified paths")
	}

	return nil
}