package routes

import (
	"example/go-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("users")
		{
			books.GET("/", controllers.ReadUsers)
			books.GET("/:id", controllers.ReadUser)
			books.POST("/", controllers.CreateUser)
			books.PUT("/", controllers.UpdateUser)
			books.DELETE("/:id", controllers.DeleteUser)
		}
	}

	return router
}
