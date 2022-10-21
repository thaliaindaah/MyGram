package database

import "gorm.io/gorm"

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/MyGram")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(models.User{})
	return db
}
