package common_configs

import (
	"database/sql"
	"log"
)

// DB connection
var db *sql.DB

// Connect to DB
func ConnectDB() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:qwerty@tcp(localhost)/golang_technical_test")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
