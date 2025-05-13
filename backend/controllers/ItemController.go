package controllers

import (
	"BackendGo/database"
	"BackendGo/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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

func (ic *ItemController) FetchAndSaveExternalItems(c *gin.Context) {
	token := os.Getenv("TOKEN")
	baseURL := os.Getenv("LINK_API_EXTERNA")

	client := &http.Client{}
	totalItems := 0
	nextPage := ""

	for {
		// Construir la URL con la query si es necesario
		url := baseURL
		if nextPage != "" {
			url += "?next_page=" + nextPage
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error al crear la petición:", err)
			break
		}
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error al hacer la petición:", err)
			break
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			break
		}

		var apiResponse structs.ApiResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar la respuesta JSON"})
			return
		}

		for _, item := range apiResponse.Items {
			dbItem, err := transformItem(item)
			if err != nil {
				fmt.Println("Error transformando item:", err)
				continue
			}
			result := ic.db.Create(&dbItem)
			if result.Error != nil {
				fmt.Println("Error al guardar item en la base de datos:", result.Error)
				continue
			}
			totalItems++
		}

		if apiResponse.NextPage == "" {
			break
		}
		nextPage = apiResponse.NextPage
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d items obtenidos y procesados.", totalItems)})
}

func (ic *ItemController) GetTopRatedItems(c *gin.Context) {
	var items []structs.DbItem

	// Obtener todos los items de la base de datos
	result := ic.db.Find(&items)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los items"})
		return
	}

	// Calcular el porcentaje de aumento y ordenar los items
	type ItemWithPercentage struct {
		Item            structs.DbItem `json:"item"`
		PercentIncrease float32        `json:"percentIncrease"`
	}

	var itemsWithPercentage []ItemWithPercentage
	for _, item := range items {
		var percentIncrease float32 = 0
		// Evitar división por cero
		if item.TargetFrom > 0 {
			// Calcular el porcentaje de aumento: ((TargetTo - TargetFrom) / TargetFrom) * 100
			percentIncrease = ((item.TargetTo - item.TargetFrom) / item.TargetFrom) * 100
		}

		itemsWithPercentage = append(itemsWithPercentage, ItemWithPercentage{
			Item:            item,
			PercentIncrease: percentIncrease,
		})
	}

	// Ordenar los items por el porcentaje de aumento en orden descendente
	sort.Slice(itemsWithPercentage, func(i, j int) bool {
		return itemsWithPercentage[i].PercentIncrease > itemsWithPercentage[j].PercentIncrease
	})

	// Devolver los items ordenados por porcentaje de aumento
	c.JSON(http.StatusOK, gin.H{"items": itemsWithPercentage})
}

func (ic *ItemController) GetBestInvestmentRecommendations(c *gin.Context) {
	var items []structs.DbItem
	result := ic.db.Find(&items)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los items"})
		return
	}

	type ScoredItem struct {
		Item  structs.DbItem `json:"item"`
		Score float32        `json:"score"`
	}

	var scoredItems []ScoredItem

	for _, item := range items {
		var score float32 = 0

		// Factor 1: Porcentaje de aumento del precio objetivo (40%)
		var targetPercentIncrease float32 = 0
		if item.TargetFrom > 0 {
			targetPercentIncrease = ((item.TargetTo - item.TargetFrom) / item.TargetFrom) * 100
			score += targetPercentIncrease * 10
		}

		// Factor 2: Mejora en la calificación (30%)
		var ratingScore float32 = 0
		// Convertir las calificaciones a puntuaciones numéricas
		ratingMap := map[string]float32{
			"Strong Buy":          5.0,
			"Strong-Buy":          5.0, // Alternativa con guion
			"Buy":                 4.0,
			"Outperform":          4.0,
			"Market Outperform":   4.0, // Nuevo
			"Sector Outperform":   4.0,
			"Overweight":          3.5,
			"Hold":                3.0,
			"Neutral":             3.0,
			"Equal Weight":        3.0,
			"Peer Perform":        3.0, // Nuevo
			"Sector Perform":      3.0,
			"Market Perform":      3.0,
			"Sector Weight":       3.0, // Nuevo
			"In-Line":             3.0, // Nuevo
			"Underperform":        2.0,
			"Sector Underperform": 2.0, // Nuevo
			"Underweight":         1.5,
			"Sell":                1.0,
			"Strong Sell":         0.5,
		}

		fromRating := ratingMap[item.RatingFrom]
		toRating := ratingMap[item.RatingTo]
		ratingScore = (toRating - fromRating) * 100 // Multiplicamos por 10 para dar más peso
		score += ratingScore * 5

		scoredItems = append(scoredItems, ScoredItem{
			Item:  item,
			Score: score,
		})
	}

	// Ordenar por puntuación más alta
	sort.Slice(scoredItems, func(i, j int) bool {
		return scoredItems[i].Score > scoredItems[j].Score
	})

	c.JSON(http.StatusOK, gin.H{"recommendations": scoredItems})
}

func (ic *ItemController) GetAllItems(c *gin.Context) {
	var items []structs.DbItem

	result := ic.db.Find(&items)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los items"})
		return
	}

	c.JSON(http.StatusOK, items)
}
func (ic *ItemController) GetUniqueRatingFrom(c *gin.Context) {
	var uniqueRatings []string

	// Usar SELECT DISTINCT con una condición para excluir valores vacíos o NULL
	result := ic.db.Model(&structs.DbItem{}).
		Where("rating_from IS NOT NULL AND rating_from != ''").
		Distinct("rating_from").
		Pluck("rating_from", &uniqueRatings)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los valores únicos de RatingFrom"})
		return
	}

	c.JSON(http.StatusOK, uniqueRatings)
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

	// Eliminar comas como separadores de miles
	cleaned := strings.ReplaceAll(s, ",", "")

	f, err := strconv.ParseFloat(cleaned, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid float value: %w", err)
	}
	return float32(f), nil
}
