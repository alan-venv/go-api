package configs

import (
	"example/go-api/src/controllers"
	"example/go-api/src/database"
	"example/go-api/src/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func GetEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok && len(value) != 0 {
		return value
	}

	return fallback
}

func ServerConfig(server *gin.Engine, environment string) *gin.Engine {
	// Repositories
	var userRepository repositories.IUserRepository
	switch environment {
	case "LOCAL":
		userRepository = repositories.UserGormRepository{
			Database: database.StartSqlite(),
		}
	case "PROD":
		userRepository = repositories.UserMongoRepository{
			Database: database.StartMongo(),
		}
	default:
		log.Fatal("Invalid environment: " + environment)
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
			users.PUT("/", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}
	}

	return server
}
