-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: Login :one
SELECT * FROM users WHERE username = ?;

-- name: CreateUser :execresult
INSERT INTO users(id, username, password, name, type)
VALUES (?, ?, ?, ?, ?);

