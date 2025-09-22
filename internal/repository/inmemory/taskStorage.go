package inmemory

import (
	taskError "br-lesson-4/internal/domain/task/errors"
	taskDomain "br-lesson-4/internal/domain/task/models"
	"github.com/google/uuid"
)

func (storage *Storage) GetTasksList(userId string) ([]taskDomain.Task, error) {
	if len(storage.tasks) == 0 {
		return nil, taskError.ErrTaskListNotFound
	}

	return storage.tasks, nil
}

func (storage *Storage) GetTaskByID(id string) (taskDomain.Task, error) {
	if len(storage.tasks) == 0 {
		return taskDomain.Task{}, taskError.ErrTaskNotFound
	}

	for _, task := range storage.tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return taskDomain.Task{}, taskError.ErrTaskNotFound
}

func (storage *Storage) CreateTask(domainTask taskDomain.Task) (taskDomain.Task, error) {
	for _, task := range storage.tasks {
		if task.Id == domainTask.Id {
			return taskDomain.Task{}, taskError.ErrTaskExists
		}
	}

	domainTask.Id = uuid.NewString()
	storage.tasks = append(storage.tasks, domainTask)
	return domainTask, nil
}

func (storage *Storage) UpdateTask(id string, domainTask taskDomain.Task) (taskDomain.Task, error) {
	for index, task := range storage.tasks {
		if task.Id == id {
			storage.tasks[index] = domainTask
			return domainTask, nil
		}
	}

	return taskDomain.Task{}, taskError.ErrTaskNotFound
}

func (storage *Storage) DeleteTask(id string) error {
	for index, task := range storage.tasks {
		if task.Id == id {
			storage.tasks = append(storage.tasks[:index], storage.tasks[index+1:]...)
			return nil
		}
	}

	return taskError.ErrTaskNotFound
}
