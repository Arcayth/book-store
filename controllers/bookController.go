package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haythamchanouni/book-store/database"
	"github.com/haythamchanouni/book-store/models"
	"github.com/haythamchanouni/book-store/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	booksCollection := database.GetCollection("books")
	filter := bson.D{{"_id", id}}
	if err := booksCollection.FindOne(context.Background(), filter).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
			return
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	utils.JSONResponse(c, http.StatusOK, book)

}

func CreateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UpdateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DeleteBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
