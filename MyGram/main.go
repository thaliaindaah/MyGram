package main

import (
	"FinalProject/database"
	"FinalProject/models"
	"FinalProject/router"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	r := router.StartApp()
	r.Run()
}
