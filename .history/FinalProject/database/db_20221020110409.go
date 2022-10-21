package database

import "gorm.io/gorm"

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db-go-sql")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(controllers.Order{}, controllers.Item{})
	return db
}
