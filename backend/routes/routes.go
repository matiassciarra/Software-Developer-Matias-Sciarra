package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola amiguito",
		})
	})
	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})

	r.GET("/api_externa", func(c *gin.Context) {
		// Tu token de autenticación
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MTcsImVtYWlsIjoibWF0aWFzc2NpYXJyYTZAZ21haWwuY29tIiwiZXhwIjoxNzQ1Nzg0MzczLCJpZCI6IjAiLCJwYXNzd29yZCI6IicgT1IgJzEnPScxIn0.d__b_x5CP3yAyDwxA_l5uMlH0lxEunnHH8ta-mj_96M"

		// Crear un nuevo request
		req, err := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list", nil)
		if err != nil {
			fmt.Println("Error al crear la petición:", err)
			return
		}
		// Agregar el encabezado de autorización Bearer
		req.Header.Add("Authorization", "Bearer "+token)

		// Ejecutar la petición
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error al hacer la petición:", err)
			return
		}
		defer resp.Body.Close()

		// Leer la respuesta
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			return
		}

		// Opción 1: Desserializar a un map[string]interface{} (para JSON genérico)
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la respuesta JSON"})
			return
		}
		// Devolver el JSON al cliente
		c.JSON(http.StatusOK, result)

	})

}
