package server

import (
	taskDomain "br-lesson-4/internal/domain/task/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ToDoAPI) getTaskList(ctx *gin.Context) {
	tasks, err := s.storage.GetTasksList()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}

func (s *ToDoAPI) getTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := s.storage.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"task": task})
	return
}

func (s *ToDoAPI) addTask(ctx *gin.Context) {
	var task taskDomain.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := s.storage.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"task": newTask})
	return
}

func (s *ToDoAPI) updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task taskDomain.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := s.storage.UpdateTask(id, task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": updatedTask})
	return
}

func (s *ToDoAPI) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := s.storage.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": nil})
	return
}
