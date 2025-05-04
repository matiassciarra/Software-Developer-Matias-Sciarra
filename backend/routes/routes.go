package routes

import (
	"BackendGo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	itemController := controllers.NewItemController()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola amiguito",
		})
	})
	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})

	r.GET("/api_externa", itemController.FetchAndSaveExternalItems)

}
