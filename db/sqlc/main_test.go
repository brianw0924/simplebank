package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/brianw0924/simplebank/util"
	_ "github.com/lib/pq" // we don't actually call any function of lib/pq directly so the go fomatter will remove it automtically. We have to use blank identifier to keep it.
)

var testQueries *Queries
var testDB *sql.DB

// main entry point of all unit tests inside 1 specific golang package (here, we are in package db)
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
