// Se encarga de crear el servidor
package main

import (
	"BackendGo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")

}
