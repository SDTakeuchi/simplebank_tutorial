package main

import (
	"database/sql"
	"fmt"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	// dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

var dbConfig = fmt.Sprintf(
	"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// "localhost",
	// for localtest
	"postgres",
	"5432",
	"root",
	"simple_bank",
	"secret",
	"disable",
)

func main() {
	conn, err := sql.Open(dbDriver, dbConfig)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}