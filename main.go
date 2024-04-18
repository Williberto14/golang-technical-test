package main

import (
	"golang-technical-test/src/common/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":7777")
}
