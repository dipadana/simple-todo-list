package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Todo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var TodoData = []Todo{}

func CreateTodo(ctx *gin.Context) {
	var newTodo Todo

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTodo.Id = uuid.New().String()
	TodoData = append(TodoData, newTodo)

	ctx.JSON(http.StatusCreated, newTodo)
}

func IndexTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, TodoData)
}

func ShowTodo(ctx *gin.Context) {
	var notFound = true
	for _, todo := range TodoData {
		if todo.Id == ctx.Param("id") {
			notFound = false
			ctx.JSON(http.StatusOK, todo)
		}
	}

	if notFound {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin)
	}
}
