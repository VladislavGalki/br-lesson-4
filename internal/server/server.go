package server

import (
	taskDomain "br-lesson-4/internal/domain/task/models"
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

type Storage interface {
	TaskStorage
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

	s.srv.Handler = router
}
