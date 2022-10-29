package main

import (
	"example/go-api/src/configs"
	"example/go-api/src/database"
	"example/go-api/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Load()
	config := configs.Api()

	database.Start()
	gin.SetMode(config.Mode)

	server := gin.Default()
	server = routes.ConfigRoutes(server)

	log.Print("Server is running at port: ", config.Port)
	log.Fatal(server.Run(":" + config.Port))
}
