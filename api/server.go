package api

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/validator/v10"
	db "github.com/cindy-cyber/simpleBank/db/sqlc"
	// "github.com/techschool/simplebank/token"
	// "github.com/techschool/simplebank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	// config     util.Config
	store      db.Store  // allow us to interact with the databse when processing API requests from clients
	// tokenMaker token.Maker	
	router     *gin.Engine	// help send each API request to the correct handler for processing
}

// NewServer creates a new HTTP server and set up routing.
func  NewServer(/*config util.Config, */store db.Store) *Server/*(*Server, error)*/ {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token maker: %w", err)
	// }

	server := &Server{
		// config:     config,
		store:      store,
		// tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }
	
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	// server.setupRouter()
	return server //, nil
}

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	router.POST("/users", server.createUser)
// 	router.POST("/users/login", server.loginUser)
// 	router.POST("/tokens/renew_access", server.renewAccessToken)

// 	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
// 	authRoutes.POST("/accounts", server.createAccount)
// 	authRoutes.GET("/accounts/:id", server.getAccount)
// 	authRoutes.GET("/accounts", server.listAccounts)

// 	authRoutes.POST("/transfers", server.createTransfer)

// 	server.router = router
// }

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}