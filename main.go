package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"personal-budget/routes"
	"personal-budget/services"
)

func main() {
	log.Println("Starting application...")


	ctx := context.Background()
	
	log.Println("Connecting to the database...")
	// Connect to the database
	conn, err := services.DbConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer conn.Close(ctx)
	log.Println("Initializing schema...")
	// Initialize schema
	err = services.InitializeSchema(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	
	log.Println("Database connected and schema initialized")
	log.Println("Starting server...")
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run()
}
