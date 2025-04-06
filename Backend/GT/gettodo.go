package gt

import (
	"context"
	"log"
	"net/http"
	"time"

	"firsttodo.com/database"
	"firsttodo.com/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.Collection.Find(ctx, bson.M{})

	if err != nil {
		log.Println("Invalid search:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo model.Todo

		err := cursor.Decode(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode"})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)

}
