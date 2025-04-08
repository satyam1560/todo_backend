package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/satyam1560/todo_backend/api/middleware"
	"github.com/satyam1560/todo_backend/api/router"
	"github.com/satyam1560/todo_backend/internal/database"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	// Initialize DB connection
	dbConn := database.InitDB()

	// Initialize SQLC query handler
	queries := db.New(dbConn)
	// Pass queries into the router
	r := router.NewRouter(queries)

	loggedRouter := middleware.Logger(r)

	addr := ":8080"
	fmt.Printf("🚀 Server started at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, loggedRouter); err != nil {
		log.Fatal(err)
	}
}
