package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	database "main.go/app/config"
	"main.go/app/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	err := database.DatabaseConnection()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	defer database.DB.Close()
	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"*"}, // or "http://localhost:3000"
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300,
}))
	routes.ApiRouters(r)

	fmt.Printf("Running on port: %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}