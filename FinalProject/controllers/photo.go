package controllers

import (
	"FinalProject/helpers"
	"FinalProject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	var Photo models.Photo
	id := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = id
	err := models.CreatePhoto(&Photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	var Photo models.Photo
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	id := int(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = id
	temp, err := models.GetPhoto()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, temp)
	}
}

func UpdatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var Photo models.Photo
	photoId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.ID = photoId

	err := models.UpdatePhoto(&Photo, photoId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := models.GetPhotoById(Photo, photoId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)

}
