package controllers

import (
	"net/http"
	"ozenshe/config"
	"ozenshe/models"

	"github.com/gin-gonic/gin"
)

func GetAges(c *gin.Context) {
	var ages []models.Age
	config.DB.Find(&ages)
	c.JSON(http.StatusOK, ages)
}

func CreateAge(c *gin.Context) {
	var age models.Age
	if err := c.ShouldBindJSON(&age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&age)
	c.JSON(http.StatusOK, age)
}

func UpdateAge(c *gin.Context) {
	var age models.Age
	id := c.Param("id")

	if err := config.DB.First(&age, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Age not found"})
		return
	}

	if err := c.ShouldBindJSON(&age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&age)
	c.JSON(http.StatusOK, age)
}

func DeleteAge(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Age{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
