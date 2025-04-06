package ud

import (
	"context"
	"net/http"
	"time"

	"firsttodo.com/database"
	"firsttodo.com/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *gin.Context) {
	idparam := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(idparam)

	var updatetodo model.Todo
	err := c.BindJSON(&updatetodo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"task":   updatetodo.Task,
		"status": updatetodo.Status,
	}}

	_, err = database.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}
