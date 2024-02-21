package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

func NewStore() (Store, error) {
	Db, err := getConnection()

	if err != nil {
		return Store{}, err
	}

	return Store{
		Db,
	}, nil
}

func getConnection() (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	if db != nil {
		return db, nil
	}

	sqlitePath, envSet := os.LookupEnv("SQLITE_PATH")
	if !envSet {
		log.Fatal("🔥 Environment variable SQLITE_PATH not set")
	}

	db, err = sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}
