package inmemory

import (
	taskDomain "br-lesson-4/internal/domain/task/models"
)

type Storage struct {
	tasks []taskDomain.Task
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		tasks: make([]taskDomain.Task, 0),
	}
}
