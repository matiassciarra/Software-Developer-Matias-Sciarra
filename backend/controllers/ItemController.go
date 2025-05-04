package controllers

import (
	"BackendGo/database"
	"BackendGo/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ItemController define los métodos para manejar las operaciones de items.
type ItemController struct {
	db *gorm.DB
}

// NewItemController crea una nueva instancia del ItemController.
func NewItemController() *ItemController {
	return &ItemController{db: database.GetDB()}
}

// FetchAndSaveExternalItems llama a la API externa y guarda los items en la base de datos.
func (ic *ItemController) FetchAndSaveExternalItems(c *gin.Context) {

	// Tu token de autenticación
	token := os.Getenv("TOKEN")

	// Crear un nuevo request
	req, err := http.NewRequest("GET", os.Getenv("LINK_API_EXTERNA"), nil)
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

	var apiResponse structs.ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la respuesta JSON"})
		return
	}

	for _, item := range apiResponse.Items {
		dbItem, err := transformItem(item)
		if err != nil {
			fmt.Println("Error parseando")
		}
		fmt.Println(dbItem)
		result := ic.db.Create(&dbItem)
		if result.Error != nil {
			fmt.Println("Error al guardar item en la base de datos:", result.Error) // Imprime el error de GORM
			// Puedes agregar un return aquí si quieres detener el proceso al primer error
			// return
			continue // Continua con el siguiente item
		} else {
			fmt.Println("Item guardado exitosamente")
		}

	}
	// Devolver el JSON al cliente
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d items obtenidos y procesados.", len(apiResponse.Items))})
}

func transformItem(item structs.Item) (structs.DbItem, error) {

	var dbItem structs.DbItem

	dbItem.Ticker = item.Ticker
	dbItem.Company = item.Company
	dbItem.Action = item.Action
	dbItem.Brokerage = item.Brokerage
	dbItem.RatingFrom = item.RatingFrom
	dbItem.RatingTo = item.RatingTo

	// Parsear TargetFrom a float32
	targetFromFloat, err := parseFloat32(item.TargetFrom)
	if err != nil {
		return structs.DbItem{}, fmt.Errorf("error parsing TargetFrom: %w", err) // Wrap original error
	}
	dbItem.TargetFrom = targetFromFloat

	// Parsear TargetTo a float32
	targetToFloat, err := parseFloat32(item.TargetTo)
	if err != nil {
		return structs.DbItem{}, fmt.Errorf("error parsing TargetTo: %w", err) // Wrap original error
	}
	dbItem.TargetTo = targetToFloat

	// Parsear la cadena de tiempo a time.Time
	parsedTime, err := time.Parse(time.RFC3339Nano, item.Time)
	if err != nil {
		return structs.DbItem{}, fmt.Errorf("error parsing Time: %w", err) // Wrap original error
	}
	dbItem.Time = parsedTime

	return dbItem, nil
}

// parseFloat32 es una función auxiliar para parsear cadenas a float32, manejando el símbolo '$'.
func parseFloat32(s string) (float32, error) {
	// Eliminar el símbolo '$' si está presente
	if len(s) > 0 && s[0] == '$' {
		s = s[1:]
	}
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid float value: %w", err) // Wrap the error
	}
	return float32(f), nil
}
