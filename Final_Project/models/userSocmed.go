package models

import "FinalProject/database"

func CreateSocmed(socmed *SocialMedia) (err error) {
	if err = database.DB.Create(&socmed).Error; err != nil {
		return err
	}
	return nil
}

func GetSocmed(id int) (out []SocialMedia, err error) {
	err = database.DB.Table("SocialMedia").Where("user_id = ?", id).Find(&out).Error
	return
}

func GetSocmedbyId(socmed SocialMedia, id int) (out SocialMedia, err error) {
	err = database.DB.Table("SocialMedia").Where("id = ?", id).First(&socmed).Error
	out = socmed
	return
}

func UpdateSocmed(out *SocialMedia, id interface{}) (err error) {
	err = database.DB.Table("SocialMedia").Where("id = ?", id).Update(&out).Error
	return err
}

func DeleteSocmed(socmed *SocialMedia, id int) (err error) {
	err = database.DB.Table("SocialMedia").Where("user_id = ?", id).Delete(nil).Error
	return nil
}

func DeleteSocmedByID(socmed *SocialMedia, id int) (err error) {
	err = database.DB.Table("SocialMedia").Where("id = ?", id).Delete(nil).Error
	return nil
}
