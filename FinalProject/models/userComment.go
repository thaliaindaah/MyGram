package models

import (
	"FinalProject/database"
)

func CreateComment(comment *Comment) (err error) {
	if err = database.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetComment() (out []Comment, err error) {
	err = database.DB.Find(&out).Error
	return
}

func GetCommentById(comment Comment, id int) (out Comment, err error) {
	err = database.DB.Table("Comment").Where("id = ?", id).First(&comment).Error
	out = comment
	return
}

func UpdateComment(out *Comment, id interface{}) (err error) {
	err = database.DB.Table("Comment").Where("id = ?", id).Update(&out).Error
	return err
}

func DeleteComment(comment *Comment, id int) (err error) {
	err = database.DB.Table("Comment").Where("photo_id = ?", id).Delete(nil).Error
	return nil
}

func DeleteCommentUserId(comment *Comment, id int) (err error) {
	err = database.DB.Table("Comment").Where("user_id = ?", id).Delete(nil).Error
	return nil
}

func DeleteCommentById(comment *Comment, id int) (err error) {
	err = database.DB.Table("Comment").Where("id = ?", id).Delete(nil).Error
	return nil
}
