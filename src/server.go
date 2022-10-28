package main

import (
	"example/go-api/src/database"
	"example/go-api/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Start()
	gin.SetMode(gin.ReleaseMode)

	port := "8000"
	server := gin.Default()
	server = routes.ConfigRoutes(server)

	log.Print("Server is running at port: ", port)
	log.Fatal(server.Run(":" + port))
}
