package configs

import (
	"example/go-api/src/controllers"
	"example/go-api/src/database"
	"example/go-api/src/repositories"

	"github.com/gin-gonic/gin"
)

func ServerConfig(server *gin.Engine) *gin.Engine {
	// Repositories
	database.Start()
	userRepository := repositories.UserMongoRepository{
		Database: database.Get(),
	}

	// Controllers
	userController := controllers.UserController{
		Repository: userRepository,
	}

	// Routes
	main := server.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", userController.ReadAll)
			users.GET("/:id", userController.Read)
			users.POST("/", userController.Create)
			//books.PUT("/", controllers.UpdateUser)
			users.DELETE("/:id", userController.Delete)
		}
	}

	return server
}
