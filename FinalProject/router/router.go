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
	return r
}
