package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/M1racle-Heen/simple_order_pizza/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSourse)
	if err != nil {
		log.Fatal("cannot connect to Db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
