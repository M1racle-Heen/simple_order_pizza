-- name: CreatePizza :one
INSERT INTO pizza (
  order_id,
  price,
  pizza_type,
  pizza_quant
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPizza :one
SELECT * FROM pizza
WHERE id = $1 LIMIT 1;

-- name: ListPizzas :many
SELECT * FROM pizza
WHERE order_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;