package main

import (
	"firsttodo.com/database"
	"firsttodo.com/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Mongo()

	r.Use(cors.Default())

	routes.Todoroute(r)

	r.Run("localhost:8080")

}
