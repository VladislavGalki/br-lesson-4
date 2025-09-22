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

func (s *ToDoAPI) LoginUser(ctx *gin.Context) {
	var userRequest userDomain.UserRequest
	if err := ctx.BindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := s.storage.GetUser(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := s.jwtSigner.NewAccessToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := s.jwtSigner.NewRefreshToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("user_id", user.Id, 3600*24*7, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
	return
}

func (s *ToDoAPI) ProfileUser(ctx *gin.Context) {
	userId, exists := ctx.Get("userID")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "userId is required"})
		return
	}

	user, err := s.storage.GetUseByID(userId.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
	return
}
