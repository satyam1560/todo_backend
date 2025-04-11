package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	// 🔥 Create Gin router and apply middleware
	r := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r.Use(gin.Recovery())   

	          // Handles panics
	r.Use(middleware.LoggerMiddleware()) // Your custom logger

	// 🧠 Register your routes
	router.RegisterRoutes(r, queries)

	// ✅ Start Gin server
	addr := ":8080"
	fmt.Printf("🚀 Server started at http://localhost%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
