package router

import (
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
	"github.com/satyam1560/todo_backend/internal/handlers"
)

func NewRouter(q *db.Queries) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	todoHandler := handlers.TodoHandler{Q: q}

	r.HandleFunc("/api/todos", todoHandler.CreateTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/todos", todoHandler.GetTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todos/{id}", todoHandler.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todos/{id}", todoHandler.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/todos/{id}", todoHandler.DeleteTodoHandler).Methods(http.MethodDelete)

	return r
}
