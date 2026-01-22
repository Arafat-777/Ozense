package config

import (
	"log"
	"ozenshe/models"

	"github.com/glebarez/sqlite" // 💡 драйвер, работающий без CGO
	"gorm.io/gorm"

)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ozenshe.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	DB.AutoMigrate(&models.Genre{}, &models.Category{})
}
