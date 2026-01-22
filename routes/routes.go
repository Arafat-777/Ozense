package routes

import (
	"ozenshe/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	genre := r.Group("/genres")
	{
		genre.GET("/", controllers.GetGenres)
		genre.POST("/", controllers.CreateGenre)
		genre.PUT("/:id", controllers.UpdateGenre)
		genre.DELETE("/:id", controllers.DeleteGenre)
	}

	category := r.Group("/categories")
	{
		category.GET("/", controllers.GetCategories)
		category.POST("/", controllers.CreateCategory)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}

	return r
}
