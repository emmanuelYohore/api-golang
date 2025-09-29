package routes

import (
	"github.com/emmanuelYohore/api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/api/books", controllers.GetBooks)
	router.GET("/api/book/:id", controllers.GetBook)
	router.POST("/api/book", controllers.CreateBook)
	router.PUT("/api/books/:id", controllers.UpdateBook)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
