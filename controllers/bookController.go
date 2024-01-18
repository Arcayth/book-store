package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haythamchanouni/book-store/database"
	"github.com/haythamchanouni/book-store/models"
	"github.com/haythamchanouni/book-store/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	var booksCollection = database.GetCollection("booksDB", "books")

	filter := bson.D{}
	cursor, err := booksCollection.Find(context.TODO(), filter)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = cursor.All(context.Background(), &books); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(c, http.StatusOK, books)
	fmt.Println(books)

}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	filter := bson.D{{Key: "_id", Value: id}}
	var booksCollection = database.GetCollection("booksDB", "books")
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
