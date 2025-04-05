package main

import (
	"fmt"
	"os"

	"github.com/RashidAssef/todo-backend/database"
	"github.com/RashidAssef/todo-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Mongo()

	r.Use(cors.Default())

	routes.Todoroute(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server starting on port:", port)
	r.Run("0.0.0.0:" + port)
}
