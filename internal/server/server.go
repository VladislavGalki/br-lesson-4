package server

import (
	taskDomain "br-lesson-4/internal/domain/task/models"
	userDomain "br-lesson-4/internal/domain/user/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskStorage interface {
	GetTasksList() ([]taskDomain.Task, error)
	GetTaskByID(string) (taskDomain.Task, error)
	CreateTask(domainTask taskDomain.Task) (taskDomain.Task, error)
	UpdateTask(id string, domainTask taskDomain.Task) (taskDomain.Task, error)
	DeleteTask(id string) error
}

type UserStorage interface {
	GetUserList() ([]userDomain.User, error)
	GetUseByID(string) (userDomain.User, error)
	CreateUser(domainTask userDomain.User) (userDomain.User, error)
	UpdateUser(id string, domainTask userDomain.User) (userDomain.User, error)
	DeleteUser(id string) error
}

type Storage interface {
	TaskStorage
	UserStorage
}

type ToDoAPI struct {
	srv     *http.Server
	storage Storage
}

func NewToDoServer(storage Storage) *ToDoAPI {
	httpServer := http.Server{
		Addr: "localhost:8080",
	}

	api := ToDoAPI{
		srv:     &httpServer,
		storage: storage,
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

	tasks := router.Group("/tasks")
	tasks.GET("/list", s.getTaskList)
	tasks.GET("/task/:id", s.getTask)
	tasks.POST("/add-task", s.addTask)
	tasks.PUT("/update-task/:id", s.updateTask)
	tasks.DELETE("/delete-task/:id", s.deleteTask)

	users := router.Group("/users")
	users.GET("/user-list", s.GetUserList)
	users.GET("/user/:id", s.GetUserById)
	users.POST("/create-user", s.CreateUser)
	users.PUT("/update-user/:id", s.UpdateUser)
	users.DELETE("/delete-user/:id", s.DeleteUser)
	s.srv.Handler = router
}
