-- name: CreatePayment :one
INSERT INTO payment (
  pizza_id,
  customer_id,
  payment_status,
  bill
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPayment :one
SELECT * FROM payment
WHERE id = $1 LIMIT 1;

-- name: ListPayments :many
SELECT * FROM payment
WHERE pizza_id = $1
OR customer_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;