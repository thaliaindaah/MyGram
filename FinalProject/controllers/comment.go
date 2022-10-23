package controllers

import (
	"FinalProject/helpers"
	"FinalProject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	var Comment models.Comment
	id := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = id
	err := models.CreateComment(&Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {
	var Comment models.Comment
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	id := int(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserID = id
	temp, err := models.GetComment(id)
	for i, v := range temp {
		item, _ := models.GetItemByID(v.UserID)
		photoItem, err := models.GetItemPhotoByID(v.UserID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		temp[i].User = item
		temp[i].Photo = photoItem
	}
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, temp)
	}
}

func UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var Comment models.Comment
	commentId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = commentId

	err := models.UpdateComment(&Comment, commentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := models.GetCommentById(Comment, commentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"message":    user.Message,
		"user_id":    user.UserID,
		"photo_id":   user.PhotoID,
	})
}

func DeleteComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var Comment models.Comment
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.ID = id

	err := models.DeleteCommentById(&Comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}
