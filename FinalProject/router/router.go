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
	r.PUT("/users/:id", middlewares.Authentication(), middlewares.Authorization(), controllers.UpdateUser)
	r.DELETE("/users/:id", middlewares.Authentication(), middlewares.Authorization(), controllers.DeleteUser)
	r.POST("/photos", middlewares.Authentication(), middlewares.Authorization(), controllers.CreatePhoto)
	r.GET("/photos", middlewares.Authentication(), middlewares.Authorization(), controllers.GetPhoto)
	r.PUT("/photos/:id", controllers.UpdatePhoto)
	r.DELETE("/photos/:id", controllers.DeletePhoto)
	r.POST("/comments", middlewares.Authentication(), middlewares.Authorization(), controllers.CreateComment)
	r.GET("/comments", middlewares.Authentication(), middlewares.Authorization(), controllers.GetComment)
	r.PUT("/comments/:id", controllers.UpdateComment)
	r.DELETE("/comments/:id", controllers.DeleteComment)
	r.POST("/socialmedias", middlewares.Authentication(), middlewares.Authorization(), controllers.CreateSocmed)
	r.GET("/socialmedias", middlewares.Authentication(), middlewares.Authorization(), controllers.GetSocmed)
	r.PUT("/socialmedias/:id", controllers.UpdateSocmed)
	r.DELETE("/socialmedias/:id", controllers.DeleteSocmed)
	return r
}
