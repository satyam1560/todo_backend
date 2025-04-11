package router

import (
	"github.com/gin-gonic/gin"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
	"github.com/satyam1560/todo_backend/internal/handlers"
	"github.com/satyam1560/todo_backend/api/middleware" // ✅ added
)

func RegisterRoutes(r *gin.Engine, q *db.Queries) {
	todoHandler := handlers.TodoHandler{Q: q}
	authHandler := handlers.AuthHandler{Q: q}

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.LoginWithFirebase)
		}

		// ✅ Protected by JWT middleware
		todos := api.Group("/todos", middleware.JWTAuthMiddleware())
		{
			todos.POST("", todoHandler.CreateTodoHandler)
			todos.GET("", todoHandler.GetTodosHandler)
			todos.GET("/:id", todoHandler.GetTodoHandler)
			todos.PUT("/:id", todoHandler.UpdateTodoHandler)
			todos.DELETE("/:id", todoHandler.DeleteTodoHandler)
		}
	}
}
