-- name: GetCart :one
SELECT * FROM carts WHERE id = ?;

-- name: GetCartItems :many
SELECT * FROM cart_items
WHERE cart_id = ?;

-- name: GetCartItem :one
SELECT * FROM cart_items
WHERE id = ?;

-- name: GetCartItemByItemID :one
SELECT id
FROM cart_items
WHERE cart_id = ? AND item_id = ?;

-- name: AddCartItem :exec
INSERT INTO cart_items(id, item_id, qty, cart_id)
VALUES (?, ?, ?, ?);

-- name: CreateCart :exec
INSERT INTO carts(id)
VALUES (?);

-- name: UpdateCartItem :exec
UPDATE cart_items
SET qty = qty + ?
WHERE cart_id = ? AND item_id = ? AND id = ?;
