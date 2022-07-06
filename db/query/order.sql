-- name: CreateOrder :one
INSERT INTO orders (
  customer_id,
  status
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE customer_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;