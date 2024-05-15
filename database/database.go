package database

import (
	"database/sql"
	"fmt"
	"golang-technical-test/config"
)

// Database struct
type Database struct {
	*sql.DB
}

// NewDatabase
// It creates a new database connection.
func NewDatabase(config *config.DBConfig) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}
