package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5220/simple_bank?sslmode=disable"
)

var testQuery *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Can't connect to database with err,", err)
	}

	testQuery = New(testDb)

	os.Exit(m.Run())
}
