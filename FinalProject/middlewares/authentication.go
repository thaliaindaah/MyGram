package middlewares

import (
	"FinalProject/helpers"
	"FinalProject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verfiyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verfiyToken)
		ctx.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		fmt.Println(userData, "data")
		userId := int(userData["id"].(float64))
		Updated := models.User{}

		temp := models.User{
			ID:       id,
			Username: Updated.Username,
			Email:    Updated.Email,
		}

		_, err := models.GetUserById(temp, userId)
		fmt.Println(&temp, userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
		}
		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		fmt.Println("id", id)
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		var Updated models.Photo

		temp := models.Photo{
			ID:       id,
			Title:    Updated.Title,
			Caption:  Updated.Caption,
			PhotoURL: Updated.PhotoURL,
			UserID:   userId,
		}
		fmt.Println("temp models", temp)
		_, err := models.GetPhotoById(temp, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
		}
		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		fmt.Println("id", id)
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		temp := models.Comment{
			ID:     id,
			UserID: userId,
		}
		fmt.Println("temp models", temp)
		_, err := models.GetCommentById(temp, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
		}
		ctx.Next()
	}
}

func SocmedAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		fmt.Println("id", id)
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))

		temp := models.SocialMedia{
			ID:     id,
			UserID: userId,
		}
		fmt.Println("temp models", temp)
		_, err := models.GetSocmedbyId(temp, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
		}
		ctx.Next()
	}
}
