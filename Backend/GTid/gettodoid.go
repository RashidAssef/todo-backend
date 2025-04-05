package gtid

import (
	"context"
	"net/http"
	"time"

	"github.com/RashidAssef/todo-backend/database"
	"github.com/RashidAssef/todo-backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodoswithid(c *gin.Context) {
	idparam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idparam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connot find id"})
		return
	}

	var todo model.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = database.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)

}
