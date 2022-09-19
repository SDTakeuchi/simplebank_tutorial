package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	// dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB
var dbConfig = fmt.Sprintf(
	"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	"localhost",
	// for localtest
	// "postgres",
	"5432",
	"root",
	"simple_bank",
	"secret",
	"disable",
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbConfig)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
