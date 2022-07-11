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
WHERE pizza_id IS NOT NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePaymentStatus :one
UPDATE payment
SET payment_status = $2
WHERE id = $1
RETURNING *;