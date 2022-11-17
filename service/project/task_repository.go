package project

import (
	"context"
	"errors"
	"fmt"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, projectID string, task Task) (string, error)
	ReadTask(ctx context.Context, projectID string, taskID string) (Task, error)
	UpdateTask(ctx context.Context, projectID string, task Task) error
	DeleteTask(ctx context.Context, projectID string, taskID string) error
	DeleteTasks(ctx context.Context, projectID string) error
	ListTasks(ctx context.Context, projectID string) ([]Task, error)
}

type taskRepository struct {
	task map[string][]Task
}

func NewTaskRepository() (TaskRepository, error) {
	return &taskRepository{task: map[string][]Task{}}, nil
}

func (s *taskRepository) CreateTask(ctx context.Context, projectID string, task Task) (string, error) {

	id, err := generateIDWhenNotGiven(task.ID)
	if err != nil {
		return "", fmt.Errorf("repository[create-task]: %w", err)
	}

	task.ID = id
	_, err = s.ReadTask(ctx, projectID, task.ID)
	if err == nil {
		return "", errors.New("repository[create-task]: task found")
	}

	s.task[projectID] = append(s.task[projectID], task)

	return id, nil
}

func (s *taskRepository) ReadTask(ctx context.Context, projectID string, taskID string) (Task, error) {

	for _, currentTask := range s.task[projectID] {
		if currentTask.ID == taskID {
			return currentTask, nil
		}
	}

	return Task{}, errors.New("repository[read-task]: task not found")
}

func (s *taskRepository) UpdateTask(ctx context.Context, projectID string, task Task) error {

	if _, err := s.ReadTask(ctx, projectID, task.ID); err != nil {
		return errors.New("repository[update-task]: task not found")
	}

	for index, currentTask := range s.task[projectID] {
		if currentTask.ID == task.ID {
			s.task[projectID][index] = task
		}
	}

	return nil
}

func (s *taskRepository) DeleteTask(ctx context.Context, projectID string, taskID string) error {

	for index, task := range s.task[projectID] {
		if task.ID == taskID {
			s.task[projectID] = append(s.task[projectID][:index], s.task[projectID][index+1:]...)
			return nil
		}
	}

	return errors.New("repository[delete-task]: task not found")
}

func (s *taskRepository) DeleteTasks(ctx context.Context, projectID string) error {

	delete(s.task, projectID)

	return nil
}

func (s *taskRepository) ListTasks(ctx context.Context, projectID string) (tasks []Task, err error) {

	for _, currentTask := range s.task[projectID] {
		tasks = append(tasks, currentTask)
	}

	return tasks, nil
}
