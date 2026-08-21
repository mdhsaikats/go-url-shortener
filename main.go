package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"main.go/app/config"
	"main.go/app/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	ctx := context.Background()
	pool, err := config.InitPool(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()
	log.Println("Database connection pool initialized successfully!")
	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.ApiRouters(r)

	fmt.Printf("Running on port: %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}