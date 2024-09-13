
-- name: GetPet :one
SELECT * FROM pets WHERE ID = ?;

-- name: CreatePet :exec
INSERT INTO pets (id, name, birthday, breed, user_id)
VALUES (?, ?, ?, ?, ?);