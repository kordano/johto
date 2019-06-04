package store

import ("database/sql"
	_  "github.com/lib/pq")

var connection *Connection

// Store is primary sql db connection
type Connection struct {
	db *sql.DB
}

func connect() error {
	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", "postgres://johto:boofar@localhost:5433/johto?sslmode=disable")

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
			return connection, err
		}
	}

	return connection, nil

}
