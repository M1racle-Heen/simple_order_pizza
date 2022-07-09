package api

type CreateOrderRequest struct {
	CustomerID int64  `json:"customer_id" binding:"required,min=1"`
	Status     string `json:"status"`
}
