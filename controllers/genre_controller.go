package controllers

import (
	"net/http"
	"ozenshe/config"
	"ozenshe/models"

	"github.com/gin-gonic/gin"

)

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	config.DB.Find(&genres)
	c.JSON(http.StatusOK, genres)
}

func CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&genre)
	c.JSON(http.StatusOK, genre)
}

func UpdateGenre(c *gin.Context) {
	var genre models.Genre
	id := c.Param("id")

	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Жанр не найден"})
		return
	}
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&genre)
	c.JSON(http.StatusOK, genre)
}

func DeleteGenre(c *gin.Context) {
	var genre models.Genre
	id := c.Param("id")

	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Жанр не найден"})
		return
	}
	config.DB.Delete(&genre)
	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}
