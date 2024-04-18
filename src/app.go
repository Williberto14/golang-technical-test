package src

import "golang-technical-test/src/common/routes"

func StartServer() {
	router := routes.SetupRouter()
	router.Run(":7777")
}
