// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: payment.sql

package db

import (
	"context"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payment (
  pizza_id,
  customer_id,
  payment_status,
  bill
) VALUES (
  $1, $2, $3, $4
) RETURNING id, pizza_id, customer_id, payment_status, bill
`

type CreatePaymentParams struct {
	PizzaID       int64  `json:"pizza_id"`
	CustomerID    int64  `json:"customer_id"`
	PaymentStatus string `json:"payment_status"`
	Bill          int64  `json:"bill"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.PizzaID,
		arg.CustomerID,
		arg.PaymentStatus,
		arg.Bill,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.PizzaID,
		&i.CustomerID,
		&i.PaymentStatus,
		&i.Bill,
	)
	return i, err
}

const getPayment = `-- name: GetPayment :one
SELECT id, pizza_id, customer_id, payment_status, bill FROM payment
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, id int64) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.PizzaID,
		&i.CustomerID,
		&i.PaymentStatus,
		&i.Bill,
	)
	return i, err
}

const listPayments = `-- name: ListPayments :many
SELECT id, pizza_id, customer_id, payment_status, bill FROM payment
WHERE pizza_id = $1
OR customer_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListPaymentsParams struct {
	PizzaID    int64 `json:"pizza_id"`
	CustomerID int64 `json:"customer_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListPayments(ctx context.Context, arg ListPaymentsParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayments,
		arg.PizzaID,
		arg.CustomerID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.PizzaID,
			&i.CustomerID,
			&i.PaymentStatus,
			&i.Bill,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
