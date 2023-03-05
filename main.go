package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq" // or else can't talk to the DB
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/cindy-cyber/simpleBank/api"
	db "github.com/cindy-cyber/simpleBank/db/sqlc"
	"github.com/cindy-cyber/simpleBank/gapi"
	"github.com/cindy-cyber/simpleBank/pb"
	"github.com/cindy-cyber/simpleBank/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	go runGatewayServer(config, store)  // run the http gateway server in another go routine
	runGrpcServer(config, store)
}


func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {	
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)  // allows gRPC client to explore what RPCs are available on the server

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func runGatewayServer(config util.Config, store db.Store) {	
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	// the response body uses the snake format as in proto files for consistency
	proto_in_json := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(proto_in_json)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()  // cancelling a context is a way of preventing the system from doing unncessary work
	
	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Cannot register handelr server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)  // routing

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal("Cannot start HTTP gateway server:", err)
	}
}