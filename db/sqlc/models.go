// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"time"
)

type Customer struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    int64  `json:"phone"`
	Address  string `json:"address"`
}

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Status     string    `json:"status"`
	OrderTime  time.Time `json:"order_time"`
}

type Payment struct {
	ID            int64  `json:"id"`
	PizzaID       int64  `json:"pizza_id"`
	CustomerID    int64  `json:"customer_id"`
	PaymentStatus string `json:"payment_status"`
	Bill          int64  `json:"bill"`
}

type Pizza struct {
	ID         int64  `json:"id"`
	OrderID    int64  `json:"order_id"`
	Price      int64  `json:"price"`
	PizzaType  string `json:"pizza_type"`
	PizzaQuant int64  `json:"pizza_quant"`
}
