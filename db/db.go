package db

import (
	"database/sql"
	"fmt"
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

	if err := createMigrations(dbName, Db); err != nil {
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

	// Init SQLite3 database
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		// log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}

func createMigrations(dbName string, db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS players (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		deck VARCHAR(255) NOT NULL
	);`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	// stmt = `CREATE TABLE IF NOT EXISTS results (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	player_id INTEGER NOT NULL,
	// 	pod_size INTEGER NOT NULL,
	// 	place INTEGER NOT NULL,
	// 	the_council_of_wizards BOOLEAN DEFAULT(FALSE),
	// 	david_and_the_goliaths BOOLEAN DEFAULT(FALSE),
	// 	untouchable BOOLEAN DEFAULT(FALSE),
	// 	cleave BOOLEAN DEFAULT(FALSE),
	// 	its_free_real_estate BOOLEAN DEFAULT(FALSE),
	// 	i_am_timmy BOOLEAN DEFAULT(FALSE),
	// 	big_bigger_huge BOOLEAN DEFAULT(FALSE),
	// 	close_but_no_cigar BOOLEAN DEFAULT(FALSE),
	// 	just_as_garfield_intended BOOLEAN DEFAULT(FALSE),
	// 	created_at DATETIME default CURRENT_TIMESTAMP,
	// 	FOREIGN KEY(player_id) REFERENCES players(id)
	// );`

	// _, err = db.Exec(stmt)
	// if err != nil {
	// 	return err
	// }

	return nil
}
