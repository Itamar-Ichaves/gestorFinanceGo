package main

import (
	"database/sql"
	"log"

	api "github.com/Itamar-Ichaves/gestorFinanceGo/Api"
	db "github.com/Itamar-Ichaves/gestorFinanceGo/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://postgres:postgres@localhost:5432/gestorfinancego?sslmode=disable"
	serverAdress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
