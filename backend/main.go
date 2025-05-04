// Se encarga de crear el servidor
package main

import (
	"BackendGo/database"
	"BackendGo/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	r := gin.Default()
	database.InitDB()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
