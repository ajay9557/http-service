package main

import (
	"github.com/gin-gonic/gin"
	"github.com/http-service/service"
	"github.com/http-service/database"

)

func main() {
	server := gin.Default()
	database.Connection()

	server.GET("/users", service.ShowAllUsers)
	server.GET("/users/:id", service.GetSingleUser)

	server.Run(":8080")
}
