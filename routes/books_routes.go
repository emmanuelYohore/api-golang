package routes

import (
	"github.com/emmanuelYohore/api-golang/controllers"
	"github.com/emmanuelYohore/api-golang/database"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()

	database.Connect()

	router.GET("/api/books", controllers.GetBooks)
	router.GET("/api/books/:id", controllers.GetBook)
	router.POST("/api/books", controllers.CreateBook)
	router.PUT("/api/books/:id", controllers.UpdateBook)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
