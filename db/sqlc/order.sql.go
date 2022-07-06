// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  customer_id,
  status
) VALUES (
  $1, $2
) RETURNING id, customer_id, status, order_time
`

type CreateOrderParams struct {
	CustomerID int64  `json:"customer_id"`
	Status     string `json:"status"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.CustomerID, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.Status,
		&i.OrderTime,
	)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT id, customer_id, status, order_time FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.Status,
		&i.OrderTime,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, customer_id, status, order_time FROM orders
WHERE customer_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListOrdersParams struct {
	CustomerID int64 `json:"customer_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.CustomerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.Status,
			&i.OrderTime,
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