package errors

import "errors"

var (
	ErrTaskListNotFound = errors.New("Tasks not found")
	ErrTaskNotFound     = errors.New("Task not found")
	ErrTaskExists       = errors.New("Task already exists")
)
