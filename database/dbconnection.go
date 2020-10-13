package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/http-service/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "project"
	password = "12345"
	dbname   = "simplehttp"
)

// DB referencing the database
var DB *gorm.DB

//Connection connects the database and creates tables
func Connection() {
	dbInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DropTableIfExists(&models.Users{})
	db.AutoMigrate(&models.Users{})
	DB = db
}
