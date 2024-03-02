package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSchema(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	getTestDb()
}

func getTestDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	err = RunMigrations(db, "../../../db/migrations")
	if err != nil {
		panic(err)
	}
	return db
}
