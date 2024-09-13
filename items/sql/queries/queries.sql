-- name: GetItem :one
SELECT * FROM items WHERE id = ?;

-- name: ListItems :many
SELECT * FROM items;

-- name: ListFoods :many
SELECT * FROM items WHERE type = "FOOD";

-- name: ListToys :many
SELECT * FROM items WHERE type = "TOY";

-- name: CreateItem :execresult
INSERT INTO items (
  id,
  name,
  price,
  image,
  expiration_time,
  weight,
  suitable_for,
  material,
  type
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);