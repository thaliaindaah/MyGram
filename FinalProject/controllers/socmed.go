package controllers

import (
	"FinalProject/helpers"
	"FinalProject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocmed(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	var Socmed models.SocialMedia
	id := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = id
	err := models.CreateSocmed(&Socmed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Socmed.ID,
		"name":             Socmed.Name,
		"social_media_url": Socmed.URL,
		"user_id":          Socmed.UserID,
	})
}

func GetSocmed(c *gin.Context) {
	var Socmed models.SocialMedia
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	id := int(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}
	Socmed.UserID = id
	temp, err := models.GetSocmed()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, temp)
	}
}

func UpdateSocmed(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var Socmed models.SocialMedia
	commentId, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.ID = commentId

	err := models.UpdateSocmed(&Socmed, commentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := models.GetSocmedbyId(Socmed, commentId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteSocmed(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var Socmed models.SocialMedia
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}
	Socmed.ID = id

	err := models.DeleteSocmed(&Socmed, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
