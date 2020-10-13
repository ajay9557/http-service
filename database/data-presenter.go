package database

import (
	"github.com/http-service/models"
)

func GetAll() []models.Users {
	var users []models.Users
	DB.Find(&users)
	return users
}

func GetWithId(id uint) models.Users {
	var user models.Users
	DB.Where("id = ?", id).Find(&user)
	return user
}
