package inmemory

import (
	taskDomain "br-lesson-4/internal/domain/task/models"
	userDomain "br-lesson-4/internal/domain/user/models"
)

type Storage struct {
	tasks []taskDomain.Task
	users []userDomain.User
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		tasks: make([]taskDomain.Task, 0),
		users: make([]userDomain.User, 0),
	}
}
