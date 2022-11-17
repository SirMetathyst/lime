package project

import (
	"context"
	"errors"
	"fmt"
)

// todo: merge task service with project service?
type TaskService interface {
	CreateTask(ctx context.Context, projectID string, task Task) (string, error)
	ReadTask(ctx context.Context, projectID string, taskID string) (Task, error)
	UpdateTask(ctx context.Context, projectID string, task Task) error
	DeleteTask(ctx context.Context, projectID string, taskID string) error
	DeleteTasks(ctx context.Context, projectID string) error
	ListTasks(ctx context.Context, projectID string) ([]Task, error)
}

type taskService struct {
	taskRepository TaskRepository
}

func NewTaskService(r TaskRepository) (TaskService, error) {
	if r == nil {
		return nil, errors.New("task-service: task repository is nil")
	}
	return &taskService{taskRepository: r}, nil
}

func (s *taskService) CreateTask(ctx context.Context, projectID string, task Task) (string, error) {

	if err := TaskValid(task); err != nil {
		return "", fmt.Errorf("service[create-task]: %w", err)
	}

	id, err := s.taskRepository.CreateTask(ctx, projectID, task)
	if err != nil {
		return "", fmt.Errorf("task-service[create-task]: %w", err)
	}

	return id, nil
}

func (s *taskService) ReadTask(ctx context.Context, projectID string, taskID string) (Task, error) {

	task, err := s.taskRepository.ReadTask(ctx, projectID, taskID)
	if err != nil {
		return Task{}, fmt.Errorf("task-service[read-task]: %w", err)
	}

	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, projectID string, task Task) error {

	if err := TaskValid(task); err != nil {
		return fmt.Errorf("service[update-task]: %w", err)
	}

	err := s.taskRepository.UpdateTask(ctx, projectID, task)
	if err != nil {
		return fmt.Errorf("task-service[update-task]: %w", err)
	}

	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, projectID string, taskID string) error {

	err := s.taskRepository.DeleteTask(ctx, projectID, taskID)
	if err != nil {
		return fmt.Errorf("task-service[delete-task]: %w", err)
	}

	return nil
}

func (s *taskService) DeleteTasks(ctx context.Context, projectID string) error {

	err := s.taskRepository.DeleteTasks(ctx, projectID)
	if err != nil {
		return fmt.Errorf("task-service[delete-tasks]: %w", err)
	}

	return nil
}

func (s *taskService) ListTasks(ctx context.Context, projectID string) ([]Task, error) {

	tasks, err := s.taskRepository.ListTasks(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("task-service[list-tasks]: %w", err)
	}

	return tasks, nil
}
