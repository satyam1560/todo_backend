package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
)

type TodoHandler struct {
	Q *db.Queries
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// POST /api/todos
func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input db.CreateTodoParams
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo, err := h.Q.CreateTodo(context.Background(), input)
	if err != nil {
		http.Error(w, "Could not create todo", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

// GET /api/todos
func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Q.ListTodos(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

// GET /api/todos/{id}
func (h *TodoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo, err := h.getTodoByID(r)

	fmt.Println(todo)
	if err != nil {
		http.Error(w, "Todo not found or invalid ID", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

// PUT /api/todos/{id}
func (h *TodoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractUUID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var input db.UpdateTodoParams
	input.ID = id

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo, err := h.Q.UpdateTodo(context.Background(), input)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Could not update todo", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

// DELETE /api/todos/{id}
func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractUUID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Q.DeleteTodo(context.Background(), id)
	if err != nil {
		http.Error(w, "Could not delete todo", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "✅ Todo successfully deleted",
	})
}

// Helper: fetch todo by ID using SQLC
func (h *TodoHandler) getTodoByID(r *http.Request) (db.Todo, error) {
	id, err := extractUUID(r)
	if err != nil {
		return db.Todo{}, err
	}
	return h.Q.GetTodo(context.Background(), id)
}

// Helper: extract UUID from path
func extractUUID(r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars["id"])
}
