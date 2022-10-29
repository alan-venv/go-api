package main

import (
	"example/go-api/src/database"
	"example/go-api/src/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := os.Remove("test.db") // remove a single file
	if err != nil {
		fmt.Println(err)
	}

	database.Start()
	gin.SetMode(gin.ReleaseMode)

	port := "8000"
	server := gin.Default()
	server = routes.ConfigRoutes(server)

	log.Print("Server is running at port: ", port)
	log.Fatal(server.Run(":" + port))
}
