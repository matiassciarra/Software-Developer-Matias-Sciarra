package database

import (
	"BackendGo/models"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		var err error
		dsn := os.Getenv("CONNECTION_STRING")
		if dsn == "" {
			log.Fatal("CONNECTION_STRING environment variable not set") // IMPORTANT
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		// AutoMigrar el modelo DbItem
		err = db.AutoMigrate(&models.DbItem{}) // Pass the DbItem struct
		if err != nil {
			log.Fatalf("failed to automigrate DbItem: %v", err)
		}

		fmt.Println("Modelo Item migrado.")
		fmt.Println("Conexión a la base de datos establecida.")
		var now time.Time
		db.Raw("SELECT NOW()").Scan(&now)
		fmt.Println(now)
	})
}

func GetDB() *gorm.DB {
	InitDB()
	return db
}
