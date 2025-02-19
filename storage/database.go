package storage

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CryptoForDB struct {
	IDDB          int       `gorm:"primaryKey"`
	NameDB        string    `gorm:"namecrypto"`
	PriceDB       int       `gorm:"price"`
	PriceChangeDB float32   `gorm:"pricechange"`
	LastUpdate    time.Time `gorm:"type:timestamptz;default:now()"`
}

type BDRepository struct {
	DB *gorm.DB
}

func NewRequestBDRepository(db *gorm.DB) *BDRepository {
	// создаем новый объект бд репозитория
	return &BDRepository{DB: db}
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	con := "host=localhost user= password= dbname=pgndb port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		log.Print("ошибка подключения к бд", err)
		os.Exit(1)
	}
	fmt.Println("Подключение к базе данных активно")

	err = DB.AutoMigrate(&CryptoForDB{})
	if err != nil {
		log.Fatal("Ошибка миграции БД:", err)
	}

	return DB

}

func (repo *BDRepository) SendMsgDB(data []CryptoForDB) error {
	if len(data) == 0 {
		log.Println("Нет данных для записи в БД")
		return nil
	}
	for _, record := range data {
		result := repo.DB.Save(&record)
		if result.Error != nil {
			log.Println("Ошибка при сохранении данных:", result.Error)
			return result.Error
		}
	}

	log.Println("Данные успешно сохранены")
	return nil
}
