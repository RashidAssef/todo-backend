package routes

import (
	ct "firsttodo.com/CT"
	dt "firsttodo.com/DT"
	gt "firsttodo.com/GT"
	gtid "firsttodo.com/GTid"
	ud "firsttodo.com/UD"
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
