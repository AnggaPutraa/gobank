package main

import (
	"database/sql"
	"log"

	"github.com/AnggaPutraa/gobank/api"
	db "github.com/AnggaPutraa/gobank/db/sqlc"
	"github.com/AnggaPutraa/gobank/utils"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Can't read the env configuration")
	}

	connection, err := sql.Open(
		conf.DbDriver,
		conf.DbSource,
	)

	if err != nil {
		log.Fatal("Can't connect to database with err,", err)
	}

	store := db.NewStore(connection)

	server, err := api.NewServer(conf, store)

	if err != nil {
		log.Fatal("Can't create the server: ", err)
	}

	err = server.Start(conf.ServerAddress)

	if err != nil {
		log.Fatal("Can't start the server: ", err)
	}
}
