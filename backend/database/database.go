package database

import (
	"database/sql"
	"log"
	"time"
)

// Timestamps has the common created_at and updated_at fields
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
