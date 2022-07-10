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
