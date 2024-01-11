package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // we don't actually call any function of lib/pq directly so the go fomatter will remove it automtically. We have to use blank identifier to keep it.
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

// main entry point of all unit tests inside 1 specific golang package (here, we are in package db)
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
