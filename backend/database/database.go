package database

import (
	"database/sql"
	"log"
)

var pool *sql.DB

// Init the database connection
func Init(connstr string) error {
	log.Println("database: init")
	var err error
	pool, err = sql.Open("postgres", connstr)
	return err
}

// Close the database connection
func Close() {
	log.Println("database: close")
	pool.Close()
}
