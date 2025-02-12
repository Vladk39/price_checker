package storage

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	con := "host=localhost user= password= dbname=pgndb port=5432 sslmode=disable"
	var err error

	DB, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		log.Print("ошибка подключения к бд", err)
		os.Exit(1)
	}
	fmt.Println("Подключение к базе данных активно", DB)
}
