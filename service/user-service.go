package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/http-service/database"
	"github.com/http-service/models"
)

//ShowAllUsers will show all the user
func ShowAllUsers(ctx *gin.Context) {
	var users []models.Users
	database.DB.Find(&users)
	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetSingleUser(ctx *gin.Context) {
	var user models.Users
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = uint(id)
	fmt.Println("ajay")
	database.DB.Where("id = ?", user.ID).Find(&user)
	ctx.JSON(200, gin.H{
		"data": user,
	})
}
