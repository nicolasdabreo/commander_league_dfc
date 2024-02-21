package main

import (
	"log"
	"fmt"
	"dfc/db"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME string = "_data/dfc.db"
)

func main() {
	store, err := db.NewStore(DB_NAME)

	if err != nil {
		log.Fatal(err)	
	}
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
