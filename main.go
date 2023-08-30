package main

import (
	"database/sql"
	"github.com/Pedroxer/avito/api"
	"github.com/Pedroxer/avito/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/avito_backend?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}
	query := sqlc.New(db)
	server := api.NewServer(query)
	if err != nil {
		log.Fatal("cannot create a server")
	}
	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
