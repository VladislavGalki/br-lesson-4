package server

import (
	userDomain "br-lesson-4/internal/domain/user/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ToDoAPI) GetUserList(ctx *gin.Context) {
	users, err := s.storage.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
	return
}

func (s *ToDoAPI) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := s.storage.GetUseByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func (s *ToDoAPI) CreateUser(ctx *gin.Context) {
	var user userDomain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := s.storage.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": createdUser})
	return
}

func (s *ToDoAPI) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user userDomain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := s.storage.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": updatedUser})
	return
}

func (s *ToDoAPI) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.storage.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": nil})
	return
}
