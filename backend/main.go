// Se encarga de crear el servidor
package main

import (
	"BackendGo/database"
	"BackendGo/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
	fmt.Println("Connection string:", os.Getenv("CONNECTION_STRING"))

	r := gin.Default()
	routes.SetupRoutes(r)

	database.TestConnection()
	r.Run(":8080")
}
