package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/http-service/database"
)

//ShowAllUsers will show all the user
func ShowAllUsers(ctx *gin.Context) {
	users := database.GetAll()
	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetSingleUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	i := uint(id)
	user := database.GetWithId(i)
	ctx.JSON(200, gin.H{
		"data": user,
	})
}
