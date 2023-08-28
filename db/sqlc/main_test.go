package sqlc

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	testDB, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/avito_backend?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())

}
