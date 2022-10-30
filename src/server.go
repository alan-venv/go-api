package main

import (
	"example/go-api/src/database"
	"example/go-api/src/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Data base
	database.Start()

	// Gin
	gin.SetMode(os.Getenv("API_MODE"))

	server := gin.Default()
	server = routes.ConfigRoutes(server)

	log.Print("Server is running at port: ", os.Getenv("API_PORT"))
	log.Fatal(server.Run(":" + os.Getenv("API_PORT")))
}
