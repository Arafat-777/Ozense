package controllers

import (
	"net/http"
	"ozenshe/config"
	"ozenshe/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMovie(c *gin.Context) {
	var input models.MovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		Trailer:     input.Trailer,
		AgeID:       input.AgeID,
	}

	if len(input.GenreIDs) > 0 {
		var genres []models.Genre
		if err := config.DB.Find(&genres, input.GenreIDs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Genres not found"})
			return
		}
		movie.Genres = genres
	}

	if len(input.CategoryIDs) > 0 {
		var categories []models.Category
		if err := config.DB.Find(&categories, input.CategoryIDs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Categories not found"})
			return
		}
		movie.Categories = categories
	}

	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	config.DB.
		Preload("Age"). // 👈 добавить это
		Preload("Genres").
		Preload("Categories").
		Preload("Seasons.Episodes").
		Preload("Screenshots").
		Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.
		Preload("Age"). // 👈 и здесь
		Preload("Genres").
		Preload("Categories").
		Preload("Seasons.Episodes").
		Preload("Screenshots").
		First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}


func UpdateMovie(c *gin.Context) {
	var m models.Movie
	id := c.Param("id")

	if err := config.DB.Preload("Genres").
		Preload("Categories").
		Preload("Seasons.Episodes").
		Preload("Screenshots").
		First(&m, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&m)
	c.JSON(http.StatusOK, m)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Movie{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
