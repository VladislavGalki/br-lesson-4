package server

import (
	"br-lesson-4/internal"
	taskDomain "br-lesson-4/internal/domain/task/models"
	userDomain "br-lesson-4/internal/domain/user/models"
	"br-lesson-4/internal/server/auth"
	"br-lesson-4/internal/server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type TaskStorage interface {
	GetTasksList(userId string) ([]taskDomain.Task, error)
	GetTaskByID(string) (taskDomain.Task, error)
	CreateTask(domainTask taskDomain.Task) (taskDomain.Task, error)
	UpdateTask(id string, domainTask taskDomain.Task) (taskDomain.Task, error)
	DeleteTask(id string) error
}

type UserStorage interface {
	GetUserList() ([]userDomain.User, error)
	GetUseByID(string) (userDomain.User, error)
	GetUser(userReq userDomain.UserRequest) (userDomain.User, error)
	CreateUser(domainUser userDomain.User) (userDomain.User, error)
	UpdateUser(id string, domainUser userDomain.User) (userDomain.User, error)
	DeleteUser(id string) error
}

type Storage interface {
	TaskStorage
	UserStorage
}

type ToDoAPI struct {
	srv       *http.Server
	storage   Storage
	jwtSigner auth.HS256Signer
}

func NewToDoServer(config internal.Config, storage Storage) *ToDoAPI {
	jwtSigner := auth.HS256Signer{
		Secret:     []byte(uuid.NewString()),
		Issuer:     "task-service",
		Audience:   "task-client",
		AccessTTL:  15 * time.Minute,
		RefreshTTL: 24 * 7 * time.Hour,
	}

	httpServer := http.Server{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	}

	api := ToDoAPI{
		srv:       &httpServer,
		storage:   storage,
		jwtSigner: jwtSigner,
	}

	api.configRouter()

	return &api
}

func (s *ToDoAPI) Start() error {
	return s.srv.ListenAndServe()
}

func (s *ToDoAPI) Stop() error {
	return s.srv.Shutdown(nil)
}

func (s *ToDoAPI) configRouter() {
	router := gin.Default()

	router.POST("/login", s.LoginUser)
	router.GET("/profile", middleware.AuthMiddleware(s.jwtSigner), s.ProfileUser)
	router.POST("/refresh", s.refresh)

	tasks := router.Group("/tasks")
	tasks.GET("/list", middleware.AuthMiddleware(s.jwtSigner), s.getTaskList)
	tasks.GET("/task/:id", middleware.AuthMiddleware(s.jwtSigner), s.getTask)
	tasks.POST("/add-task", middleware.AuthMiddleware(s.jwtSigner), s.addTask)
	tasks.PUT("/update-task/:id", middleware.AuthMiddleware(s.jwtSigner), s.updateTask)
	tasks.DELETE("/delete-task/:id", middleware.AuthMiddleware(s.jwtSigner), s.deleteTask)

	users := router.Group("/users")
	users.GET("/user-list", s.GetUserList)
	users.GET("/user/:id", s.GetUserById)
	users.POST("/create-user", s.CreateUser)
	users.PUT("/update-user/:id", middleware.AuthMiddleware(s.jwtSigner), s.UpdateUser)
	users.DELETE("/delete-user/:id", middleware.AuthMiddleware(s.jwtSigner), s.DeleteUser)
	s.srv.Handler = router
}

func (s *ToDoAPI) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, err := s.jwtSigner.ParseRefreshToken(refreshToken, auth.ParseOptions{
		ExpectedIssuer:   s.jwtSigner.Issuer,
		ExpectedAudience: s.jwtSigner.Audience,
		AllowedMethods:   []string{"HS256"},
		Leeway:           60 * time.Second,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	access, err := s.jwtSigner.NewAccessToken(claims.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newRefresh, err := s.jwtSigner.NewRefreshToken(claims.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", newRefresh, 3600*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"access": access})
}
