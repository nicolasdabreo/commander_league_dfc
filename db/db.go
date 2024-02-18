package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

func NewStore(dbName string) (Store, error) {
	Db, err := getConnection(dbName)

	if err != nil {
		return Store{}, err
	}

	return Store{
		Db,
	}, nil
}

func getConnection(dbName string) (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	if db != nil {
		return db, nil
	}

	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}

