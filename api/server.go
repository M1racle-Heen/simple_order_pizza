package api

import (
	db "github.com/M1racle-Heen/simple_order_pizza/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/customers", server.createCustomer)
	router.GET("/customers/:id", server.getCustomer)
	router.GET("/customers", server.listCustomers)

	server.router = router
	return server
}

//Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
