package models

import (
	"FinalProject/database"
	"fmt"
)

func CreatePhoto(photo *Photo) (err error) {
	if err = database.DB.Create(&photo).Error; err != nil {
		return err
	}
	return nil
}

func GetPhoto() (out []Photo, err error) {
	err = database.DB.Find(&out).Error
	return
}

func GetPhotoById(photo Photo, id int) (out Photo, err error) {
	err = database.DB.Table("Photo").Where("id = ?", id).First(&photo).Error
	out = photo
	fmt.Println("Photo", photo)
	fmt.Println("Out", out)
	return
}

func GetItemPhotoByID(id int) (out []Photo, err error) {
	err = database.DB.Table("Photo").Where("user_id = ?", id).Scan(&out).Error
	return
}

func UpdatePhoto(out *Photo, id interface{}) (err error) {
	err = database.DB.Table("Photo").Where("id = ?", id).Update(&out).Error
	return err
}

func DeletePhoto(photo *Photo, id int) (err error) {
	err = database.DB.Table("Photo").Where("user_id = ?", id).Delete(nil).Error
	return nil
}

func DeletePhotoByID(photo *Photo, id int) (err error) {
	err = database.DB.Table("Photo").Where("id = ?", id).Delete(nil).Error
	return nil
}
