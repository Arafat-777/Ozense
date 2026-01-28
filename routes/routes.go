package routes

import (
	"net/http"
	"ozenshe/controllers"
	"ozenshe/middleware"

	"github.com/gin-gonic/gin"

)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/genres", controllers.GetGenres)
	r.GET("/categories", controllers.GetCategories)
	r.GET("/ages", controllers.GetAges)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		admin.GET("/check", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Добро пожаловать, админ!"})
		})

		admin.POST("/movies", controllers.CreateMovie)
		admin.PUT("/movies/:id", controllers.UpdateMovie)
		admin.DELETE("/movies/:id", controllers.DeleteMovie)

		admin.POST("/genres", controllers.CreateGenre)
		admin.PUT("/genres/:id", controllers.UpdateGenre)
		admin.DELETE("/genres/:id", controllers.DeleteGenre)

		admin.POST("/categories", controllers.CreateCategory)
		admin.PUT("/categories/:id", controllers.UpdateCategory)
		admin.DELETE("/categories/:id", controllers.DeleteCategory)

		admin.POST("/ages", controllers.CreateAge)
		admin.PUT("/ages/:id", controllers.UpdateAge)
		admin.DELETE("/ages/:id", controllers.DeleteAge)
	}

	// общие просмотровые маршруты
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)

	return r
}
