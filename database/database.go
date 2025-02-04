package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"premium_microservice/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Bayadilov06kz dbname=restapi_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	DB = db
	db.AutoMigrate(&models.Transaction{})
	log.Println("База данных подключена и мигрирована")
}
