package main

import (
    "os"
    "log"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/sqlite"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    m, err := migrate.New(
        "file://cmd/migrate/migrations",
        "sqlite://_data/dfc.db",
    )
    if err != nil {
       log.Fatal(err)   
    }
    defer m.Close()

    cmd := os.Args[len(os.Args)-1]

    if cmd == "up" {
        log.Printf("Up")
        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	        log.Printf("Failed with %s", err)                                         
	}
    }

    if cmd == "down" {
	
		log.Printf("Down")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			 log.Printf("Failed with %s", err)                                       
		}
	}

}

