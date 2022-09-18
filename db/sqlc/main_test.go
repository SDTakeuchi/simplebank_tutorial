package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	// dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var dbConfig = fmt.Sprintf(
	"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	"postgres",
	"5432",
	"root",
	"simple_bank",
	"secret",
	"disable",
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbConfig)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
