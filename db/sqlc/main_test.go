package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/AnggaPutraa/gobank/utils"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	conf, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatal("Can't read the env configuration")
	}

	testDb, err = sql.Open(conf.DbDriver, conf.DbSource)

	if err != nil {
		log.Fatal("Can't connect to database with err,", err)
	}

	testQuery = New(testDb)

	os.Exit(m.Run())
}
