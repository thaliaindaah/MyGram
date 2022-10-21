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
			ID:        id,
			Password:  Updated.Password,
			CreatedAt: Updated.CreatedAt,
			UpdatedAt: Updated.UpdatedAt,
			Age:       Updated.Age,
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

// func PhotoAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id, _ := strconv.Atoi(ctx.Param("id"))
// 		userData := ctx.MustGet("userData").(jwt.MapClaims)
// 		userId := int(userData["id"].(float64))
// 		Updated := models.Photo{}

// 		temp := models.Photo{
// 			ID:        id,
// 			CreatedAt: Updated.CreatedAt,
// 			UpdatedAt: Updated.UpdatedAt,
// 			PhotoURL:  Updated.PhotoURL,
// 			Caption:   Updated.Caption,
// 			Title:     Updated.Title,
// 			UserID:    Updated.UserID,
// 		}

// 		_, err := models.GetPhotoById(temp, id)
// 		fmt.Println(err)
// 		fmt.Println(temp.UserID)

// 		if temp.UserID != userId {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "unauthorized",
// 				"message": "you are not allowed to access this data",
// 			})
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
