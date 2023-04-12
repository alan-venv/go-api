package main

import (
	"example/go-api/src/configs"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	env := configs.GetEnv("API_ENV", "LOCAL")
	mode := configs.GetEnv("API_MODE", "release")
	port := configs.GetEnv("API_PORT", "8080")

	gin.SetMode(mode)

	server := gin.Default()
	server = configs.ServerConfig(server, env)

	log.Print("Server is running at port: ", port)
	log.Fatal(server.Run(":" + port))
}
