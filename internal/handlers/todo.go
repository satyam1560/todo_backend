package handlers

import (
	"context"
	"database/sql"
	

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
)

type TodoHandler struct {
	Q *db.Queries
}



// POST /api/todos
func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	var input db.CreateTodoParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo, err := h.Q.CreateTodo(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GET /api/todos
func (h *TodoHandler) GetTodosHandler(c *gin.Context) {

	todos, err := h.Q.ListTodos(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GET /api/todos/:id
func (h *TodoHandler) GetTodoHandler(c *gin.Context) {
	todo, err := h.getTodoByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found or invalid ID"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// PUT /api/todos/:id
func (h *TodoHandler) UpdateTodoHandler(c *gin.Context) {
	id, err := parseUUIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input db.UpdateTodoParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	input.ID = id

	todo, err := h.Q.UpdateTodo(context.Background(), input)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DELETE /api/todos/:id
func (h *TodoHandler) DeleteTodoHandler(c *gin.Context) {
	id, err := parseUUIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Q.DeleteTodo(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ Todo successfully deleted"})
}

// Helper: get todo by ID
func (h *TodoHandler) getTodoByID(c *gin.Context) (db.Todo, error) {
	id, err := parseUUIDParam(c)
	if err != nil {
		return db.Todo{}, err
	}
	return h.Q.GetTodo(context.Background(), id)
}

// Helper: parse UUID from Gin param
func parseUUIDParam(c *gin.Context) (uuid.UUID, error) {
	idParam := c.Param("id")
	return uuid.Parse(idParam)
}
