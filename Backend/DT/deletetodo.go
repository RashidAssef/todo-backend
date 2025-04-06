package dt

import (
	"context"
	"net/http"
	"time"

	"firsttodo.com/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *gin.Context) {
	idparam := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idparam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id not Found"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo Deleted"})

}
