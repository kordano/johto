package store

import (
	"os"
	"fmt"
	"database/sql"
    _ "github.com/lib/pq"

)

var connection *Connection

// Connection is primary sql db connection
type Connection struct {
	db *sql.DB
}

func connect() error {
	var err error
	var db *sql.DB
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	url := "postgres://"

	if postgresUser != "" {
		url += postgresUser
	} else {
		url += "johto"
	}

	url += ":"

	if postgresPassword != "" {
		url += postgresPassword
	} else {
		url += "boofar"
	}

	url += "@"

	if postgresUser != "" {
		url += "db:5432"
	} else {
		url += "localhost:5433"
	}

	url += "/johto?sslmode=disable"

	db, err = sql.Open("postgres", url)

	if err != nil {
		return err
	}
	connection = &Connection{db}
	return nil
}

// GetConnection creates or retrieves current sql connection
func GetConnection() (*Connection, error) {
	if connection == nil {
		err := connect()
		if err != nil {
			fmt.Println("db is not there")
			return connection, err
		}
	}

	return connection, nil

}
