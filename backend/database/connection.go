package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnection() {

	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println("Conexion establecida correctamente")
}
