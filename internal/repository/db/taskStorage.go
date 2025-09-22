package db

import (
	"br-lesson-4/internal/domain"
	taskDomain "br-lesson-4/internal/domain/task/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type taskStorage struct {
	db *pgx.Conn
}

func (s *taskStorage) GetTasksList(userId string) ([]taskDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var tasks []taskDomain.Task

	rows, err := s.db.Query(ctx, "SELECT * FROM tasks WHERE userid = $1", userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task taskDomain.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Status, &task.UserId); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskStorage) GetTaskByID(id string) (taskDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var task taskDomain.Task
	row := s.db.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id)
	if err := row.Scan(&task.Id, &task.Name, &task.Description, &task.Status); err != nil {
		return taskDomain.Task{}, err
	}

	return task, nil
}

func (s *taskStorage) CreateTask(domainTask taskDomain.Task) (taskDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := s.db.Exec(
		ctx,
		"INSERT into tasks (id, name, description, status, userId) VALUES ($1, $2, $3, $4, $5)",
		domainTask.Id,
		domainTask.Name,
		domainTask.Description,
		domainTask.Status,
		domainTask.UserId,
	)

	if err != nil {
		return taskDomain.Task{}, err
	}

	return domainTask, nil
}

func (s *taskStorage) UpdateTask(id string, domainTask taskDomain.Task) (taskDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := s.db.Exec(
		ctx,
		"UPDATE tasks SET name = $1, description = $2, status = $3 WHERE id = $4",
		domainTask.Name,
		domainTask.Description,
		domainTask.Status,
		id,
	)
	if err != nil {
		return taskDomain.Task{}, err
	}

	return domainTask, nil
}

func (s *taskStorage) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := s.db.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
