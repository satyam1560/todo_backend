// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: todo.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (title, completed)
VALUES ($1, $2)
RETURNING id, title, completed
`

type CreateTodoParams struct {
	Title     string
	Completed bool
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.Title, arg.Completed)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Completed)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, completed FROM todos
WHERE id = $1
`

func (q *Queries) GetTodo(ctx context.Context, id uuid.UUID) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Completed)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, completed FROM todos
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.Completed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $2, completed = $3
WHERE id = $1
RETURNING id, title, completed
`

type UpdateTodoParams struct {
	ID        uuid.UUID
	Title     string
	Completed bool
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.ID, arg.Title, arg.Completed)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Completed)
	return i, err
}
