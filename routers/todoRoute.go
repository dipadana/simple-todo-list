package routers

import (
	"github.com/gin-gonic/gin"
	"simple-todo/controllers"
)

func StartServer() *gin.Engine {
	var router = gin.Default()

	router.GET("/todo", controllers.IndexTodo)
	router.GET("/todo/:id", controllers.ShowTodo)
	router.POST("/todo", controllers.CreateTodo)
	router.PUT("/todo/:id", controllers.EditTodo)
	router.DELETE("/todo/:id", controllers.DeleteTodo)

	return router
}
