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
WHERE customer_id IS NOT NULL
ORDER BY id
LIMIT $1
OFFSET $2;