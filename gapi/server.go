package gapi

import (
	"fmt"

	db "github.com/cindy-cyber/simpleBank/db/sqlc"
	"github.com/cindy-cyber/simpleBank/pb"
	"github.com/cindy-cyber/simpleBank/token"
	"github.com/cindy-cyber/simpleBank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store  // allow us to interact with the databse when processing API requests from clients
	tokenMaker token.Maker	
}

// NewServer creates a new gRPC server
func  NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
