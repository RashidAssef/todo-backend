package ct

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/RashidAssef/todo-backend/database"
	"github.com/RashidAssef/todo-backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodos(c *gin.Context) {
	var todo model.Todo

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		log.Println("Invalid Input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	todo.ID = primitive.NewObjectID()
	todo.Status = false

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.Collection.InsertOne(ctx, todo)

	if err != nil {
		log.Println("Insert Failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)

}
