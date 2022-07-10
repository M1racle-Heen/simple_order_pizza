package api

import (
	"database/sql"
	"net/http"

	db "github.com/M1racle-Heen/simple_order_pizza/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPizzaRequest struct {
	OrderID    int64  `json:"order_id" binding:"required,min=1"`
	Price      int64  `json:"price" binding:"required,min=1000"`
	PizzaType  string `json:"pizza_type" binding:"required,oneof=Cheese Veggie Pepperoni Meat Margherita BBQChicken Hawaiian Buffalo"`
	PizzaQuant int64  `json:"pizza_quant" binding:"required,min=1"`
}

func (server *Server) createPizza(ctx *gin.Context) {
	var req createPizzaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePizzaParams{
		OrderID:    req.OrderID,
		Price:      req.Price,
		PizzaType:  req.PizzaType,
		PizzaQuant: req.PizzaQuant,
	}

	pizza, err := server.store.CreatePizza(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}

type getPizzaRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPizza(ctx *gin.Context) {
	var req getPizzaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pizza, err := server.store.GetPizza(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}

type listPizzasRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPizzas(ctx *gin.Context) {
	var req listPizzasRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPizzasParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	pizzas, err := server.store.ListPizzas(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pizzas)
}
