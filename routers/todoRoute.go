package routers

import (
	"github.com/gin-gonic/gin"
	"simple-todo/controllers"
)

func StartServer() *gin.Engine {
	var router = gin.Default()

	router.GET("/todo", controllers.IndexTodo)
	router.GET("/todo/:id", controllers.IndexTodo)
	router.POST("/todo", controllers.CreateTodo)

	return router
}
