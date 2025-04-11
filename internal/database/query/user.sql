-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: CreateUser :one
INSERT INTO users (id, phone)
VALUES ($1, $2)
RETURNING *;
