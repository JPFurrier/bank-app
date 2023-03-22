package main

import (
	"bank-app/api"
	db "bank-app/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
const serverAddress = "0.0.0.0:8080"

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
