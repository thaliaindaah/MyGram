package controllers

import (
	"FinalProject/helpers"
	"FinalProject/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var User models.User

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := models.CreateUsers(&User)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	temp, err := models.GetUserByEmail(User.Email)
	fmt.Println(User.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(temp.Password), []byte(User.Password))
	fmt.Println(User.Password, temp.Password)
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Password",
		})
		return
	}
	fmt.Println(temp.ID, "id")
	token := helpers.GenerateToken(temp.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func UpdateUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	var User models.User
	id := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := models.UpdateUser(&User, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	user, err := models.GetUserById(User, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"age":          user.Age,
		"email":        user.Email,
		"id":           user.ID,
		"username":     user.Username,
		"updated_date": user.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	var User models.User
	var Photo models.Photo
	var Socmed models.SocialMedia
	var Comment models.Comment
	id := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = id

	err := models.DeleteSocmed(&Socmed, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	err = models.DeleteCommentUserId(&Comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	err = models.DeletePhoto(&Photo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	err = models.DeleteUser(&User, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})
}
