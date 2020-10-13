package main

import (
	"github.com/gin-gonic/gin"
	"github.com/http-service/database"
	"github.com/http-service/service"
)

func main() {
	server := Setup()
	server.Run(":8080")

	server.GET("/users", service.ShowAllUsers)
	server.GET("/users/:id", service.GetSingleUser)

}

func Setup() *gin.Engine {
	database.Connection()

	router := gin.Default()

	router.GET("/users", service.ShowAllUsers)
	router.GET("/users/:id", service.GetSingleUser)

	return router
}
