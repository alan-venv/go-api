package main

import (
	"example/go-api/src/configs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("API_MODE"))

	server := gin.Default()
	server = configs.ServerConfig(server)

	log.Print("Server is running at port: ", os.Getenv("API_PORT"))
	log.Fatal(server.Run(":" + os.Getenv("API_PORT")))
}
