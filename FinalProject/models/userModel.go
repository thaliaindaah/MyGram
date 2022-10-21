package models

import (
	"FinalProject/database"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUsers(user *User) (err error) {
	if err = database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (out User, err error) {
	err = database.DB.Table("Users").Where("email = ?", email).Last(&out).Error
	return
}

func GetUserById(user User, id int) (out User, err error) {
	err = database.DB.Table("Users").Where("id = ?", id).First(&user).Error
	out = user
	return
}

func UpdateUser(user *User, id interface{}) (err error) {
	err = database.DB.Table("Users").Where("id = ?", id).Update(&user).Error
	return err
}

func DeleteUser(user *User, id int) (err error) {
	err = database.DB.Table("Users").Where("id = ?", id).Delete(nil).Error
	return nil
}
