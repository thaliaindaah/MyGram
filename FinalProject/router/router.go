package router

import (
	"FinalProject/controllers"
	"FinalProject/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.POST("/users/register", controllers.UserRegister)
	r.POST("/users/login", controllers.UserLogin)
	r.PUT("/users", middlewares.Authentication(), controllers.UpdateUser)
	r.DELETE("/users", middlewares.Authentication(), controllers.DeleteUser)
	r.POST("/photos", middlewares.Authentication(), controllers.CreatePhoto)
	r.GET("/photos", middlewares.Authentication(), controllers.GetPhoto)
	r.PUT("/photos/:id", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
	r.DELETE("/photos/:id", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	r.POST("/comments", middlewares.Authentication(), controllers.CreateComment)
	r.GET("/comments", middlewares.Authentication(), controllers.GetComment)
	r.PUT("/comments/:id", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.UpdateComment)
	r.DELETE("/comments/:id", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.DeleteComment)
	r.POST("/socialmedias", middlewares.Authentication(), controllers.CreateSocmed)
	r.GET("/socialmedias", middlewares.Authentication(), controllers.GetSocmed)
	r.PUT("/socialmedias/:id", middlewares.Authentication(), middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
	r.DELETE("/socialmedias/:id", middlewares.Authentication(), middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	return r
}
