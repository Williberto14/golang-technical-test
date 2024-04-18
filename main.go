package main

import (
	app "golang-technical-test/src"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.StartServer()
}
