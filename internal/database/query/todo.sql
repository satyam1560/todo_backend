-- name: CreateTodo :one
INSERT INTO todos (title, completed)
VALUES ($1, $2)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, completed = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;
