-- name: CreateCustomer :one
INSERT INTO customers (
  full_name,
  email,
  phone,
  address
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: GetCustomerForUpdate :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListCustomers :many
SELECT * FROM customers
WHERE full_name IS NOT NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCustomerAddress :one
UPDATE customers
SET address = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;