package api

import (
	"fmt"

	db "github.com/AnggaPutraa/gobank/db/sqlc"
	"github.com/AnggaPutraa/gobank/token"
	"github.com/AnggaPutraa/gobank/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Can't create token maker")
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authenticatedRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authenticatedRoutes.POST("/accounts", server.createAccount)
	authenticatedRoutes.GET("/accounts/:id", server.getAccount)
	authenticatedRoutes.GET("/accounts", server.getAccounts)

	authenticatedRoutes.POST("/transfer", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
