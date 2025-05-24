package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	testDBName = "simple_bank_test"
)

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	
	testQueries = New(conn)

	os.Exit(m.Run())
}