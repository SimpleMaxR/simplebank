package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simplemaxr/simplebank/api"
	db "github.com/simplemaxr/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:Aa@09120828@localhost:5432/simple_bank?sslmode=disable"

	serverAddress = "0.0.0.0:12123"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
