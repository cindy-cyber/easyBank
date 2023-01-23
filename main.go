package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // or else can't talk to the DB

	"github.com/cindy-cyber/simpleBank/api"
	db "github.com/cindy-cyber/simpleBank/db/sqlc"
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
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}