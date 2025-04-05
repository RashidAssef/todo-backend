package main

import (
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
	r.Run(":" + port)
}
