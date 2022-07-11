package api

import (
	"database/sql"
	"net/http"

	db "github.com/M1racle-Heen/simple_order_pizza/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPaymentRequest struct {
	PizzaID       int64  `json:"pizza_id" binding:"required,min=1"`
	CustomerID    int64  `json:"customer_id" binding:"required,min=1"`
	PaymentStatus string `json:"payment_status" binding:"required"`
}

func (server *Server) createPayment(ctx *gin.Context) {
	var req createPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePaymentParams{
		PizzaID:       req.PizzaID,
		CustomerID:    req.CustomerID,
		PaymentStatus: req.PaymentStatus,
	}

	PizzaID, err := server.store.GetPizza(ctx, req.PizzaID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg.Bill = PizzaID.Price

	payment, err := server.store.CreatePayment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

type getPaymentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPayment(ctx *gin.Context) {
	var req getPaymentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payment, err := server.store.GetPayment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

type listPaymentsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPayments(ctx *gin.Context) {
	var req listPaymentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPaymentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	payments, err := server.store.ListPayments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

type UpdatePaymentStatusRequest struct {
	ID            int64  `json:"id" binding:"required,min=1"`
	PaymentStatus string `json:"status" binding:"required,oneof=Paid NotPaid"`
}

func (server *Server) updatePaymentStatus(ctx *gin.Context) {
	var req UpdatePaymentStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePaymentStatusParams{
		ID:            req.ID,
		PaymentStatus: req.PaymentStatus,
	}

	payment, err := server.store.GetPayment(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pizza, err := server.store.GetPizza(ctx, payment.PizzaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := server.store.GetOrder(ctx, pizza.OrderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.PaymentStatus == "Paid" {
		order.Status = "Delivered"
	} else {
		order.Status = "Hold"
	}

	arg1 := db.UpdateOrderStatusParams{
		ID:     order.ID,
		Status: order.Status,
	}

	order, err = server.store.UpdateOrderStatus(ctx, arg1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	status, err := server.store.UpdatePaymentStatus(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, status)
}
