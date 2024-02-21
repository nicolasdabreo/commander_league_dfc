package main

import (
	"dfc/db"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store, _ := db.NewStore()

	defer store.Db.Close()

	tables := []string{
		"players",
		"results",
		"leagues",
		"schema_migrations",
	}

	for _, table := range tables {
		query := fmt.Sprintf("drop table if exists %s", table)
		if _, err := store.Db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Tables for database have been dropped")
}
