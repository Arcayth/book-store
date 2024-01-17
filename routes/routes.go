package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/haythamchanouni/book-store/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	bookRoutes := router.Group("/api/books")

	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBookById)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
	return router
}
