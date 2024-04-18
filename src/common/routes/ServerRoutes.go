package routes

import (
	"golang-technical-test/src/common/middlewares"
	"golang-technical-test/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/students", handlers.GetStudents)
	router.GET("/students/:id", middlewares.JWTAuthMiddleware(), handlers.GetStudent)
	router.POST("/students", middlewares.JWTAuthMiddleware(), handlers.CreateStudent)
	router.PUT("/students/:id", middlewares.JWTAuthMiddleware(), handlers.UpdateStudent)
	router.DELETE("/students/:id", middlewares.JWTAuthMiddleware(), handlers.DeleteStudent)

	router.POST("/login", handlers.LoginHandler)

	return router
}
