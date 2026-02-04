package config

import (
	"log"
	"ozenshe/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ozenshe.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Age{},
		&models.Genre{},
		&models.Category{},
		&models.Season{},
		&models.Episode{},
		&models.Screenshot{},
		&models.Movie{},
	)
}
