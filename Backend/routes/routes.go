package routes

import (
	ct "github.com/RashidAssef/todo-backend/CT"
	dt "github.com/RashidAssef/todo-backend/DT"
	gt "github.com/RashidAssef/todo-backend/GT"
	gtid "github.com/RashidAssef/todo-backend/GTid"
	ud "github.com/RashidAssef/todo-backend/UD"
	"github.com/gin-gonic/gin"
)

func Todoroute(r *gin.Engine) {
	todoRT := r.Group("/todos")

	{
		todoRT.GET("/", gt.GetTodos)
		todoRT.GET("/:id", gtid.GetTodoswithid)
		todoRT.POST("/", ct.CreateTodos)
		todoRT.DELETE("/:id", dt.DeleteTodo)
		todoRT.PUT("/:id", ud.UpdateTodo)
	}
}
