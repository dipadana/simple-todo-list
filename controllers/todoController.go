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
			return
		}
	}

	if notFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not Found"})
		return
	}
}

func EditTodo(ctx *gin.Context) {
	var notFound = true
	var newTodo Todo
	var foundTodo Todo
	var tempTodo = []Todo{}

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, todo := range TodoData {
		if todo.Id == ctx.Param("id") {
			notFound = false
			foundTodo = todo
			break
		}
	}

	if notFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not Found"})
		return
	}

	for _, todo := range TodoData {
		if todo.Id != ctx.Param("id") {
			tempTodo = append(tempTodo, todo)
		}
	}

	foundTodo.Title = newTodo.Title
	TodoData = append(tempTodo, foundTodo)
	ctx.JSON(http.StatusOK, gin.H{"message": "Data Edited"})
}

func DeleteTodo(ctx *gin.Context) {
	var tempTodo = []Todo{}

	for _, todo := range TodoData {
		if todo.Id != ctx.Param("id") {
			tempTodo = append(tempTodo, todo)
		}
	}

	TodoData = tempTodo
	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
