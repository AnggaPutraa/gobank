package main

import (
	"database/sql"
	"log"

	"github.com/AnggaPutraa/gobank/api"
	db "github.com/AnggaPutraa/gobank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5220/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:9000"
)

func main() {
	var err error

	connection, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Can't connect to database with err,", err)
	}

	store := db.NewStore(connection)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Can't start the server: ", err)
	}
}
